// 代码生成时间: 2025-09-19 09:37:01
package main

import (
    "net"
    "time"
    "github.com/gin-gonic/gin"
    "net/http"
)

// NetworkStatusChecker 结构体用于表示网络连接状态检查器
type NetworkStatusChecker struct {
    // 可以添加配置字段，比如超时时间等
}

// NewNetworkStatusChecker 创建一个新的网络状态检查器实例
func NewNetworkStatusChecker() *NetworkStatusChecker {
    return &NetworkStatusChecker{}
}

// CheckNetworkStatus 检查指定地址的网络连接状态
func (nsc *NetworkStatusChecker) CheckNetworkStatus(c *gin.Context) {
    // 从请求中获取目标地址
    target := c.Query("target")
    if target == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Target address is required",
        })
        return
    }

    // 设置超时时间
    timeout := 5 * time.Second
    conn, err := net.DialTimeout("tcp", target, timeout)
    if err != nil {
        // 处理错误情况
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    defer conn.Close()

    // 网络连接成功，返回状态码200
    c.JSON(http.StatusOK, gin.H{
        "status": "connected",
        "target": target,
    })
}

func main() {
    // 创建路由器
    router := gin.Default()

    // 创建网络状态检查器实例
    nsc := NewNetworkStatusChecker()

    // 定义路由，使用网络状态检查器处理GET请求
    router.GET("/check", nsc.CheckNetworkStatus)

    // 启动服务
    router.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
