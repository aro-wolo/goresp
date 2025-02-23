package goresp

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestOkResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	OkResponse(c, "test data", "Success")

	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `{"msg":"Success","data":"test data","err":false}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestServerErrorResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	ServerErrorResponse(c, "Internal server error")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	expectedBody := `{"msg":"Internal server error","err":true}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestBadRequestResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	BadRequestResponse(c, "Bad request")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	expectedBody := `{"msg":"Bad request","err":true}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestError404Response(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	Error404Response(c, "Resource not found")

	assert.Equal(t, http.StatusNotFound, w.Code)
	expectedBody := `{"msg":"Resource not found","err":true}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestAccessDeniedResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	AccessDeniedResponse(c, "Access denied")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	expectedBody := `{"msg":"Access denied","err":true}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestJSONResp(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	JSONResp(c, http.StatusAccepted, "Accepted", "test data", false)

	assert.Equal(t, http.StatusAccepted, w.Code)
	expectedBody := `{"msg":"Accepted","data":"test data","err":false}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestShouldBindJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Test with valid JSON
	c.Request = httptest.NewRequest("POST", "/", nil)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Body = io.NopCloser(strings.NewReader(`{"key":"value"}`))

	var obj map[string]interface{}
	result := ShouldBindJSON(c, &obj)
	assert.True(t, result)

	// Test with invalid JSON
	c.Request = httptest.NewRequest("POST", "/", nil)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Body = io.NopCloser(strings.NewReader(`invalid json`))

	result = ShouldBindJSON(c, &obj)
	assert.False(t, result)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	expectedBody := `{"msg":"Invalid request payload","err":true}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
