package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code       uint64      `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	ErrorField interface{} `json:"errorField,omitempty"`
	StatusCode int         `json:"statusCode"`
}

func HandlerSuccessResponse(c *gin.Context, response any) {
	c.JSON(http.StatusOK, OK(response))
}

func OK(i interface{}) *Response {
	return &Response{
		Code:       0,
		Message:    "success",
		Data:       i,
		StatusCode: 200,
	}
}

func Err(code uint64, statusCode int, message string) *Response {
	return &Response{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
	}
}
