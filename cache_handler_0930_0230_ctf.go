// 代码生成时间: 2025-09-30 02:30:25
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

// CacheableData represents data that can be cached
type CacheableData struct {
    Data string `json:"data"`
}

// cacheKey generates a cache key based on the request
func cacheKey(c *gin.Context) string {
    // You can customize this function to generate a more specific cache key
    return c.Request.URL.String()
}

// CacheMiddleware is a Gin middleware that implements caching
func CacheMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Define a cache map
        cache := make(map[string]CacheableData)
        // Define cache expiration time
        cacheDuration := 5 * time.Minute

        key := cacheKey(c)
        if cachedData, exists := cache[key]; exists && time.Since(cachedData.CreatedAt) < cacheDuration {
            // Return cached data
            c.JSON(http.StatusOK, cachedData)
            c.Abort()
            return
        }

        // Proceed with the request
        c.Next()

        // Cache the response if it was successful
        if c.Writer.Status() == http.StatusOK {
            var responseData CacheableData
            if err := c.ShouldBindJSON(&responseData); err == nil {
                responseData.CreatedAt = time.Now()
                cache[key] = responseData
           }
       }
    }
}

func main() {
    router := gin.Default()
    // Use the cache middleware
    router.Use(CacheMiddleware())

    router.GET("/data", func(c *gin.Context) {
        // Simulate data fetching
        responseData := CacheableData{Data: "This is cached data"}
        c.JSON(http.StatusOK, responseData)
    })

    // Handle errors
    router.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{
            "message": "404 Not Found",
        })
    })

    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %s", err)
    }
}
