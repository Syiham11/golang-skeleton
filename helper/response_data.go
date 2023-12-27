package helper

import (
	"github.com/labstack/echo/v4"
	"injection.javamifi.com/constants"
	"injection.javamifi.com/core"
)

type ErrorData struct {
	Message string      `json:"message"`
	Err     error       `json:"-"`
	Reason  interface{} `json:"reason"`
}

func (e ErrorData) Error() string {
	return e.Message
}

type PaginationData struct {
	TotalData   int `json:"total_data"`
	TotalPage   int `json:"total_page"`
	CurrentPage int `json:"current_page"`
	PageSize    int `json:"page_size"`
}

type MetaData struct {
	Pagination *PaginationData `json:"pagination"`
	Error      *ErrorData      `json:"errors"`
}

type ResponseData struct {
	Status   int         `json:"status_code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	MetaData *MetaData   `json:"meta_data"`
}

type ResponseDataView struct {
	Status     int         `json:"status_code"`
	StatusEdit bool        `json:"status_edit"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	MetaData   *MetaData   `json:"meta_data"`
}

func ReturnInvalidResponse(httpCode int, statusCode constants.MessageEnum, err error) error {
	var errData ErrorData
	if val, ok := err.(ErrorData); ok {
		errData = val
	} else {
		errData = ErrorData{
			Message: err.Error(),
			Err:     err,
		}
	}

	responseBody := ResponseData{
		Status:   statusCode.Int(),
		Message:  errData.Message,
		MetaData: &MetaData{Pagination: nil, Error: &errData},
	}
	return echo.NewHTTPError(httpCode, responseBody)
}

func ReturnResponse(httpCode int, statusCode constants.MessageEnum, c echo.Context, data interface{}) error {
	meta := MetaData{}
	var resData interface{}
	if val, ok := data.(core.PagedFindResult); ok {
		pagination := PaginationData{
			TotalData:   val.TotalData,
			PageSize:    val.Rows,
			CurrentPage: val.CurrentPage,
			TotalPage:   val.LastPage,
		}
		meta.Pagination = &pagination
		resData = val.Data
	} else {
		resData = data
	}

	response := ResponseData{
		Status:   statusCode.Int(),
		Message:  statusCode.String(),
		MetaData: &meta,
		Data:     resData,
	}

	return c.JSON(httpCode, response)
}

func ReturnResponseDetail(httpCode int, statusCode constants.MessageEnum, c echo.Context, statusEdit bool, data interface{}) error {
	meta := MetaData{}
	var resData interface{}
	if val, ok := data.(core.PagedFindResult); ok {
		pagination := PaginationData{
			TotalData:   val.TotalData,
			PageSize:    val.Rows,
			CurrentPage: val.CurrentPage,
			TotalPage:   val.LastPage,
		}
		meta.Pagination = &pagination
		resData = val.Data
	} else {
		resData = data
	}

	response := ResponseDataView{
		Status:     statusCode.Int(),
		StatusEdit: statusEdit,
		Message:    statusCode.String(),
		MetaData:   &meta,
		Data:       resData,
	}

	return c.JSON(httpCode, response)
}

func WrapError(errConst constants.MessageEnum, err error, reason interface{}) error {
	return ErrorData{
		Message: err.Error(),
		Err:     err,
		Reason:  reason,
	}
}
