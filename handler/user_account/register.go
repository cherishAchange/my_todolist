package user_account

import (
	"my_todolist/middleware"
	"my_todolist/service/user_account"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRegisterResponse struct {
	StatusCode int32  `json:"statusCode"`
	StatusMsg  string `json:"statusMsg,omitempty"`
	*user_account.LoginResponse
}

func UserRegisterHandler(c *gin.Context) {
	username := middleware.RegisterParams.Username
	password := middleware.RegisterParams.Password

	if password == "" {
		c.JSON(http.StatusOK, UserRegisterResponse{
			StatusCode: 1,
			StatusMsg:  "密码解析出错",
		})
		return
	}

	registerResponse, err := user_account.Register(username, password)

	if err != nil {
		c.JSON(http.StatusOK, UserRegisterResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	res := UserRegisterResponse{
		StatusCode:    0,
		StatusMsg:     "操作成功",
		LoginResponse: registerResponse,
	}

	c.JSON(http.StatusOK, res)
}
