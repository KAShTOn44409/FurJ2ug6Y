// 代码生成时间: 2025-08-14 10:18:30
package main

import (
    "fmt"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
)

// TestDataGeneratorHandler is the handler function for generating test data.
// It generates a slice of integers and returns a JSON response.
func TestDataGeneratorHandler(c *gin.Context) {
    numItems := c.DefaultQuery("count", "10") // default to 10 items
    _, err := strconv.Atoi(numItems)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid count parameter, must be a positive integer.",
        })
        return
    }

    // Generate the test data.
    testData := make([]int, 0, 10)
    for i := 0; i < numItems; i++ {
        testData = append(testData, i)
    }

    // Return the generated test data in JSON format.
    c.JSON(http.StatusOK, testData)
}

// SetupGinRouter initializes the Gin router and sets up the route for the test data generator.
func SetupGinRouter() *gin.Engine {
    r := gin.Default() // Initialize a new Gin router.

    // Setup a route for the test data generator.
    r.GET("/test-data", TestDataGeneratorHandler)

    return r
}

// main function to start the server.
func main() {
    r := SetupGinRouter()
    r.Run(":8080") // Start the server on port 8080.
}
