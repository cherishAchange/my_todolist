package router

import (
	"my_todolist/handler/user_account"
	"my_todolist/middleware"
	"my_todolist/model"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	model.InitDB()

	r := gin.Default()

	baseGroup := r.Group("api")

	baseGroup.POST("/user/register", middleware.SHAMiddleWare(), user_account.UserRegisterHandler)

	return r
}
