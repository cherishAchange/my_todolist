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
	DB.Where("username=?", username).First(&userLogin)

	return userLogin.ID != 0
}
