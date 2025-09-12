// 代码生成时间: 2025-09-12 21:52:52
package main

import (
    "net"
    "time"
    "github.com/gin-gonic/gin"
    "log"
)

// NetworkStatusChecker 结构体用于定义网络连接状态检查器
type NetworkStatusChecker struct {
    // HostName 是要检查的主机名
    HostName string
    // Port 是要检查的端口号
    Port int
}

// NewNetworkStatusChecker 是 NetworkStatusChecker 的构造函数
func NewNetworkStatusChecker(hostName string, port int) *NetworkStatusChecker {
    return &NetworkStatusChecker{
        HostName: hostName,
        Port:     port,
    }
}

// CheckNetworkStatus 检查指定主机和端口的网络连接状态
func (n *NetworkStatusChecker) CheckNetworkStatus() (bool, error) {
    conn, err := net.DialTimeout("tcp", net.JoinHostPort(n.HostName, strconv.Itoa(n.Port)), 5*time.Second)
    defer func() {
        if conn != nil {
            conn.Close() // 确保连接被关闭
        }
    }()
    if err != nil {
        return false, err // 如果有错误发生，则返回 false 和错误信息
    }
    return true, nil // 成功连接则返回 true
}

// networkStatusHandler 是 Gin 的处理器函数，用于处理网络状态检查的 HTTP 请求
func networkStatusHandler(c *gin.Context, checker *NetworkStatusChecker) {
    connected, err := checker.CheckNetworkStatus()
    if err != nil {
        c.JSON(500, gin.H{
            "error": err.Error(),
        })
    } else {
        c.JSON(200, gin.H{
            "connected": connected,
        })
    }
}

func main() {
    router := gin.Default()

    // 创建网络状态检查器实例
    checker := NewNetworkStatusChecker("example.com", 80)

    // 设置 Gin 的处理器函数
    router.GET("/network_status", func(c *gin.Context) {
        networkStatusHandler(c, checker)
    })

    // 启动 Gin 服务
    log.Fatal(router.Run(":8080"))
}
