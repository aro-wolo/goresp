package goresp_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aro-wolo/goresp"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestResponder(t *testing.T) {
	gin.SetMode(gin.TestMode)

	responses := []struct {
		name     string
		method   func(*goresp.Responder)
		expected int
	}{
		{"Ok response", func(res *goresp.Responder) { res.Ok("success", "Operation successful") }, http.StatusOK},
		{"Created response", func(res *goresp.Responder) { res.Created("created", "Resource created") }, http.StatusCreated},
		{"Bad Request response", func(res *goresp.Responder) { res.BadRequest("Invalid input") }, http.StatusBadRequest},
		{"Server Error response", func(res *goresp.Responder) { res.ServerError("Internal error") }, http.StatusInternalServerError},
		{"Unauthorized response", func(res *goresp.Responder) { res.AccessDenied("Unauthorized access") }, http.StatusUnauthorized},
		{"Forbidden response", func(res *goresp.Responder) { res.Forbidden("Forbidden action") }, http.StatusForbidden},
		{"Not Found response", func(res *goresp.Responder) { res.Error404("Resource not found") }, http.StatusNotFound},
		{"Conflict response", func(res *goresp.Responder) { res.Conflict("Conflict detected") }, http.StatusConflict},
		{"Unprocessable Entity response", func(res *goresp.Responder) { res.UnprocessableEntity("Unprocessable entity") }, http.StatusUnprocessableEntity},
		{"Too Many Requests response", func(res *goresp.Responder) { res.TooManyRequests("Rate limit exceeded") }, http.StatusTooManyRequests},
	}

	for _, tt := range responses {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			res := goresp.New(c)
			tt.method(res)
			assert.Equal(t, tt.expected, w.Code)
		})
	}

	t.Run("JSON response", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		res := goresp.New(c)
		res.JSON(http.StatusTeapot, "I'm a teapot", "data", true)
		assert.Equal(t, http.StatusTeapot, w.Code)
	})

	t.Run("ShouldBind success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"John"}`))
		c.Request.Header.Set("Content-Type", "application/json")

		type RequestBody struct {
			Name string `json:"name"`
		}
		var body RequestBody
		res := goresp.New(c)
		bound := res.ShouldBind(&body)
		assert.True(t, bound)
		assert.Equal(t, "John", body.Name)
	})

	t.Run("ShouldBind failure", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`invalid json`))
		c.Request.Header.Set("Content-Type", "application/json")

		type RequestBody struct {
			Name string `json:"name"`
		}
		var body RequestBody
		res := goresp.New(c)
		bound := res.ShouldBind(&body)
		assert.False(t, bound)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
