// 代码生成时间: 2025-09-01 18:01:55
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

// BulkRenameHandler 结构体包含旧文件名和新文件名的映射
type BulkRenameHandler struct {
    OldToNew map[string]string
}

// NewBulkRenameHandler 创建一个新的BulkRenameHandler实例
func NewBulkRenameHandler(oldToNew map[string]string) *BulkRenameHandler {
    return &BulkRenameHandler{OldToNew: oldToNew}
}

// Handle 处理批量重命名请求
func (b *BulkRenameHandler) Handle(c *gin.Context) {
    // 从请求中获取文件路径列表
    paths := c.PostFormArray("paths")

    // 遍历所有路径
    for _, path := range paths {
        // 获取旧文件名和新文件名
        oldName, newName := filepath.Base(path), path + c.PostForm("newName")

        // 检查旧文件名是否在映射中
        if newOldName, exists := b.OldToNew[oldName]; exists {
            oldName = newOldName
        }

        // 构造完整的旧路径和新路径
        oldPath := filepath.Join(c.PostForm("dir"), oldName)
        newPath := filepath.Join(c.PostForm("dir"), newName)

        // 检查旧文件是否存在
        if _, err := os.Stat(oldPath); os.IsNotExist(err) {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": fmt.Sprintf("File %s does not exist", oldPath),
            })
            return
        }

        // 重命名文件
        if err := os.Rename(oldPath, newPath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": fmt.Sprintf("Failed to rename file %s to %s: %v", oldPath, newPath, err),
            })
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Files renamed successfully",
    })
}

func main() {
    r := gin.Default()

    // 中间件，用于日志记录
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    // 创建批量重命名处理器实例
    renameHandler := NewBulkRenameHandler(map[string]string{
        "oldfile1.txt": "newfile1.txt",
        "oldfile2.txt": "newfile2.txt",
    })

    // 注册POST路由，用于处理批量文件重命名
    r.POST("/bulk-rename", renameHandler.Handle)

    // 启动服务
    if err := r.Run(); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}