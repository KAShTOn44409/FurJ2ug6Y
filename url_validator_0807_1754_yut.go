// 代码生成时间: 2025-08-07 17:54:28
package main

import (
    "net/http"
    "
    "github.com/gin-gonic/gin"
    "net/url"
    "strings"
)

// UrlValidator checks if a given URL is valid
func UrlValidator(c *gin.Context) {
    var urlStr string
    if err := c.ShouldBind(&urlStr); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // Trim spaces to avoid 'space' URL errors
    urlStr = strings.TrimSpace(urlStr)

    // Validate URL
    parsedUrl, err := url.Parse(urlStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid URL format",
        })
        return
    }

    // Check if scheme and host are present
    if parsedUrl.Scheme == "" || parsedUrl.Host == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "URL must contain a scheme and a host",
        })
        return
    }

    // If all checks pass, return a success message
    c.JSON(http.StatusOK, gin.H{
        "message": "URL is valid",
    })
}

func main() {
    r := gin.Default()

    // Register the URL validator handler
    r.POST("/check_url", UrlValidator)

    // Start the server
    r.Run() // listening on 0.0.0.0:8080
}
