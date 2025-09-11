// 代码生成时间: 2025-09-11 08:59:47
package main

import (
    "fmt"
    "log"
    "net/http"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

// RenameRequest 定义了请求中需要的信息
type RenameRequest struct {
    From string `json:"from" binding:"required"` // 旧文件名
    To   string `json:"to" binding:"required"`   // 新文件名
}

func main() {
    r := gin.Default()
    r.POST("/rename", handleRename)
    r.Run() // 默认在0.0.0.0:8080上启动服务
}

// handleRename 是处理POST请求的处理器
func handleRename(c *gin.Context) {
    var renameReq RenameRequest
    // 绑定请求体到RenameRequest结构体
    if err := c.ShouldBindJSON(&renameReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    // 检查文件是否存在
    if _, err := os.Stat(renameReq.From); os.IsNotExist(err) {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "File does not exist",
        })
        return
    }
    // 检查新文件名是否已存在
    if _, err := os.Stat(renameReq.To); !os.IsNotExist(err) {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "File with the new name already exists",
        })
        return
    }
    // 重命名文件
    if err := os.Rename(renameReq.From, renameReq.To); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    // 成功响应
    c.JSON(http.StatusOK, gin.H{
        "message": "File renamed successfully",
    })
}
