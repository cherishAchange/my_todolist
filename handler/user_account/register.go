package user_account

import (
	"fmt"
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
	username := c.Query("username")
	rowVal, _ := c.Get("password")
	password, ok := rowVal.(string)

	fmt.Println("username", c.PostForm("username"))

	if !ok {
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
