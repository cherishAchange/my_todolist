package user_account

import (
	"my_todolist/middleware"
	"my_todolist/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseResponse struct {
	StatusCode int32  `json:"statusCode"`
	StatusMsg  string `json:"statusMsg,omitempty"`
}

func UserLoginHandler(c *gin.Context) {
	username := middleware.LoginParams.LoginUsername
	password := middleware.LoginParams.LoginPassword

	if username == "" || password == "" {
		c.JSON(http.StatusOK, BaseResponse{
			StatusCode: 1,
			StatusMsg:  "缺失用户名或密码",
		})
		return
	}

	userLogin := model.NewUserLoginDao()
	if !userLogin.IsUserExistByUsername(username) {
		c.JSON(http.StatusOK, BaseResponse{
			StatusCode: 1,
			StatusMsg:  "不存在该用户，请先注册",
		})
		return
	}

	user, err := userLogin.IsPasswordMatched(username, password)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, BaseResponse{
			StatusCode: 1,
			StatusMsg:  "请输入正确的密码",
		})
		return
	}

	token, err := middleware.ReleaseToken(user)
	if err == nil {
		c.SetCookie("todolist_token", token, 3600, "/", "127.0.0.1", false, true)
		c.JSON(http.StatusOK, BaseResponse{
			StatusCode: 0,
			StatusMsg:  "操作成功",
		})
		return
	}

	c.JSON(http.StatusOK, BaseResponse{
		StatusCode: 1,
		StatusMsg:  "登录失败",
	})
}
