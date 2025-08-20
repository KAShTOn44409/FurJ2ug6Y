// 代码生成时间: 2025-08-20 21:31:50
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"

    "github.com/gin-gonic/gin"
)

// 解析日志文件并返回解析结果
func parseLogFile(c *gin.Context) {
    filePath := c.PostForm("file_path")
    if filePath == "" {
        c.JSON(400, gin.H{
            "error": "file_path is required",
        })
        return
    }

    // 检查文件是否存在
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        c.JSON(400, gin.H{
            "error": "file not found",
        })
        return
    }

    // 读取文件内容并解析
    file, err := os.Open(filePath)
    if err != nil {
        c.JSON(500, gin.H{
            "error": "failed to open file",
        })
        return
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        // 这里可以添加具体的日志解析逻辑
        if strings.Contains(line, "ERROR") {
            lines = append(lines, line)
        }
    }
    if err := scanner.Err(); err != nil {
        c.JSON(500, gin.H{
            "error": "failed to read file",
        })
        return
    }

    c.JSON(200, gin.H{
        "parsed_lines": lines,
    })
}

func main() {
    r := gin.Default()
    r.POST("/parse", parseLogFile)
    r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
