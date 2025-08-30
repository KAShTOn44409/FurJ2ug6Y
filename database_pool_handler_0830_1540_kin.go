// 代码生成时间: 2025-08-30 15:40:28
package main

import (
    "fmt"
    "log"
    "database/sql"
    \_ "github.com/go-sql-driver/mysql" // MySQL driver
    "github.com/gin-gonic/gin"
)

// DatabaseConfig 用于配置数据库连接
type DatabaseConfig struct {
    Username string
    Password string
    Host     string
    Port     string
    Database string
}

// DatabasePool 管理数据库连接池
type DatabasePool struct {
    *sql.DB
}

// NewDatabasePool 创建一个新的数据库连接池
func NewDatabasePool(config *DatabaseConfig) (*DatabasePool, error) {
    // 构建DSN（数据源名称）
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.Database)

    // 打开数据库连接
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // 设置数据库连接池参数
    db.SetMaxOpenConns(100)      // 设置最大打开的连接数，默认值为0没有限制
    db.SetMaxIdleConns(50)       // 设置连接池中的最大闲置连接数
    db.SetConnMaxLifetime(3600) // 设置了连接的最大存活时间，单位为秒，这里设置为1小时

    // 测试数据库连接
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &DatabasePool{db}, nil
}

func main() {
    // 定义数据库配置
    config := &DatabaseConfig{
        Username: "your_username",
        Password: "your_password",
        Host:     "localhost",
        Port:     "3306",
        Database: "your_database",
    }

    // 创建数据库连接池
    dbPool, err := NewDatabasePool(config)
    if err != nil {
        log.Fatalf("Failed to create database pool: %v", err)
    }
    defer dbPool.Close() // 确保程序退出时关闭数据库连接池

    // 设置Gin中间件
    r := gin.Default()

    // 定义一个简单的API来测试数据库连接
    r.GET("/ping", func(c *gin.Context) {
        // 从连接池中获取一个连接
        db := dbPool.DB
        // 测试连接
        if err := db.Ping(); err != nil {
            c.JSON(500, gin.H{"error": "Database connection failed"})
            return
        }
        c.JSON(200, gin.H{"message": "Database connection is alive"})
    })

    // 启动Gin服务器
    r.Run() //默认监听并在 0.0.0.0:8080 上启动服务
}