// 代码生成时间: 2025-08-05 21:39:17
package main

import (
    "math/rand"
    "time"
    "github.com/gin-gonic/gin"
)

// RandomNumberGeneratorHandler is a handler function that generates a random number.
// It takes two parameters, min and max, and returns a random number within that range.
// If the input parameters are not valid, it returns an error.
func RandomNumberGeneratorHandler(c *gin.Context) {
    // Parse the query parameters for min and max
    minQuery := c.DefaultQuery("min", "1")
    maxQuery := c.DefaultQuery("max", "100")
    min, errMin := strconv.Atoi(minQuery)
    max, errMax := strconv.Atoi(maxQuery)

    // Check for errors in parsing the query parameters
    if errMin != nil || errMax != nil {
        c.JSON(400, gin.H{
            "error": "Invalid input parameters",
        })
        return
    }

    // Check if min is less than max
    if min >= max {
        c.JSON(400, gin.H{
            "error": "Minimum value should be less than maximum value",
        })
        return
    }

    // Generate a random number within the specified range
    rand.Seed(time.Now().UnixNano())
    randomNumber := rand.Intn(max-min) + min

    // Send the random number back as JSON
    c.JSON(200, gin.H{
        "randomNumber": randomNumber,
    })
}

func main() {
    // Create a new Gin router
    router := gin.Default()

    // Use Recover middleware to handle any panic
    router.Use(gin.Recovery())

    // Define a route for the random number generator handler
    router.GET("/random", RandomNumberGeneratorHandler)

    // Start the server on port 8080
    router.Run(":8080")
}
