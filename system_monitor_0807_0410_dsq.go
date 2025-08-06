// 代码生成时间: 2025-08-07 04:10:36
package main

import (
    "fmt"
    "net/http"
    "os"
    "runtime"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
)

// SystemMonitorHandler 定义系统性能监控的处理函数
// 该函数返回系统的内存使用情况、goroutine数量、CPU使用率等信息
func SystemMonitorHandler(c *gin.Context) {
    // 获取内存使用情况
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    memoryUsage := float64(m.Alloc) / 1024 / 1024

    // 获取goroutine数量
    goroutineCount := runtime.NumGoroutine()

    // 获取CPU使用率
    var startStats, endStats os.ProcStat
    if err := startStats.Get(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取系统性能数据"})
        return
    }
    time.Sleep(100 * time.Millisecond)
    if err := endStats.Get(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取系统性能数据"})
        return
    }
    cpuUsage := calculateCPUUsage(startStats, endStats)

    // 响应系统性能数据
    c.JSON(http.StatusOK, gin.H{
        "memory_usage": memoryUsage,
        "goroutine_count": goroutineCount,
        "cpu_usage": cpuUsage,
    })
}

// calculateCPUUsage 计算CPU使用率
func calculateCPUUsage(start, end os.ProcStat) float64 {
    totalStart := start.CPUTotalTime
    totalEnd := end.CPUTotalTime
    idleStart := start.CPUIdleTime
    idleEnd := end.CPUIdleTime

    totalCPU := float64(totalEnd - totalStart)
    idleCPU := float64(idleEnd - idleStart)
    return 1 - (idleCPU / totalCPU)
}

func main() {
    r := gin.Default()

    // 使用中间件记录请求日志
    r.Use(gin.Logger())
    // 使用中间件恢复panic，以保持服务器稳定
    r.Use(gin.Recovery())

    // 注册系统性能监控的路由
    r.GET("/monitor", SystemMonitorHandler)

    // 启动服务器
    if err := r.Run(":8080"); err != nil {
        fmt.Printf("服务器启动失败: %v
", err)
    }
}
