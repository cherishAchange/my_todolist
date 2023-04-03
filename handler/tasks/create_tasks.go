package tasks

import (
	"my_todolist/service/takes"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateTaskParams struct {
	Title    string `json:"title" binding:"required"`
	Describe string `json:"describe" binding:"required"`
}

func CreateTasks(c *gin.Context) {
	createTaskParams := CreateTaskParams{}
	err := c.BindJSON(&createTaskParams)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userId, _ := c.Get("userId")

	response, createErr := takes.NewCreateTaskFlow(createTaskParams.Title, createTaskParams.Describe, userId.(uint))

	if createErr != nil {
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
