// 代码生成时间: 2025-08-10 05:17:44
package main
# 优化算法效率

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

// ErrorLoggerMiddleware is a Gin middleware that handles error logging.
func ErrorLoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Start time
        startTime := time.Now()

        // Process request
        c.Next()

        // End time
        endTime := time.Now()

        // Execution time
        latencyTime := endTime.Sub(startTime)

        // Get status code
        status := c.Writer.Status()

        // Log only if there is an error
        if status != http.StatusOK {
            logger := log.New(c.Writer, "ERROR: ", log.LstdFlags|log.Lshortfile)
            logger.Printf("Status: %d | Method: %s | Path: %s | Latency: %v | Request Body: %+v",
                status, c.Request.Method, c.Request.URL.Path, latencyTime, c.Request)
        }
    }
}

// main function to initialize Gin router and attach the middleware
func main() {
    router := gin.Default()

    // Attach ErrorLoggerMiddleware to the router
    router.Use(ErrorLoggerMiddleware())

    // Define a test route
    router.GET("/test", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })

    // Start server
    router.Run(":8080")
}
