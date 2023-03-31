package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type RegisterParamsStruct struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginParamsStruct struct {
	LoginUsername string `json:"loginUsername" binding:"required"`
	LoginPassword string `json:"loginPassword" binding:"required"`
}

var RegisterParams RegisterParamsStruct
var LoginParams LoginParamsStruct

func BindParamsToCTX() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		if err = ctx.ShouldBindBodyWith(&RegisterParams, binding.JSON); err == nil {
			fmt.Println("绑定1")
			RegisterParams.Password = SHA1ForPassword(RegisterParams.Password)
			ctx.Next()
		} else if err = ctx.ShouldBindBodyWith(&LoginParams, binding.JSON); err == nil {
			fmt.Println("绑定2")
			LoginParams.LoginPassword = SHA1ForPassword(LoginParams.LoginPassword)
			ctx.Next()
		}

		fmt.Println("err", err)
		ctx.Next()
	}
}
