package middleware

import (
	"my_todolist/model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_todolist")

const TOKENNAME string = "todolist_token"

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.UserLogin) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.UserInfoID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "my_todolist",
			Subject:   "L_B__",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, bool) {
	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if token != nil {
		if key, ok := token.Claims.(*Claims); ok {
			if token.Valid {
				return key, true
			} else {
				return key, false
			}
		}
	}

	return nil, false
}

// 鉴权
func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr, err := ctx.Cookie(TOKENNAME)

		if tokenStr == "" || err != nil {
			ctx.JSON(http.StatusOK, GetRes(401, "请先登录"))
			ctx.Abort()
			return
		}

		token, ok := ParseToken(tokenStr)
		if !ok {
			ctx.JSON(http.StatusOK, GetRes(403, "鉴权失败"))
			ctx.Abort()
			return
		}

		if time.Now().Unix() > token.ExpiresAt {
			ctx.JSON(402, "token已失效")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
