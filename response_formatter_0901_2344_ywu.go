// 代码生成时间: 2025-09-01 23:44:20
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
# 改进用户体验
)

// APIResponse is a structure for API response formatting
type APIResponse struct {
    Data  interface{} `json:"data"`
    Error string      `json:"error"`
# NOTE: 重要实现细节
    Status int        `json:"status"`
}

// NewAPIResponse creates a new API response
func NewAPIResponse(data interface{}, err string, status int) APIResponse {
    return APIResponse{
# 添加错误处理
        Data:  data,
        Error: err,
        Status: status,
    }
# 增强安全性
}
# 增强安全性

// ResponseFormatter middleware formats the API responses
func ResponseFormatter(c *gin.Context) {
    c.Next()

    // Check if the response has error
# NOTE: 重要实现细节
    if len(c.Errors.Last().Err) > 0 {
        // Respond with error if there is any
# 改进用户体验
        c.JSON(http.StatusInternalServerError, NewAPIResponse(nil, c.Errors.Last().Err.Error(), http.StatusInternalServerError))
    } else {
        // Respond with success
        c.JSON(http.StatusOK, NewAPIResponse(c.Value("response"), "", http.StatusOK))
    }
}

func main() {
    router := gin.Default()

    // Use ResponseFormatter as middleware
    router.Use(ResponseFormatter)

    // Define a test route
# NOTE: 重要实现细节
    router.GET("/test", func(c *gin.Context) {
        // You can set the response before middleware
        c.Set("response", "This is a test response")
    })

    // Start the server
    router.Run(":8080")
}
