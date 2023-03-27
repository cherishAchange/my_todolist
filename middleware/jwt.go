package middleware

import (
	"my_todolist/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_todolist")

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

func ReleaseToken(user model.UserLogin) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: int64(user.UserInfoID),
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
