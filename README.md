# Response

A simple and consistent HTTP response builder for Go.

## Installation

```bash
go get github.com/hiroaki-th/response
```

## Features

- **Method Chaining** - Fluent API for easy response building
- **Consistent Format** - Standardized JSON response structure
- **Simple API** - Minimal setup, maximum productivity

## Quick Start

```go
package main

import (
    "net/http"
    "github.com/hiroaki-th/response"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Success response
    response.Success(w).
        Json().
        BasicSecurity().
        SetBody(map[string]string{"message": "Hello World"}).
        Return()
}
```

## API Reference

### Status Code Helpers

```go
response.Success(w)      // 200 OK
response.Accepted(w)     // 202 Accepted  
response.BadRequest(w)   // 400 Bad Request
response.ServerError(w)  // 500 Internal Server Error
```

### Method Chaining

```go
response.Success(w).
    Json().                                    // Set Content-Type
    BasicSecurity().                          // Add security headers
    SetBody(data).                           // Set response data
    SetError(err).                           // Set error (if any)
    SetHeader("key", "value").               // Add custom header
    SetHeaders(map[string]string{...}).      // Add multiple headers
    Return()                                 // Send response
```

## Response Structure

```json
{
  "status": 200,
  "data": {...},
  "error": {
    "message": "error message"
  },
  "timestamp": "2025-06-05T10:30:00Z"
}
```

## Examples

### Success Response

```go
func getUser(w http.ResponseWriter, r *http.Request) {
    user := User{Name: "John", Email: "john@example.com"}
    
    response.Success(w).
        Json().
        BasicSecurity().
        SetBody(user).
        Return()
}
```

### Error Response

```go
func loginHandler(w http.ResponseWriter, r *http.Request) {
    if err := validateCredentials(); err != nil {
        response.BadRequest(w).
            Json().
            BasicSecurity().
            SetError(err).
            Return()
        return
    }
    // ... success logic
}
```

### Custom Headers

```go
response.Success(w).
    Json().
    SetHeaders(map[string]string{
        "X-API-Version": "v1.0",
        "X-Rate-Limit": "1000",
    }).
    SetBody(data).
    Return()
```
