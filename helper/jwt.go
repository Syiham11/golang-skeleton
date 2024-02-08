package helper

import (
	"greebel.core.be/core"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type (
	JWTclaims struct {
		ID        int `json:"id"`
		IsPartner int `json:"is_partner"`
		jwt.StandardClaims
	}
	JWTclaimsTicket struct {
		ID      int    `json:"id"`
		OrderID string `json:"order_id"`
		jwt.StandardClaims
	}
)

func CreateJwtToken(userID int, isPartner int) (string, error) {
	claim := JWTclaims{
		userID,
		isPartner,
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

func CreateJwtTokenPDF(userID int, OrderID string) (string, error) {
	expireToken := time.Now().Add(time.Hour * 12).Unix()
	claim := JWTclaimsTicket{
		userID,
		OrderID,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Id:        strconv.Itoa(userID),
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
