package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"greebel.core.be/core"
	"greebel.core.be/helper"
	"greebel.core.be/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetClientJWTmiddlewares(g *echo.Group, role string) {
	env := core.App.Config

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(env.JWT_SECRET),
	}))

	g.Use(ValidateJWTLogin)
}

func ValidateJWTLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if claims["id"] != 0 {
				userBanned := models.UserBanned{}
				currentDate := fmt.Sprintf("%v", time.Now().Format("2006-01-02"))
				db := core.App.DB
				notFound := db.Where("user_id = ?", claims["id"]).
					Where("status = 1 OR (status = 2 AND end_date >= ?)", currentDate).
					First(&userBanned).
					RecordNotFound()

				if !notFound {
					return helper.Response(http.StatusForbidden, "", fmt.Sprintf("%s", "Your account has been banned"))
				}

				return next(c)
			} else {
				return helper.Response(http.StatusUnauthorized, "", fmt.Sprintf("%s", "Please Sign In \n Woops! Gonna sign in first\n Only a click away and you can continue to enjoy"))
			}
		}
		return helper.Response(http.StatusUnauthorized, "", fmt.Sprintf("%s", "Please Sign In \n Woops! Gonna sign in first\n Only a click away and you can continue to enjoy"))
	}
}
