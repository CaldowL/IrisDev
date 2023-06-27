package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecretToken = []byte("ly05661265")

type userClaim struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

// generateToken 根据userId生成token
func generateToken(user string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim{
		UserId: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 3)),
		},
	})
	s, _ := token.SignedString(jwtSecretToken)
	return s
}

// checkToken 检查token是否合法
func checkToken(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &userClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("ly05661265"), nil
	})
	if token.Valid {
		return true, nil
	} else {
		fmt.Println(err.Error())
		return false, err
	}
}

func run() {
	s := generateToken("a")
	fmt.Println(s)
	time.Sleep(1 * time.Second)
	fmt.Println(checkToken(s))
	time.Sleep(3 * time.Second)
	fmt.Println(checkToken(s))
}
