package tasks

import (
	"my_todolist/service/takes"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FolowTaskParam struct {
	TaskId uint `json:"taskId" binding:"required"`
}

func FollowTask(ctx *gin.Context) {
	followTaskParam := FolowTaskParam{}
	err := ctx.BindJSON(&followTaskParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userId, _ := ctx.Get("userId")

	response, serviceErr := takes.Follow(userId.(uint), followTaskParam.TaskId)

	if serviceErr != nil {
		ctx.JSON(500, response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
