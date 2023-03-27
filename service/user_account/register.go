package user_account

import (
	"errors"
	"my_todolist/middleware"
	"my_todolist/model"
)

const (
	MaxUsernameLength = 100
	MaxPasswordLength = 20
	MinPasswordLength = 8
)

type LoginResponse struct {
	UserId int64  `json:"userId"`
	Token  string `json:"token"`
}

func Register(username string, password string) (*LoginResponse, error) {
	return NewRegisterFlow(username, password).Do()
}

func NewRegisterFlow(username string, password string) *RegisterFlow {
	return &RegisterFlow{username: username, password: password}
}

type RegisterFlow struct {
	username string
	password string

	data   *LoginResponse
	userid int64
	token  string
}

func (r *RegisterFlow) Do() (*LoginResponse, error) {
	// 校验参数
	if err := r.checkParams(); err != nil {
		return nil, err
	}

	// 更新数据到数据库
	if err := r.updateData(); err != nil {
		return nil, err
	}

	// 打包数据，返回
	if err := r.packResponse(); err != nil {
		return nil, err
	}

	return r.data, nil
}

func (r *RegisterFlow) checkParams() error {
	if r.username == "" {
		return errors.New("用户名为空")
	}

	if len(r.username) > MaxUsernameLength {
		return errors.New("用户名长度超出限制")
	}

	if r.password == "" {
		return errors.New("密码为空")
	}

	return nil
}

func (r *RegisterFlow) updateData() error {
	userLogin := model.UserLogin{Username: r.username, Password: r.password}
	userInfo := model.UserInfo{User: userLogin, Username: r.username}

	// 判断用户是否存在
	userLoginDAO := model.NewUserLoginDao()
	if userLoginDAO.IsUserExistByUsername(r.username) {
		return errors.New("该用户名已存在")
	}

	// 更新操作
	userInfoDAO := model.NewUserInfoDAO()
	err := userInfoDAO.AddUserInfo(&userInfo)
	if err != nil {
		return err
	}

	// 颁发token
	token, err := middleware.ReleaseToken(userLogin)
	if err != nil {
		return err
	}

	r.token = token
	r.userid = int64(userInfo.ID)
	return nil
}

func (r *RegisterFlow) packResponse() error {
	r.data = &LoginResponse{
		UserId: r.userid,
		Token:  r.token,
	}

	return nil
}
