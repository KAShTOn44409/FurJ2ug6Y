// 代码生成时间: 2025-09-21 01:43:51
package main

import (
    "fmt"
    "net/http"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/PuerkitoBio/goquery" // 用于HTML解析
)

// WebContentScraper 是该应用的主要结构体
type WebContentScraper struct {
    // 可以添加配置字段
}

func main() {
    r := gin.Default()
    r.Use(gin.Recovery()) // 使用Gin的中间件来处理panic

    // 创建WebContentScraper实例
    scraper := &WebContentScraper{}

    // 注册路由和处理器
    r.GET("/scrape", scraper.scrapeHandler)

    // 启动服务器
    r.Run() // 默认在8080端口启动
}

// scrapeHandler 是处理网页内容抓取请求的处理器
func (s *WebContentScraper) scrapeHandler(c *gin.Context) {
    url := c.Query("url")
    if url == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Missing URL parameter",
        })
        return
    }

    // 尝试获取网页内容
    content, err := scrapeContent(url)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("Failed to scrape content: %v", err),
        })
        return
    }

    // 返回网页内容
    c.JSON(http.StatusOK, gin.H{
        "url": url,
        "content": content,
    })
}

// scrapeContent 用于从提供的URL抓取网页内容
func scrapeContent(url string) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", fmt.Errorf("http.Get: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("server returned status: %s", resp.Status)
    }

    // 使用goquery解析HTML
    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return "", fmt.Errorf("goquery.NewDocumentFromReader: %v", err)
    }

    // 提取网页的body标签内容
    var content strings.Builder
    doc.Find("body").Html(&content)

    return content.String(), nil
}
