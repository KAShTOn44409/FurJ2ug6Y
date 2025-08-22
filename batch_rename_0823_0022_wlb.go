// 代码生成时间: 2025-08-23 00:22:00
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

// RenameRequest 定义了重命名请求的结构
type RenameRequest struct {
    From string `json:"from"`
    To   string `json:"to"`
}

// RenameResponse 定义了重命名响应的结构
type RenameResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

func main() {
    // 初始化Gin引擎
    router := gin.Default()

    // 使用中间件记录日志
    router.Use(gin.Logger())

    // 使用中间件恢复panic以防止程序崩溃
    router.Use(gin.Recovery())

    // 批量重命名的路由
    router.POST("/rename", renameFiles)

    // 启动服务器
    log.Fatal(router.Run(":8080"))
}

// renameFiles 处理文件重命名请求
func renameFiles(c *gin.Context) {
    var renameRequests []RenameRequest

    // 绑定请求体到RenameRequest结构
    if err := c.ShouldBindJSON(&renameRequests); err != nil {
        // 如果绑定失败，返回错误信息
        c.JSON(http.StatusBadRequest, RenameResponse{
            Status:  "error",
            Message: err.Error(),
        })
        return
    }

    // 遍历所有重命名请求
    for _, req := range renameRequests {
        // 检查源文件是否存在
        if _, err := os.Stat(req.From); os.IsNotExist(err) {
            c.JSON(http.StatusBadRequest, RenameResponse{
                Status:  "error",
                Message: fmt.Sprintf("Source file '%s' does not exist.", req.From),
            })
            return
        }

        // 检查目标文件是否已存在
        if _, err := os.Stat(req.To); err == nil {
            c.JSON(http.StatusBadRequest, RenameResponse{
                Status:  "error",
                Message: fmt.Sprintf("Target file '%s' already exists.", req.To),
            })
            return
        }

        // 重命名文件
        if err := os.Rename(req.From, req.To); err != nil {
            c.JSON(http.StatusInternalServerError, RenameResponse{
                Status:  "error",
                Message: fmt.Sprintf("Failed to rename '%s' to '%s': %v", req.From, req.To, err),
            })
            return
        }
    }

    // 如果所有文件都成功重命名，返回成功响应
    c.JSON(http.StatusOK, RenameResponse{
        Status:  "success",
        Message: "Files renamed successfully.",
    })
}
