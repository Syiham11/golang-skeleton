package helper

import (
	"strconv"

	"injection.javamifi.com/core"

	"github.com/golang-jwt/jwt"
)

type (
	JWTclaims struct {
		ID   int `json:"id"`
		Tipe int `json:"tipe"`
		jwt.StandardClaims
	}
)

func CreateJwtToken(userID int, tipe int) (string, error) {
	claim := JWTclaims{
		userID,
		tipe,
		jwt.StandardClaims{
			Id: strconv.Itoa(userID),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := rawToken.SignedString([]byte(core.App.Config.JWT_SECRET))
	if err != nil {
		return "", err
	}

	return token, nil
}

func DecodeJwtToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(core.App.Config.JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}
