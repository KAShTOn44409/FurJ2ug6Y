// 代码生成时间: 2025-09-05 07:47:03
package main

import (
    "net/http"
    "net/url"
    "strings"

    "github.com/gin-gonic/gin"
)

// URLValidatorMiddleware 是一个Gin中间件，用于验证请求中的URL链接是否有效
func URLValidatorMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 获取URL参数
        rawURL := c.Query("url")
        if rawURL == "" {
            // 如果URL参数为空，返回400 Bad Request
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "URL parameter is required",
            })
            c.Abort()
            return
        }

        // 验证URL格式
        u, err := url.ParseRequestURI(rawURL)
        if err != nil || !strings.HasPrefix(u.Scheme, "http") || !strings.HasSuffix(u.Host, ".com") {
            // 如果URL格式不正确或不是有效的.com域名，返回400 Bad Request
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid URL format or host",
            })
            c.Abort()
            return
        }

        // 如果URL有效，向下传递请求
        c.Next()
    }
}

// validateURLHandler 是一个处理器函数，用于处理请求并验证URL
func validateURLHandler(c *gin.Context) {
    // 此处可以添加额外的逻辑，比如数据库查询或业务逻辑处理
    c.JSON(http.StatusOK, gin.H{
        "message": "URL is valid",
    })
}

func main() {
    r := gin.Default()

    // 使用URLValidatorMiddleware中间件
    r.Use(URLValidatorMiddleware())

    // 定义路由和处理器函数
    r.GET("/validate", validateURLHandler)

    // 启动服务器
    r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
}
