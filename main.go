package goresp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DataResponse struct {
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
	IsError bool        `json:"err"`
}

func OkResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, DataResponse{
		Message: message,
		Data:    data,
		IsError: false,
	})
}

func ServerErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, DataResponse{
		Message: message,
		IsError: true,
	})
}

func BadRequestResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, DataResponse{
		Message: message,
		IsError: true,
	})
}

func Error404Response(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, DataResponse{
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

// ShouldBindJSON is a helper function that binds JSON and returns an error response if binding fails
func ShouldBindJSON(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		BadRequestResponse(c, "Invalid request payload")
		return false
	}
	return true
}
