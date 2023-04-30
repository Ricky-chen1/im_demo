package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID                 uint   `json:"id"`
	UserName           string `json:"user_name"`
	jwt.StandardClaims        // 实现了valid方法
}

var jwtKey = []byte("go_im_demo")
var expireTime = time.Minute * 60

func SignToken(id uint, userName string) (string, error) {
	claims := Claims{
		ID:       id,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireTime).Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString(jwtKey)
	if err != nil {
		LogInstance.Info("token signed error")
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (*Claims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		LogInstance.Info(err)
		return nil, err
	}
	claims, ok := jwtToken.Claims.(*Claims)
	if !ok || !jwtToken.Valid {
		LogInstance.Info("get claims fail")
		return nil, err
	}

	return claims, nil
}
