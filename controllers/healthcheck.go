package controllers

import (
	"fmt"
	"greebel.core.be/core"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Info struct {
	Time string `json:"time"`
	DB   bool   `json:"database"`
}

var (
	err  error
	info Info
)

// HealthCheck check API health status
// @Summary Check API health status
// @Description Check API health status
// @Tags healthcheck
// @ID healthcheck-healthcheck
// @Accept json
// @Produce application/json
// @Success 200 {object} interface{}
// @Router /healthcheck [get]
func HealthCheck(c echo.Context) error {
	defer c.Request().Body.Close()

	info.Time = fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05"))
	info.DB = true

	if err = core.App.DB.DB().Ping(); err != nil {
		info.DB = false
	}

	return c.JSON(http.StatusOK, info)
}
