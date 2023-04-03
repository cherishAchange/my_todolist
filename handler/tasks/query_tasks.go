package tasks

import (
	"my_todolist/service/takes"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AllTasksResponse struct {
	StatusCode int         `json:"statusCode"`
	StatusMsg  string      `json:"statusMsg"`
	Data       interface{} `json:"data"`
}

func QueryAllTasks(ctx *gin.Context) {
	results, err := takes.QueryAllTasks()
	response := AllTasksResponse{}

	if err != nil {
		response.StatusCode = 1
		response.StatusMsg = err.Error()
		ctx.JSON(500, response)
		return
	}

	response.StatusCode = 0
	response.StatusMsg = "操作成功"
	response.Data = results
	ctx.JSON(http.StatusOK, response)
}

func QueryTasksByOwner(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	response := AllTasksResponse{}

	results, err := takes.QueryCreated(userId.(uint))

	if err != nil {
		response.StatusCode = 1
		response.StatusMsg = err.Error()
		ctx.JSON(500, response)
		return
	}

	response.StatusCode = 0
	response.StatusMsg = "操作成功"
	response.Data = results
	ctx.JSON(http.StatusOK, response)
}

func QueryTasksByFollower(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	response := AllTasksResponse{}

	results, err := takes.QueryFollowedTasks(userId.(uint))

	if err != nil {
		response.StatusCode = 1
		response.StatusMsg = err.Error()
		ctx.JSON(500, response)
		return
	}

	response.StatusCode = 0
	response.StatusMsg = "操作成功"
	response.Data = results
	ctx.JSON(http.StatusOK, response)
}

type QueryFollowerParam struct {
	TaskID uint `json:"taskId" binding:"required"`
}

func QueryUsersByTaskId(ctx *gin.Context) {
	var param *QueryFollowerParam
	err := ctx.BindJSON(&param)
	response := AllTasksResponse{}

	if err != nil {
		response.StatusCode = 1
		response.StatusMsg = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	results, queryErr := takes.QueryFollower(param.TaskID)
	if queryErr != nil {
		response.StatusCode = 1
		response.StatusMsg = queryErr.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response.StatusCode = 0
	response.StatusMsg = "操作成功"
	response.Data = results
	ctx.JSON(http.StatusOK, response)
}
