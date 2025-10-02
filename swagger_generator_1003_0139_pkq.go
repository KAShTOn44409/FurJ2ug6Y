// 代码生成时间: 2025-10-03 01:39:25
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/swag"
)

// @title Swagger API
// @version 1.0
// @description This is a sample Gin application
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.basic BasicAuth
// @in header
// @name Authorization

func main() {
    // Initialize the Gin router
    router := gin.Default()

    // Swagger
    // Automatically generate Swagger documentation for API
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // @Summary Generate Swagger Documentation
    // @Description Generate Swagger Documentation for the API
    // @Tags Swagger
    // @Produce json
    // @Success 200 {string} string "swagger documentation generated"
    // @Failure 500 {string} string "internal server error"
    // @Router /api/v1/swagger [GET]
    router.GET("/api/v1/swagger", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "swagger documentation generated",
        })
    })

    // Use Gin middleware to handle errors
    router.Use(func(c *gin.Context) {
        startTime := time.Now()
        c.Next()
        duration := time.Since(startTime)
        fmt.Printf("%s %s %s %d %s", c.Request.Method, c.Request.URL.Path,
            c.Request.URL.RawQuery, c.Writer.Status(), duration)
    })

    // Start the server
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}

// Generate Swagger documentation for the API
// This function is called by the `swag init` command
// Uncomment the following lines to use `swag init` for auto-generating Swagger documentation
// func init() {
//     // Generate Swagger documentation
//     if err := swag. InitializeGenerator(
//         swag.GeneralAPI{
//             OpenAPI: "2.0",
//             Info: swag.Info{
//                 Title:      "Swagger API",
//                 Version:    "1.0",
//                 Description:"This is a sample Gin application",
//                 TermsOfService: "http://swagger.io/terms/",
//                 Contact: swag.ContactInfo{
//                     Name:  "API Support",
//                     URL:   "http://www.swagger.io/support",
//                     Email: "support@swagger.io",
//                 },
//                 License: swag.LicenseInfo{
//                     Name: "Apache 2.0",
//                     URL:  "http://www.apache.org/licenses/LICENSE-2.0.html",
//                 },
//             },
//         },
//         []swag.Swagger{
//             {
//                 URL: "http://localhost:8080/swagger/doc.json",
//             },
    ); err != nil {
    //     // Error handling
    //     log.Fatalf("failed to generate swagger: %v", err)
    // }
    //}
}
