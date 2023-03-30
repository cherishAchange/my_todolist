package user_account

import (
	"my_todolist/middleware"
	"my_todolist/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserLoginHandler(c *gin.Context) {
	username := middleware.LoginParams.LoginUsername
	password := middleware.LoginParams.LoginPassword

	if username == "" || password == "" {
		c.JSON(http.StatusOK, middleware.GetRes(1, "缺失用户名或密码"))
		return
	}

	userLogin := model.NewUserLoginDao()
	if !userLogin.IsUserExistByUsername(username) {
		c.JSON(http.StatusOK, middleware.GetRes(1, "不存在该用户，请先注册"))
		return
	}

	user, err := userLogin.IsPasswordMatched(username, password)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, middleware.GetRes(1, "请输入正确的密码"))
		return
	}

	token, err := middleware.ReleaseToken(user)
	if err == nil {
		c.SetCookie(middleware.TOKENNAME, token, 3600, "/", "127.0.0.1", false, true)
		c.JSON(http.StatusOK, middleware.GetRes(1, "操作成功"))
		return
	}

	c.JSON(http.StatusOK, middleware.GetRes(1, "登录失败"))
}
