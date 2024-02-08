package main

import (
	"greebel.core.be/controllers"
	"greebel.core.be/core"
	_ "greebel.core.be/docs"
	"greebel.core.be/middlewares"

	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	defer core.App.Close()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.Static("/", "tests/experiments")
	e.Static("/assets", "assets")

	e.Pre(middleware.Rewrite(map[string]string{
		"/api/*": "/$1",
	}))

	e.GET("/docs/*", echoSwagger.WrapHandler)

	e.GET("/healthcheck", controllers.HealthCheck)

	api := e.Group("/v1")
	{
		auth := api.Group("/auth")
		auth.POST("/register-user", controllers.RegisterUser)
		auth.POST("/registeruser", controllers.RegisterUserNoOtp)
		auth.POST("/verify-otp", controllers.VerifyOTP)
		auth.POST("/login", controllers.Login)
		auth.POST("/google-login", controllers.GoogleLogin)
		auth.POST("/google-login-x", controllers.GoogleLoginX)
		auth.GET("/google-login-callback", controllers.GoogleLoginCallback)
		auth.POST("/facebook-login-x", controllers.FacebookLoginX)
		auth.POST("/facebook-login", controllers.FacebookLogin)
		auth.GET("/facebook-login-callback", controllers.FacebookLoginCallback)
		auth.POST("/upload-image", controllers.UploadImage)
		auth.DELETE("/delete-image", controllers.DeleteImage)
		auth.POST("/request-otp", controllers.RequestOTPNoAuth)
		auth.POST("/forgot-password", controllers.ForgotPassword)
		auth.POST("/apple-login", controllers.AppleLogin)

		users := api.Group("/user")
		middlewares.SetClientJWTmiddlewares(users, "user")
		users.GET("/my-profile", controllers.MyProfile)
		users.POST("/request-otp", controllers.RequestOTP)
		users.POST("/upload-profile-photo", controllers.UploadProfilePhoto)
		users.PATCH("/edit-user-profile", controllers.EditUserProfile)
		users.PATCH("/change-password", controllers.ChangePassword)

		usersBlock := api.Group("/user/block")
		middlewares.SetClientJWTmiddlewares(usersBlock, "user")
		usersBlock.POST("/:user_id", controllers.BlockUser)

	}

	var host string
	if core.App.Config.DB_NAME == "db_eventori" {
		host = "127.0.0.1"
	}

	e.Logger.Fatal(e.Start(host + ":" + core.App.Port))
	os.Exit(0)
}
