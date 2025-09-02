// 代码生成时间: 2025-09-02 10:02:51
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strings"

    "github.com/gin-gonic/gin"
)

// BackupSyncService 结构体包含源目录和目标目录
type BackupSyncService struct {
    SourceDir string
    TargetDir string
}

// NewBackupSyncService 构造函数
func NewBackupSyncService(sourceDir, targetDir string) *BackupSyncService {
    return &BackupSyncService{
        SourceDir: sourceDir,
        TargetDir: targetDir,
    }
}

// Sync 同步文件方法
func (b *BackupSyncService) Sync(c *gin.Context) {
    // 错误处理
    if b.SourceDir == "" || b.TargetDir == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "SourceDir and TargetDir cannot be empty",
        })
        return
    }

    // 检查源目录是否存在
    if _, err := os.Stat(b.SourceDir); os.IsNotExist(err) {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("Source directory does not exist: %s", b.SourceDir),
        })
        return
    }

    // 同步文件
    err := filepath.Walk(b.SourceDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            // 构建目标文件路径
            targetPath := filepath.Join(b.TargetDir, strings.TrimPrefix(path, b.SourceDir))
            // 确保目标目录存在
            if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
                return err
            }
            // 复制文件
            if _, err := CopyFile(targetPath, path); err != nil {
                return err
            }
        }
        return nil
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            {
                "error": fmt.Sprintf("Error syncing files: %s", err),
            }
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Files synced successfully",
    })
}

// CopyFile 复制文件
func CopyFile(dstName, srcName string) (int64, error) {
    src, err := os.Open(srcName)
    if err != nil {
        return 0, err
    }
    defer src.Close()
    dst, err := os.Create(dstName)
    if err != nil {
        return 0, err
    }
    defer dst.Close()
    return io.Copy(dst, src)
}

func main() {
    router := gin.Default()

    // 可以添加中间件如Logger、Recovery等
    // router.Use(gin.Logger())
    // router.Use(gin.Recovery())

    sourceDir := "/path/to/source"
    targetDir := "/path/to/target"
    service := NewBackupSyncService(sourceDir, targetDir)

    // 设置路由
    router.POST("/sync", service.Sync)

    fmt.Println("Server starting on port 8080")
    log.Fatal(router.Run(":8080"))
}