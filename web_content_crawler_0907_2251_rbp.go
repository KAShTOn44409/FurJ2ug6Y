// 代码生成时间: 2025-09-07 22:51:21
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "time"
    "github.com/gin-gonic/gin"
)

// WebContentCrawlerHandler 处理HTTP请求并抓取网页内容
func WebContentCrawlerHandler(c *gin.Context) {
    url := c.Query("url")

    // 验证URL是否存在
    if url == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "URL parameter is missing",
        })
        return
    }

    // 发送HTTP请求
    resp, err := http.Get(url)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to fetch URL",
        })
        return
    }
    defer resp.Body.Close()

    // 读取网页内容
    content, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to read response body",
        })
        return
    }

    // 将网页内容转换为字符串
    contentStr := strings.TrimSpace(string(content))

    // 返回网页内容
    c.JSON(http.StatusOK, gin.H{
        "url": url,
        "content": contentStr,
    })
}

// setupRouter 配置Gin路由器和中间件
func setupRouter() *gin.Engine {
    router := gin.Default()

    // 添加中间件
    router.Use(gin.Recovery()) // Recovery中间件用于恢复panic，提供更优雅的错误处理

    // 添加处理器
    router.GET("/crawl", WebContentCrawlerHandler)

    return router
}

func main() {
    // 设置路由器
    router := setupRouter()

    // 启动服务器
    if err := router.Run(":8080"); err != nil {
        fmt.Printf("Failed to start server: %s
", err)
    }
}
