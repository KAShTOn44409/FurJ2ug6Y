// 代码生成时间: 2025-08-21 15:53:11
package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

// CSVBatchProcessor 结构体用于文件处理
type CSVBatchProcessor struct{
    // 添加任何需要的字段
}

// NewCSVBatchProcessor 创建新的CSVBatchProcessor实例
func NewCSVBatchProcessor() *CSVBatchProcessor {
    return &CSVBatchProcessor{}
}

// ProcessCSV 处理上传的CSV文件
func (p *CSVBatchProcessor) ProcessCSV(fileHeader *multipart.FileHeader) (err error) {
    // 打开文件
    file, err := fileHeader.Open()
    if err != nil {
        return err
    }
    defer file.Close()

    // 创建CSV读取器
    reader := csv.NewReader(file)

    // 读取CSV数据
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
        // 处理CSV记录，这里只是个示例，可以根据需要进行修改
        fmt.Println(record)
    }

    return nil
}

// setupRoutes 设置Gin路由
func setupRoutes(r *gin.Engine) {
    r.POST("/process", func(c *gin.Context) {
        // 多文件上传
        form, err := c.MultipartForm()
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }

        files := form.File["file"]
        if len(files) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "No files were uploaded",
            })
            return
        }

        processor := NewCSVBatchProcessor()
        for _, fileHeader := range files {
            if err := processor.ProcessCSV(fileHeader); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{
                    "error":   "Failed to process CSV file",
                    "file_name": fileHeader.Filename,
                    "message":  err.Error(),
                })
                return
            }
        }

        c.JSON(http.StatusOK, gin.H{
            "message":   "All files were processed successfully",
            "processed": len(files),
        })
    })
}

func main() {
    r := gin.Default()
    setupRoutes(r)

    // 运行服务器
    if err := r.Run(":8080"); err != nil {
        fmt.Println("Failed to start server: ", err)
        os.Exit(1)
    }
}