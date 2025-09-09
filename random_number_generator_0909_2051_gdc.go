// 代码生成时间: 2025-09-09 20:51:49
package main

import (
    "math/rand"
    "time"
    "github.com/gin-gonic/gin"
)

// RandomNumberGeneratorHandler is the handler function that generates a random number.
// It takes two query parameters: min and max, which define the range of the random number.
// If the parameters are not provided or are invalid, it returns an error.
func RandomNumberGeneratorHandler(c *gin.Context) {
    min := c.DefaultQuery("min", "0")
    max := c.DefaultQuery("max", "100")
    var minInt, maxInt int
    var err error

    // Convert the string query parameters to integers
    minInt, err = strconv.Atoi(min)
    if err != nil {
        c.JSON(400, gin.H{
            "error": "Invalid 'min' parameter. It must be an integer.",
        })
        return
    }

    maxInt, err = strconv.Atoi(max)
    if err != nil {
        c.JSON(400, gin.H{
            "error": "Invalid 'max' parameter. It must be an integer.",
        })
        return
    }

    // Ensure max is greater than min
    if maxInt <= minInt {
        c.JSON(400, gin.H{
            "error": "The 'max' parameter must be greater than 'min'.",
        })
        return
    }

    // Generate a random number within the specified range
    rand.Seed(time.Now().UnixNano())
    randomNumber := rand.Intn(maxInt-minInt) + minInt

    // Return the random number as JSON
    c.JSON(200, gin.H{
        "random_number": randomNumber,
    })
}

func main() {
    router := gin.Default()

    // Use the RandomNumberGeneratorHandler for the path /random
    router.GET("/random", RandomNumberGeneratorHandler)

    // Start the server on port 8080
    router.Run(":8080")
}
