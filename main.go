package main

import (
	"github.com/labstack/echo/v4"
	"injection.javamifi.com/controllers"
	"injection.javamifi.com/core"
	_ "injection.javamifi.com/docs"

	"os"

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

	var host string
	if core.App.Config.DB_NAME == "db_javamifi" {
		host = "127.0.0.1"
	}

	e.Logger.Fatal(e.Start(host + ":" + core.App.Port))
	os.Exit(0)
}
