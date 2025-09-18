// 代码生成时间: 2025-09-18 14:08:08
package main

import (
    "fmt"
    "os"
    "time"

    "github.com/gin-gonic/gin"
)

// 日志记录器，用于将日志写入文件
type Logger struct {
    File *os.File
}

// Write 实现了日志写入操作
func (l *Logger) Write(p []byte) (n int, err error) {
    return l.File.Write(append(p, '
'))
}

// NewLogger 创建一个新的日志记录器实例
func NewLogger(filePath string) (*Logger, error) {
    file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        return nil, err
    }
    return &Logger{File: file}, nil
}

// AuditLogMiddleware 是Gin中间件，用于记录安全审计日志
func AuditLogMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 开始时间
        start := time.Now()

        // 处理请求
        c.Next()

        // 结束时间
        duration := time.Since(start)

        // 获取请求相关信息
        method := c.Request.Method
        path := c.Request.URL.Path
        status := c.Writer.Status()

        // 记录日志
        logMsg := fmt.Sprintf("[AUDIT] %s - %s - %d - %v", method, path, status, duration)
        c.Get("logger").(*Logger).Write([]byte(logMsg))
    }
}

func main() {
    // 创建Gin路由器
    r := gin.New()

    // 设置日志文件路径
    logFile := "audit_log.txt"
    logger, err := NewLogger(logFile)
    if err != nil {
        fmt.Printf("Failed to create logger: %s", err)
        return
    }

    // 将日志记录器添加到Gin上下文中
    r.SetLogger(logger)

    // 添加中间件
    r.Use(AuditLogMiddleware())

    // 路由和处理函数
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    // 启动服务器
    if err := r.Run(); err != nil {
        fmt.Printf("Failed to run server: %s", err)
    }
}
