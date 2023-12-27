package helper

import (
	"github.com/labstack/echo/v4"
)

type (
	HttpResponse struct {
		Status  int                    `json:"status"`
		Message string                 `json:"message"`
		Data    map[string]interface{} `json:"data"`
	}
	HttpResponseData struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	MetaList struct {
		CurrentPage int `json:"current_page"`
		PerPage     int `json:"per_page"`
		Total       int `json:"total"`
		TotalPage   int `json:"total_page"`
	}

	HttpResponseWithMeta struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta"`
	}
)

func Response(statuscode int, data interface{}, message string) error {
	response := map[string]interface{}{
		"status":  statuscode,
		"message": message,
		"data":    data,
	}

	return echo.NewHTTPError(statuscode, response)
}
