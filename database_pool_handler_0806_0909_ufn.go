// 代码生成时间: 2025-08-06 09:09:48
package main

import (
    "fmt"
    "log"
    "net/http"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/gin-gonic/gin"
)

// DatabaseConfig contains configuration parameters for the database connection.
type DatabaseConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
}

// DB is the global database connection pool.
var DB *gorm.DB

func initDB(cfg DatabaseConfig) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
    
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Migrations can be run here if needed
    // DB.AutoMigrate(&YourModel{})
}

func main() {
    router := gin.Default()

    // Database configuration
    dbConfig := DatabaseConfig{
        Host:     "localhost",
        Port:     "3306",
        User:     "your_user",
        Password: "your_password",
        DBName:   "your_db",
    }
    
    // Initialize database connection pool
    initDB(dbConfig)
    
    router.GET("/ping", func(c *gin.Context) {
        // Ping the database to check the connection
        err := DB.Ping()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to connect to database",
            })
        } else {
            c.JSON(http.StatusOK, gin.H{
                "message": "Database connection is alive",
            })
        }
    })
    
    // Add more routes here
    
    // Start the server
    log.Fatal(router.Run(":8080"))
}
