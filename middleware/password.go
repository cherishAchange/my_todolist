package middleware

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/gin-gonic/gin"
)

type RegisterParams struct {
	Username string `json:"username" xml:"username" binding:"required"`
	Password string `json:"password" xml:"password" binding:"required"`
}

func SHA1(s string) string {
	o := sha1.New()

	o.Write([]byte(s))

	return hex.EncodeToString(o.Sum(nil))
}

func SHAMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := RegisterParams{}
		ctx.ShouldBind(&params)
		fmt.Println("再看看", params.Username, params.Password)

		password := ctx.Query("password")
		if password == "" {
			password = ctx.PostForm("password")
		}
		fmt.Println("中间件：", password)
		ctx.Set("password", SHA1(password))
		ctx.Next()
	}
}
