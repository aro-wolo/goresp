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

type Responder struct {
	ctx *gin.Context
}

func New(c *gin.Context) *Responder {
	return &Responder{ctx: c}
}

func (r *Responder) Ok(data interface{}, messages ...string) {
	message := ""

	if len(messages) > 0 {
		message = messages[0]
	}

	r.ctx.JSON(http.StatusOK, DataResponse{
		Message: message,
		Data:    data,
		IsError: false,
	})
}

func (r *Responder) Created(data interface{}, messages ...string) {
	message := "Resource created successfully"

	if len(messages) > 0 {
		message = messages[0]
	}

	r.ctx.JSON(http.StatusCreated, DataResponse{
		Message: message,
		Data:    data,
		IsError: false,
	})
}

func (r *Responder) ServerError(messages ...string) {
	message := "Internal Server Error" // Default message

	if len(messages) > 0 {
		message = messages[0]
	}

	r.ctx.JSON(http.StatusInternalServerError, DataResponse{
		Message: message,
		IsError: true,
	})
}

func (r *Responder) BadRequest(messages ...string) {
	message := "Bad Request" // Default message

	if len(messages) > 0 {
		message = messages[0]
	}

	r.ctx.JSON(http.StatusBadRequest, DataResponse{
		Message: message,
		IsError: true,
	})
}

func (r *Responder) Error404(messages ...string) {
	message := "Not Found" // Default message

	if len(messages) > 0 {
		message = messages[0]
	}

	r.ctx.JSON(http.StatusNotFound, DataResponse{
		Message: message,
		IsError: true,
	})
}

func (r *Responder) AccessDenied(messages ...string) {
	message := "Access denied"

	if len(messages) > 0 {
		message = messages[0]
	}

	r.ctx.JSON(http.StatusUnauthorized, DataResponse{
		Message: message,
		IsError: true,
	})

	r.ctx.Abort()
}

func (r *Responder) Forbidden(messages ...string) {
	message := "Forbidden"

	if len(messages) > 0 {
		message = messages[0]
	}

	r.ctx.JSON(http.StatusForbidden, DataResponse{
		Message: message,
		IsError: true,
	})
}

func (r *Responder) Conflict(messages ...string) {
	message := "Conflict" // Default message

	if len(messages) > 0 {
		message = messages[0]
	}

	r.ctx.JSON(http.StatusConflict, DataResponse{
		Message: message,
		IsError: true,
	})
}

func (r *Responder) UnprocessableEntity(messages ...string) {
	message := "Unprocessable Entity" // Default message

	if len(messages) > 0 {
		message = messages[0]
	}

	r.ctx.JSON(http.StatusUnprocessableEntity, DataResponse{
		Message: message,
		IsError: true,
	})
}

func (r *Responder) TooManyRequests(messages ...string) {
	message := "Too Many Requests" // Default message

	if len(messages) > 0 {
		message = messages[0]
	}

	r.ctx.JSON(http.StatusTooManyRequests, DataResponse{
		Message: message,
		IsError: true,
	})
}

func (r *Responder) JSON(code int, message string, data interface{}, isError bool) {
	r.ctx.JSON(code, DataResponse{
		Message: message,
		Data:    data,
		IsError: isError,
	})
}

func (r *Responder) ShouldBind(obj interface{}, customErr ...string) bool {
	if err := r.ctx.ShouldBindJSON(obj); err != nil {
		message := "Invalid request payload: " + err.Error()

		if len(customErr) > 0 {
			message = customErr[0]
		}

		r.BadRequest(message)
		return false
	}
	return true
}
