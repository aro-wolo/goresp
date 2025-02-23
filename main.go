package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DataResponse struct {
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
	IsError bool        `json:"err"`
}

func SuccessResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, DataResponse{
		Message: message,
		Data:    data,
		IsError: false,
	})
}

func ErrorResponse(c *gin.Context, message string, code int) {
	c.JSON(code, DataResponse{
		Message: message,
		IsError: true,
	})
}

func JSONResp(c *gin.Context, code int, message string, data interface{}, isError bool) {
	c.JSON(code, DataResponse{
		Message: message,
		Data:    data,
		IsError: isError,
	})
}
