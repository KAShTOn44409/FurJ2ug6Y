// 代码生成时间: 2025-09-23 15:52:09
package main

import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql" // MySQL driver
    "github.com/gin-gonic/gin"
)

// DatabaseConfig 用于配置数据库连接
type DatabaseConfig struct {
    Host     string
    Port     int
    Username string
    Password string
    DBName   string
}

// DBManager 管理数据库连接池
type DBManager struct {
    *sql.DB
    Config DatabaseConfig
}

// NewDBManager 创建一个新的DBManager实例
func NewDBManager(config DatabaseConfig) (*DBManager, error) {
    // 构建DSN（数据源名称）
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.Username, config.Password, config.Host, config.Port, config.DBName)
    
    // 打开数据库连接
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    
    // 设置数据库连接池参数
    db.SetMaxOpenConns(100) // 设置最大打开的连接数，默认值为0没有限制
    db.SetMaxIdleConns(25)   // 设置连接池中的最大闲置连接，默认值为2
    db.SetConnMaxLifetime(5 * 60 * 60) // 设置了连接的最大存活时间，5小时
    
    // 测试连接池中的连接是否可用，不可用则报错
    if err := db.Ping(); err != nil {
        return nil, err
    }
    
    return &DBManager{DB: db, Config: config}, nil
}

// Close 关闭数据库连接池
func (m *DBManager) Close() error {
    return m.DB.Close()
}

func main() {
    // 配置数据库连接参数
    config := DatabaseConfig{
        Host:     "localhost",
        Port:     3306,
        Username: "your_username",
        Password: "your_password",
        DBName:   "your_dbname",
    }
    
    // 创建数据库管理器
    dbManager, err := NewDBManager(config)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer dbManager.Close()
    
    // 创建Gin路由器
    router := gin.Default()
    
    // 使用中间件记录请求日志
    router.Use(gin.Logger())
    
    // 使用中间件恢复 panic，返回500错误页
    router.Use(gin.Recovery())
    
    // 定义GET路由，测试数据库连接
    router.GET("/testdb", func(c *gin.Context) {
        // 执行数据库查询
        rows, err := dbManager.DB.Query("SELECT 1")
        if err != nil {
            c.JSON(500, gin.H{"error": "Database query failed"})
            return
        }
        defer rows.Close()
        
        // 检查查询结果
        if rows.Next() {
            c.JSON(200, gin.H{"message": "Database connection is working"})
        } else {
            c.JSON(500, gin.H{"error": "No data received from database"})
        }
    })
    
    // 启动服务
    log.Println("Server started on :8080")
    router.Run(":8080")
}
