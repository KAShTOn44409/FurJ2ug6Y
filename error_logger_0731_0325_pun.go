// 代码生成时间: 2025-07-31 03:25:03
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

// ErrorLoggerMiddleware is a Gin middleware that logs errors.
func ErrorLoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        var err error
        defer func() {
            if e := recover(); e != nil {
                err = fmt.Errorf("recovered from a panic: %v", e)
                c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
                    "error": err.Error(),
                })
                log.Printf("[ERROR] %v", err)
            }
        }()

        c.Next()
    }
}

// CustomRecovery is a custom recovery middleware for Gin.
func CustomRecovery(c *gin.Context) {
    panicObj := recover()
    if panicObj != nil {
        log.Printf("[PANIC] Recovered from panic: %+v", panicObj)
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "error": "Internal Server Error",
        })
    }
}

// Logger logs the request and response details.
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Start timer
        start := time.Now()

        path := c.Request.URL.Path
        raw := c.Request.URL.RawQuery

        // Process request
        c.Next()

        // Calculate resolution time
        latency := time.Since(start)

        // Log only when method is not HEAD
        if c.Request.Method != "HEAD" {
            fmt.Printf("[INFO] %s %s %d %s "%s" %s"
", c.Request.Method, path, c.Writer.Status(), latency,
                c.Request.UserAgent(), raw)
        }
    }
}

func main() {
    router := gin.Default()

    router.Use(Logger()) // Log requests
    router.Use(ErrorLoggerMiddleware()) // Error logging
    router.Use(gin.Recovery()) // Default recovery middleware

    // Custom recovery middleware with logging
    router.Use(gin.CustomRecoveryWithWriter(CustomRecovery, os.Stdout))

    // Define a simple route
    router.GET("/error", func(c *gin.Context) {
        // Intentionally cause an error
        panic("something went wrong")
    })

    // Start the server
    router.Run(":8080")
}
