// 代码生成时间: 2025-10-13 16:21:50
package main

import (
    "fmt"
    "mime"
    "net/http"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

// FileTypeResponse 定义了文件类型识别的响应结构体
type FileTypeResponse struct {
    FileName string `json:"filename"`
    FileType string `json:"filetype"`
    Error    string `json:"error"`
}

//识别文件类型处理函数
func identifyFileType(c *gin.Context) {
    file, err := c.GetRawData()
    if err != nil {
        c.JSON(http.StatusBadRequest, FileTypeResponse{
            FileName: "unknown",
            FileType: "unknown",
            Error:    err.Error(),
        })
        return
    }

    fileType := http.DetectContentType(file)
    fileName := c.PostForm("filename")

    // 如果没有提供文件名，尝试从Content-Disposition头中解析
    if fileName == "" {
        cd := c.Request.Header.Get("Content-Disposition")
        if cd != "" {
            fileName = filepath.Base(c.Request.Header.Get("Content-Disposition"))
        }
    }

    c.JSON(http.StatusOK, FileTypeResponse{
        FileName: fileName,
        FileType: fileType,
        Error:    "",
    })
}

func main() {
    r := gin.Default()

    // 可以添加中间件，例如Logger和Recovery
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    r.POST("/identify", identifyFileType)
    r.Run() // 默认在0.0.0.0:8080上运行
}
