package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GennerateJWT(userId string, email string, exp int64, secret string) (string, error) {
	jwtClaims := jwt.MapClaims{}
	jwtClaims["uuId"] = userId
	jwtClaims["email"] = email
	jwtClaims["exp"] = exp

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	tokenString, err := token.SignedString([]byte(secret))
	fmt.Println(err)

	return tokenString, err
}
