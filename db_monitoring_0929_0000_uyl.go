// 代码生成时间: 2025-09-29 00:00:37
package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql" // MySQL driver
)

// DatabaseConfig holds the configuration for the database connection.
type DatabaseConfig struct {
    User     string
    Password string
    Host     string
    Port     string
    Name     string
}

// DBMonitor is a handler for monitoring database.
func DBMonitor(dbConfig DatabaseConfig) gin.HandlerFunc {
    return func(c *gin.Context) {
        var db *sqlx.DB
        var err error

        // Connect to the database
        db, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Unable to connect to the database",
            })
            return
        }
        defer db.Close()

        // Perform a health check query
        var dbTime time.Time
        err = db.Get(&dbTime, "SELECT NOW()")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Database query failed",
            })
            return
        }

        // Return the database time
        c.JSON(http.StatusOK, gin.H{
            "status":   "success",
            "database_time": dbTime.Format(time.RFC1123),
        })
    }
}

func main() {
    // Define the database configuration
    dbConfig := DatabaseConfig{
        User:     "your_username",
        Password: "your_password",
        Host:     "localhost",
        Port:     "3306",
        Name:     "your_db_name",
    }

    // Create a new Gin router
    router := gin.Default()

    // Register the DBMonitor handler
    router.GET("/db/monitor", DBMonitor(dbConfig))

    // Start the server
    router.Run(":8080")
}