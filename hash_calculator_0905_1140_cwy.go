// 代码生成时间: 2025-09-05 11:40:27
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "log"
# 增强安全性
    "net/http"
    "github.com/gin-gonic/gin"
)

// HashCalculatorHandler handles the request for calculating hash values.
func HashCalculatorHandler(c *gin.Context) {
    // Get the input string from the query parameter.
    input := c.Query("input")
    if input == "" {
        // Return an error response if input is not provided.
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Input parameter 'input' is required.",
        })
        return
    }
# 添加错误处理

    // Calculate the SHA256 hash of the input.
    hash := sha256.Sum256([]byte(input))
    hashString := hex.EncodeToString(hash[:])

    // Return the hash value in the response.
# 增强安全性
    c.JSON(http.StatusOK, gin.H{
        "hash": hashString,
    })
}

func main() {
    // Create a new Gin router.
    router := gin.Default()

    // Register the HashCalculatorHandler for the route /hash.
    router.GET("/hash", HashCalculatorHandler)

    // Start the server on port 8080.
    log.Fatal(router.Run(":8080"))
}
