package model

import (
	"fmt"
	"my_todolist/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.DBConnectString()), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})

	fmt.Println("您的DB是：", DB)

	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&UserInfo{}, &UserLogin{}, &Task{})

	if err != nil {
		panic(err)
	}
}
