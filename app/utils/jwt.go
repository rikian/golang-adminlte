package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtPassword = "S4n94t_R4h4S14_BRO..."
var SecretJwt = []byte(jwtPassword)

func EncryptToken(email string, expired int) (string, error) {
	secretJWT := SecretJwt
	client := jwt.MapClaims{
		"user_email": email,
		"exp":        time.Now().Add(time.Second * time.Duration(expired)).Unix(),
	}
	encrypt := jwt.NewWithClaims(jwt.SigningMethodHS256, client)
	token, err := encrypt.SignedString([]byte(secretJWT))
	if err != nil {
		return "", err
	}
	return token, nil
}

func DecryptToken(tokenStr string) (jwt.MapClaims, bool) {
	secretJWT := SecretJwt
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretJWT, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}
