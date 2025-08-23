// 代码生成时间: 2025-08-23 19:27:45
package main

import (
    "fmt"
    "net/http"
    "runtime"
    "github.com/gin-gonic/gin"
)

// MemoryUsageResponse defines the structure for the response
type MemoryUsageResponse struct {
    // Memory usage in bytes
    MemoryUsage int64 \"json:memoryUsage\"
    // Allocations count
    Allocations uint64 \"json:allocations\"
    // GC pause duration
    GCPauseDuration float64 \"json:gcPauseDuration\"
}

// MemoryUsageHandler handles the request to get memory usage stats
func MemoryUsageHandler(c *gin.Context) {
    // Get memory stats
    var m runtime.MemStats
    runtime.ReadMemStats(&m)

    // Prepare the response
    response := MemoryUsageResponse{
        MemoryUsage:       int64(m.Alloc), // bytes allocated and not yet freed
        Allocations:       m.Mallocs,
        GCPauseDuration:   m.PauseTotalNs / (1000 * 1000), // convert nanoseconds to milliseconds
    }

    // Respond with the memory usage stats
    c.JSON(http.StatusOK, response)
}

func main() {
    // Initialize Gin router
    r := gin.Default()

    // Register the memory usage handler
    r.GET("/memory", MemoryUsageHandler)

    // Start the server
    r.Run() // listening and serving on 0.0.0.0:8080
}