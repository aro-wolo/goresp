# goresp Go API Response 

![GitHub Repo stars](https://img.shields.io/github/stars/aro-wolo/goresp?style=social)
![GitHub last commit](https://img.shields.io/github/last-commit/aro-wolo/goresp)
![GitHub issues](https://img.shields.io/github/issues/aro-wolo/goresp)
![GitHub contributors](https://img.shields.io/github/contributors/aro-wolo/goresp)
[![Go Report Card](https://goreportcard.com/badge/github.com/aro-wolo/goresp)](https://goreportcard.com/report/github.com/aro-wolo/goresp)
[![Build Status](https://github.com/aro-wolo/goresp/actions/workflows/go.yml/badge.svg)](https://github.com/aro-wolo/goresp/actions)
[![codecov](https://codecov.io/gh/aro-wolo/goresp/branch/main/graph/badge.svg)](https://codecov.io/gh/aro-wolo/goresp)
![Go Module](https://img.shields.io/github/go-mod/go-version/aro-wolo/goresp)
[![License](https://img.shields.io/github/license/aro-wolo/goresp.svg)](https://github.com/aro-wolo/goresp/blob/main/LICENSE)
[![wakatime](https://wakatime.com/badge/user/c78c31fe-9c97-4b21-b865-91bc182f2d42.svg)](https://wakatime.com/@c78c31fe-9c97-4b21-b865-91bc182f2d42)

**goresp** is a lightweight Go library for handling standardized API responses using the Gin framework.

## Installation

Install `goresp` using:

```sh
go get github.com/aro-wolo/goresp
```

Then import it into your project:

```go
import "github.com/aro-wolo/goresp"
```

## Features

- Standardized JSON response format
- Predefined response methods for common HTTP statuses
- Easy JSON request binding with validation handling
- Customizable responses with flexible JSON formatting

## Usage

goresp simplifies API response handling with predefined functions.

### 1. Standard Success Response

Use `Ok` to send a successful response:

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/aro-wolo/goresp"
)

func handler(c *gin.Context) {
	res := goresp.New(c)
	data := map[string]string{"message": "Hello, World!"}
	res.Ok(data, "Request successful")
}

func main() {
	r := gin.Default()
	r.GET("/", handler)
	r.Run(":8080")
}
```

### 2. Standard Error Responses

#### Bad Request
```go
res.BadRequest("Invalid request data")
```

#### Unauthorized
```go
res.AccessDenied("Unauthorized access")
```

#### Not Found
```go
res.Error404("Resource not found")
```

#### Internal Server Error
```go
res.ServerError("Internal server error")
```

### 3. Custom Response

Use `JSON` for full control over the response:

```go
res.JSON(418, "I'm a teapot", nil, true)
```

### 4. Request Binding with Automatic Error Handling

Use `ShouldBind` to automatically bind JSON request bodies and handle errors:

```go
func createUserHandler(c *gin.Context) {
	var user struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	res := goresp.New(c)
	if !res.ShouldBind(&user) {
		return
	}

	res.Ok(user, "User created successfully")
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

