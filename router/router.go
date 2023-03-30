package router

import (
	"my_todolist/handler/tasks"
	"my_todolist/handler/user_account"
	"my_todolist/middleware"
	"my_todolist/model"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	model.InitDB()

	r := gin.Default()

	baseGroup := r.Group("api")

	baseGroup.POST("/user/register", middleware.BindParamsToCTX(), user_account.UserRegisterHandler)
	baseGroup.POST("/user/login", middleware.BindParamsToCTX(), user_account.UserLoginHandler)
	baseGroup.POST("/task/create", middleware.BindParamsToCTX(), middleware.JWTMiddleware(), tasks.CreateTasks)

	return r
}
