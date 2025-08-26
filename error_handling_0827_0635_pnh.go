// 代码生成时间: 2025-08-27 06:35:43
package main

import (
    "fmt"
    "net/http"
    "log"

    "github.com/gin-gonic/gin"
)

// ErrorResponse 定义了一个错误响应的结构体
type ErrorResponse struct {
    Error string `json:"error"`
}

// response 用于返回JSON响应
func response(c *gin.Context, code int, message string) {
    c.JSON(code, gin.H{
        "message": message,
    })
}

// errorHandler 是一个处理错误的中间件
func errorHandler(c *gin.Context) {
    c.Next()
    
    if len(c.Errors) > 0 {
        for _, e := range c.Errors {
            // 这里可以根据错误类型做不同的处理
            // 例如，可以根据错误类型记录日志，并返回不同的错误消息
            response(c, http.StatusInternalServerError, e.Err)
        }
    }
}

// IndexHandler 是一个示例处理器，返回一个简单的欢迎信息
func IndexHandler(c *gin.Context) {
    response(c, http.StatusOK, "Welcome to the Gin Error Handling Example")
}

func main() {
    // 创建一个新的Gin路由器
    router := gin.Default()
    
    // 注册中间件
    router.Use(errorHandler)
    
    // 注册处理器
    router.GET("/", IndexHandler)
    
    // 启动服务
    log.Printf("Server is running on :8080")
    router.Run(":8080")
}
