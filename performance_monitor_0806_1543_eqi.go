// 代码生成时间: 2025-08-06 15:43:07
package main

import (
    "fmt"
    "net/http"
    "os"
    "runtime"
    "runtime/pprof"
    "time"

    "github.com/gin-gonic/gin"
)

// PerformanceMonitor 结构体，用于封装监控工具的状态
type PerformanceMonitor struct {
    // 可以添加更多字段，例如监控的指标等
}

// NewPerformanceMonitor 创建一个新的性能监控工具
func NewPerformanceMonitor() *PerformanceMonitor {
    return &PerformanceMonitor{}
}

// StartCPUProfile 开始CPU性能分析
func (pm *PerformanceMonitor) StartCPUProfile(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // 设置CPU性能分析文件
    file, err := os.Create("cpu.prof")
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    defer file.Close()
    
    // 启动CPU性能分析
    if err := pprof.StartCPUProfile(file); err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    
    // 模拟一些工作
    time.Sleep(10 * time.Second)
    
    // 停止CPU性能分析
    pprof.StopCPUProfile()
    fmt.Fprint(w, "CPU profiling started and stopped")
}

// StartMemoryProfile 开始内存性能分析
func (pm *PerformanceMonitor) StartMemoryProfile(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // 设置内存性能分析文件
    file, err := os.Create("mem.prof")
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    defer file.Close()
    
    // 开启内存分析
    runtime.GC()
    if err := pprof.WriteHeapProfile(file); err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    fmt.Fprint(w, "Memory profiling started and stopped")
}

func main() {
    router := gin.Default()
    pm := NewPerformanceMonitor()
    
    // 注册性能监控工具的路由
    router.GET("/cpu", pm.StartCPUProfile)
    router.GET("/mem", pm.StartMemoryProfile)
    
    // 启动服务
    router.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
