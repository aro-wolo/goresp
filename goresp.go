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

func (r *Responder) sendJSON(status int, defaultMsg string, isError bool, data interface{}, messages ...string) {
	message := defaultMsg
	if len(messages) > 0 {
		message = messages[0]
	}

	r.ctx.JSON(status, DataResponse{
		Message: message,
		Data:    data,
		IsError: isError,
	})
}

// Success responses
func (r *Responder) Ok(data interface{}, messages ...string) {
	r.sendJSON(http.StatusOK, "Success", false, data, messages...)
}

func (r *Responder) Created(data interface{}, messages ...string) {
	r.sendJSON(http.StatusCreated, "Resource created successfully", false, data, messages...)
}

// Error responses
func (r *Responder) BadRequest(messages ...string) {
	r.sendJSON(http.StatusBadRequest, "Bad Request", true, nil, messages...)
}

func (r *Responder) UnprocessableEntity(messages ...string) {
	r.sendJSON(http.StatusUnprocessableEntity, "Unprocessable Entity", true, nil, messages...)
}

func (r *Responder) ServerError(messages ...string) {
	r.sendJSON(http.StatusInternalServerError, "Internal Server Error", true, nil, messages...)
}

func (r *Responder) Error404(messages ...string) {
	r.sendJSON(http.StatusNotFound, "Not Found", true, nil, messages...)
}

func (r *Responder) NotFound(messages ...string) {
	r.sendJSON(http.StatusNotFound, "Not Found", true, nil, messages...)
}

func (r *Responder) Forbidden(messages ...string) {
	r.sendJSON(http.StatusForbidden, "Forbidden", true, nil, messages...)
}

func (r *Responder) Conflict(messages ...string) {
	r.sendJSON(http.StatusConflict, "Conflict", true, nil, messages...)
}

func (r *Responder) TooManyRequests(messages ...string) {
	r.sendJSON(http.StatusTooManyRequests, "Too Many Requests", true, nil, messages...)
}

func (r *Responder) AccessDenied(messages ...string) {
	r.sendJSON(http.StatusUnauthorized, "Access Denied", true, nil, messages...)
	r.ctx.Abort()
}

func (r *Responder) NotAllowed(messages ...string) {
	r.sendJSON(http.StatusMethodNotAllowed, "Method Not Allowed", true, nil, messages...)
}

func (r *Responder) NotAcceptable(messages ...string) {
	r.sendJSON(http.StatusNotAcceptable, "Not Acceptable", true, nil, messages...)
}

func (r *Responder) UnsupportedMedia(messages ...string) {
	r.sendJSON(http.StatusUnsupportedMediaType, "Unsupported Media Type", true, nil, messages...)
}

func (r *Responder) ReqTimeout(messages ...string) {
	r.sendJSON(http.StatusRequestTimeout, "Request Timeout", true, nil, messages...)
}

func (r *Responder) JSON(code int, message string, data interface{}, isError bool) {
	r.sendJSON(code, message, isError, data)
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
