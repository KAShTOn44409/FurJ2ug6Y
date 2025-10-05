// 代码生成时间: 2025-10-06 01:34:24
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// JsonTransformerHandler is a Gin handler function for JSON data format conversion.
func JsonTransformerHandler(c *gin.Context) {
    // Define the input and output structs for JSON data.
    type Input struct {
        Data string `json:"data"`
    }
    type Output struct {
        TransformedData string `json:"transformedData"`
    }

    var input Input

    // Bind the JSON data from the request body to the input struct.
    if err := c.ShouldBindJSON(&input); err != nil {
        // Return an error response if JSON binding fails.
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid JSON input",
        })
        return
    }

    // Transform the input data (for demonstration, just reverse the string).
    transformedData := reverseString(input.Data)

    // Create an output struct with the transformed data.
    output := Output{TransformedData: transformedData}

    // Return the transformed data as JSON response.
    c.JSON(http.StatusOK, output)
}

// reverseString reverses a given string.
func reverseString(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func main() {
    // Initialize a new Gin router.
    router := gin.Default()

    // Register the JSON transformer handler with a route.
    router.POST("/transform", JsonTransformerHandler)

    // Start the server on port 8080.
    router.Run(":8080")
}
