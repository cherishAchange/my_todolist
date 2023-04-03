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

func (u *UserLoginDAO) IsUserExistByUsername(username string) (bool, error) {
	var userLogin UserLogin
	err := DB.Where("username=?", username).First(&userLogin).Error

	return userLogin.ID != 0, err
}

func (u *UserLoginDAO) IsPasswordMatched(username string, password string) (UserLogin, error) {
	var userLogin UserLogin
	err := DB.Where("username=? and password=?", username, password).First(&userLogin).Error

	return userLogin, err
}
