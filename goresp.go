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

func (r *Responder) Ok(data interface{}, message string) {
	r.ctx.JSON(http.StatusOK, DataResponse{
		Message: message,
		Data:    data,
		IsError: false,
	})
}

func (r *Responder) Created(data interface{}, message string) {
	r.ctx.JSON(http.StatusCreated, DataResponse{
		Message: message,
		Data:    data,
		IsError: false,
	})
}


func (r *Responder) ServerError(message string) {
	r.ctx.JSON(http.StatusInternalServerError, DataResponse{
		Message: message,
		IsError: true,
	})
}

func (r *Responder) BadRequest(message string) {
	r.ctx.JSON(http.StatusBadRequest, DataResponse{
		Message: message,
		IsError: true,
	})
}

func (r *Responder) Error404(message string) {
	r.ctx.JSON(http.StatusNotFound, DataResponse{
		Message: message,
		IsError: true,
	})
}

func (r *Responder) AccessDenied(message string) {
	r.ctx.JSON(http.StatusUnauthorized, DataResponse{
		Message: message,
		IsError: true,
	})
}

func (r *Responder) Forbidden(message string) {
	r.ctx.JSON(http.StatusForbidden, DataResponse{
		Message: message,
		IsError: true,
	})
}

func (r *Responder) Conflict(message string) {
	r.ctx.JSON(http.StatusConflict, DataResponse{
		Message: message,
		IsError: true,
	})
}

func (r *Responder) UnprocessableEntity(message string) {
	r.ctx.JSON(http.StatusUnprocessableEntity, DataResponse{
		Message: message,
		IsError: true,
	})
}

func (r *Responder) TooManyRequests(message string) {
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

func (r *Responder) ShouldBind(obj interface{}) bool {
	if err := r.ctx.ShouldBindJSON(obj); err != nil {
		r.BadRequest("Invalid request payload")
		return false
	}
	return true
}
