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
	baseGroup.POST("/task/create", middleware.JWTMiddleware(), tasks.CreateTasks)
	baseGroup.POST("/task/follow", middleware.JWTMiddleware(), tasks.FollowTask)
	baseGroup.POST("/query/getAllTasks", middleware.JWTMiddleware(), tasks.QueryAllTasks)
	baseGroup.POST("/query/getOwnTasks", middleware.JWTMiddleware(), tasks.QueryTasksByOwner)
	baseGroup.POST("/query/getFollowed", middleware.JWTMiddleware(), tasks.QueryTasksByFollower)
	baseGroup.POST("/query/getUsers", middleware.JWTMiddleware(), tasks.QueryUsersByTaskId)

	return r
}
