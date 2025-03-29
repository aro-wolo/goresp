package goresp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// setupTest initializes a Gin context and returns a Responder instance.
func setupTest() (*gin.Context, *httptest.ResponseRecorder, *Responder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	res := New(c)
	return c, w, res
}

// Helper function to extract JSON response
func getJSONResponse(w *httptest.ResponseRecorder) DataResponse {
	var response DataResponse
	_ = json.NewDecoder(w.Body).Decode(&response)
	return response
}

// Test New Responder
func TestNewResponder(t *testing.T) {
	c, _, res := setupTest()
	assert.NotNil(t, res)
	assert.Equal(t, c, res.ctx)
}

// Test Ok Response
func TestResponder_Ok(t *testing.T) {
	_, w, res := setupTest()

	res.Ok(map[string]string{"key": "value"}, "Success Message")

	assert.Equal(t, http.StatusOK, w.Code)
}

// Test Created Response
func TestResponder_Created(t *testing.T) {
	_, w, res := setupTest()

	res.Created(map[string]string{"id": "123"}, "Created Successfully")

	assert.Equal(t, http.StatusCreated, w.Code)
}

// Test BadRequest Response
func TestResponder_BadRequest(t *testing.T) {
	_, w, res := setupTest()

	res.BadRequest("Invalid input")

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Test UnprocessableEntity Response
func TestResponder_UnprocessableEntity(t *testing.T) {
	_, w, res := setupTest()

	res.UnprocessableEntity("Cannot process request")

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

// Test ServerError Response
func TestResponder_ServerError(t *testing.T) {
	_, w, res := setupTest()

	res.ServerError("Something went wrong")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// Test NotFound Response
func TestResponder_NotFound(t *testing.T) {
	_, w, res := setupTest()

	res.NotFound("Resource not found")

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// Test Forbidden Response
func TestResponder_Forbidden(t *testing.T) {
	_, w, res := setupTest()

	res.Forbidden("You are not allowed")

	assert.Equal(t, http.StatusForbidden, w.Code)
}

// Test Conflict Response
func TestResponder_Conflict(t *testing.T) {
	_, w, res := setupTest()

	res.Conflict("Resource conflict")

	assert.Equal(t, http.StatusConflict, w.Code)
}

// Test TooManyRequests Response
func TestResponder_TooManyRequests(t *testing.T) {
	_, w, res := setupTest()

	res.TooManyRequests("Too many requests made")

	assert.Equal(t, http.StatusTooManyRequests, w.Code)
}

// Test AccessDenied Response
func TestResponder_AccessDenied(t *testing.T) {
	_, w, res := setupTest()

	res.AccessDenied("Unauthorized Access")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// Test Method Not Allowed Response
func TestResponder_NotAllowed(t *testing.T) {
	_, w, res := setupTest()

	res.NotAllowed("Method Not Allowed")

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
}

// Test Not Acceptable Response
func TestResponder_NotAcceptable(t *testing.T) {
	_, w, res := setupTest()

	res.NotAcceptable("Not Acceptable Format")

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

// Test Unsupported Media Type Response
func TestResponder_UnsupportedMedia(t *testing.T) {
	_, w, res := setupTest()

	res.UnsupportedMedia("Unsupported Media Type")

	assert.Equal(t, http.StatusUnsupportedMediaType, w.Code)
}

// Test Request Timeout Response
func TestResponder_ReqTimeout(t *testing.T) {
	_, w, res := setupTest()

	res.ReqTimeout("Request timed out")

	assert.Equal(t, http.StatusRequestTimeout, w.Code)
}

// Test Custom JSON Response
func TestResponder_JSON(t *testing.T) {
	_, w, res := setupTest()

	res.JSON(http.StatusTeapot, "I'm a teapot", nil, false)

	assert.Equal(t, http.StatusTeapot, w.Code)
}

// Test ShouldBind with Valid JSON
func TestResponder_ShouldBind_Success(t *testing.T) {
	_, w, res := setupTest()

	res.ctx.Request = httptest.NewRequest(http.MethodPost, "/", nil)
	res.ctx.Request.Header.Set("Content-Type", "application/json")

	var obj map[string]string
	success := res.ShouldBind(&obj)

	assert.True(t, success)
	assert.Equal(t, http.StatusOK, w.Code)
}

// Test ShouldBind with Invalid JSON
func TestResponder_ShouldBind_Fail(t *testing.T) {
	_, w, res := setupTest()

	res.ctx.Request = httptest.NewRequest(http.MethodPost, "/", nil)
	res.ctx.Request.Header.Set("Content-Type", "application/json")

	var obj map[string]string
	success := res.ShouldBind(&obj, "Custom Error Message")

	assert.False(t, success)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
