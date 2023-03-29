package model

import (
	"sync"

	"gorm.io/gorm"
)

type UserLogin struct {
	gorm.Model
	UserInfoID uint
	Username   string `gorm:"primary_key"`
	Password   string `gorm:"size:200;notnull"`
}

type UserLoginDAO struct {
}

var (
	userLoginDao  *UserLoginDAO
	userLoginOnce sync.Once
)

func NewUserLoginDao() *UserLoginDAO {
	userLoginOnce.Do(func() {
		userLoginDao = new(UserLoginDAO)
	})

	return userLoginDao
}

func (u *UserLoginDAO) IsUserExistByUsername(username string) bool {
	var userLogin UserLogin
	err := DB.Where("username=?", username).First(&userLogin).Error

	// 如果此处错误不是ErrRecordNotFound，则需要中端流程
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}
	return userLogin.ID != 0
}

func (u *UserLoginDAO) IsPasswordMatched(username string, password string) (UserLogin, error) {
	var userLogin UserLogin
	err := DB.Where("username=? and password=?", username, password).First(&userLogin).Error

	return userLogin, err
}
