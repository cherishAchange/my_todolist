package model

import (
	"errors"
	"sync"

	"gorm.io/gorm"
)

var (
	ErrIvdPtr        = errors.New("空指针错误")
	ErrEmptyUserList = errors.New("用户列表为空")
)

type UserInfo struct {
	gorm.Model
	Username string    `json:"name" gorm:"name,omitempty"`
	User     UserLogin `json:"-"`
	Tasks    []Task    `json:"tasks"`                                       // 用户与task [1对多]：一个用户可以创建多个task，一个task不能被多个用户创建
	Follows  []*Task   `json:"follows" gorm:"many2many:user_follow_tasks;"` // 用户与task [多对多]：一个用户可以follow多个task，一个task也可以被多个用户follow
}

type UserInfoDAO struct {
}

var (
	userInfoDao  *UserInfoDAO
	userInfoOnce sync.Once
)

func NewUserInfoDAO() *UserInfoDAO {
	userInfoOnce.Do(func() {
		userInfoDao = new(UserInfoDAO)
	})

	return userInfoDao
}

func (u *UserInfoDAO) AddUserInfo(userInfo *UserInfo) error {
	if userInfo == nil {
		return ErrIvdPtr
	}

	return DB.Create(userInfo).Error
}
