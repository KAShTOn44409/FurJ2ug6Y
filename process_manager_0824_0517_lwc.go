// 代码生成时间: 2025-08-24 05:17:14
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "os/exec"
    "syscall"

    "github.com/gin-gonic/gin"
)

// ProcessManagerHandler 用于管理进程的处理器
func ProcessManagerHandler(c *gin.Context) {
    command := c.DefaultQuery("command", "")
    if command == "" {
        // 如果没有提供命令，返回错误
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "No command provided"
        })
        return
    }

    // 执行命令
    output, err := exec.Command("sh", "-c", command).CombinedOutput()
    if err != nil {
        // 错误处理
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":  fmt.Sprintf("Failed to execute command: %s", err),
            "output": string(output),
        })
        return
    }

    // 返回命令执行结果
    c.JSON(http.StatusOK, gin.H{
        "output": string(output),
    })
}

// KillProcess 用于结束指定进程的处理器
func KillProcess(c *gin.Context) {
    processID := c.Param("processID")
    if processID == "" {
        // 如果没有提供进程ID，返回错误
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "No process ID provided"
        })
        return
    }

    // 将进程ID转换为int
    pid, err := strconv.Atoi(processID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":  fmt.Sprintf("Invalid process ID: %s", err),
        })
        return
    }

    // 获取进程
    process, err := os.FindProcess(pid)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":  fmt.Sprintf("Failed to find process: %s", err),
        })
        return
    }

    // 结束进程
    if err := process.Signal(syscall.SIGTERM); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":  fmt.Sprintf("Failed to kill process: %s", err),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Process terminated successfully",
    })
}

func main() {
    r := gin.Default()

    // 使用中间件记录请求日志
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    // 管理进程的路由
    r.GET("/process", ProcessManagerHandler)

    // 结束进程的路由
    r.DELETE("/process/:processID", KillProcess)

    // 启动服务器
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %s
", err)
    }
}
