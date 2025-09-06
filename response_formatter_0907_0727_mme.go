// 代码生成时间: 2025-09-07 07:27:20
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// ResponseData is a struct to format API responses
type ResponseData struct {
    Data  interface{} `json:"data"`  // Data field for the response payload
    Error string      `json:"error"` // Error field for any error message
}

// ErrorResponse represents an error response
func ErrorResponse(c *gin.Context, httpStatus, errCode int, message string) {
    c.JSON(httpStatus, ResponseData{
        Data:  nil,
        Error: message,
    })
}

// SuccessResponse represents a successful response
func SuccessResponse(c *gin.Context, httpStatus int, payload interface{}) {
    c.JSON(httpStatus, ResponseData{
        Data:  payload,
        Error: "",
    })
}

// Middleware to handle errors and format responses
func ErrorHandlingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next() // proceed to the next middleware/handler
        
        // Check if there is any error stored in the context
        if len(c.Errors) > 0 {
            // If there is an error, return it in the response
            err := c.Errors.Last().Err
            ErrorResponse(c, http.StatusInternalServerError, 500, err.Error())
        }
    }
}

func main() {
    r := gin.Default()
    
    // Register middleware
    r.Use(ErrorHandlingMiddleware())
    
    // Define a route with a sample API handler
    r.GET("/api/example", func(c *gin.Context) {
        // Sample payload
        payload := map[string]string{"message": "Hello, World!"}
        
        // Call SuccessResponse to format the response
        SuccessResponse(c, http.StatusOK, payload)
    })
    
    // Start the server
    r.Run() // listening and serving on 0.0.0.0:8080
}
