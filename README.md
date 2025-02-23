# goresp

[![Go Reference](https://pkg.go.dev/badge/github.com/aro-wolo/goresp.svg)](https://pkg.go.dev/github.com/aro-wolo/goresp)
[![Go Report Card](https://goreportcard.com/badge/github.com/aro-wolo/goresp)](https://goreportcard.com/report/github.com/aro-wolo/goresp)
[![Go Test](https://github.com/aro-wolo/goresp/actions/workflows/test.yml/badge.svg)](https://github.com/aro-wolo/goresp/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/github/aro-wolo/goresp/badge.svg?branch=main)](https://coveralls.io/github/aro-wolo/goresp?branch=main)


**goresp** is a lightweight Go library for handling standardized API responses using Gin.

## Installation

To install goresp, use:

```sh
go get github.com/aro-wolo/goresp
```

Then import it in your project:

```go
import "github.com/aro-wolo/goresp"
```

## Usage

goresp provides functions to simplify API responses:

### 1. Success Response

Use `SuccessResponse` to send a standard success response:

```go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/aro-wolo/goresp"
)

func handler(c *gin.Context) {
	data := map[string]string{"message": "Hello, World!"}
	goresp.SuccessResponse(c, data, "Request successful")
}

func main() {
	r := gin.Default()
	r.GET("/", handler)
	r.Run(":8080")
}
```

### 2. Error Response

Use `ErrorResponse` to return an error response with a status code:

```go
func errorHandler(c *gin.Context) {
	goresp.ErrorResponse(c, "Invalid request", http.StatusBadRequest)
}
```

### 3. Custom Response

Use `JSONResp` for full control over the response:

```go
func customHandler(c *gin.Context) {
	data := map[string]string{"error": "Unauthorized access"}
	goresp.JSONResp(c, http.StatusUnauthorized, "Access denied", data, true)
}
```

### 4. ShouldBind with Automatic Error Handling

Use `ShouldBind` to automatically bind JSON request bodies and return a bad request response if binding fails:

```go
func createUserHandler(c *gin.Context) {
	var user struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := goresp.ShouldBind(c, &user); err != nil {
		return
	}

	goresp.SuccessResponse(c, user, "User created successfully")
}
```

## Response Format

All responses follow a standard format:

```json
{
  "msg": "Request successful",
  "data": { "message": "Hello, World!" },
  "err": false
}
```

## License

goresp is licensed under the MIT License.

