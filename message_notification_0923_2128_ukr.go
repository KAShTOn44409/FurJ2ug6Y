// 代码生成时间: 2025-09-23 21:28:42
package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

// ErrorResponse defines the structure of an error response.
type ErrorResponse struct {
    Error string `json:"error"`
}

// NotificationRequest defines the structure of a notification request.
type NotificationRequest struct {
    Message string `json:"message"`
    Recipient string `json:"recipient"`
}

func main() {
    r := gin.Default()

    // Example middleware for logging requests.
    r.Use(func(c *gin.Context) {
        fmt.Println("Request made to", c.Request.URL.Path)
        c.Next()
    })

    r.POST("/notify", func(c *gin.Context) {
        var req NotificationRequest

        // Bind the JSON body to the NotificationRequest struct.
        if err := c.ShouldBindJSON(&req); err != nil {
            // Return an error response if the binding fails.
            c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request format"})
            return
        }

        // Here you would implement the actual notification logic.
        fmt.Printf("Sending notification to %s: %s
", req.Recipient, req.Message)

        // Return a success response.
        c.JSON(http.StatusOK, gin.H{
            "message": "Notification sent successfully",
            "recipient": req.Recipient,
        })
    })

    // Start the server on port 8080.
    r.Run(":8080")
}
