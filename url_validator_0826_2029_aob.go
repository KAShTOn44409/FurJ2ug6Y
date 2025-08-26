// 代码生成时间: 2025-08-26 20:29:17
package main

import (
    "net/http"
    "net/url"
    "strings"
    "github.com/gin-gonic/gin"
)

// URLValidatorHandler 是一个Gin处理器，用于验证URL链接的有效性
func URLValidatorHandler(c *gin.Context) {
    // 从请求中获取URL参数
    urlStr := c.Query("url")
    if urlStr == "" {
        // 如果URL参数不存在，返回400错误
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "URL parameter is required",
        })
        return
    }

    // 解析URL
    parsedURL, err := url.ParseRequestURI(urlStr)
    if err != nil {
        // 如果URL解析失败，返回400错误
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid URL format",
        })
        return
    }

    // 检查URL是否有效
    if !strings.HasPrefix(parsedURL.Scheme, "http") {
        // 如果URL协议不是http或https，返回400错误
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "URL must start with 'http' or 'https'",
        })
        return
    }

    // 如果URL有效，返回200 OK
    c.JSON(http.StatusOK, gin.H{
        "message": "URL is valid",
    })
}

func main() {
    // 创建Gin路由器
    router := gin.Default()

    // 注册URL验证处理器
    router.GET("/validate_url", URLValidatorHandler)

    // 启动服务器
    router.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
}
