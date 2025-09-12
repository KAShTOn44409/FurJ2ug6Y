// 代码生成时间: 2025-09-12 11:23:22
package main

import (
    "archive/zip"
    "bytes"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "time"
    "github.com/gin-gonic/gin"
)

// 解压文件到指定目录
func unzipFile(archivePath, destPath string) error {
    reader, err := zip.OpenReader(archivePath)
    if err != nil {
        return err
    }
    defer reader.Close()

    for _, file := range reader.File {
        filePath := filepath.Join(destPath, file.Name)
        if file.FileInfo().IsDir() {
            os.MkdirAll(filePath, os.ModePerm)
        } else {
            if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
                return err
            }
            outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
            if err != nil {
                return err
            }
            defer outFile.Close()

            fileContent, err := file.Open()
            if err != nil {
                return err
            }
            defer fileContent.Close()
            _, err = io.Copy(outFile, fileContent)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

// API处理器，用于文件上传和解压
func uploadAndUnzip(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    src, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    defer src.Close()

    // 将上传的文件保存到临时文件
    tempFile, err := os.Create("temp.zip")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    defer tempFile.Close()
    _, err = io.Copy(tempFile, src)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    // 解压文件到指定目录
    err = unzipFile("temp.zip", "unzipped")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "File uploaded and unzipped successfully",
    })
}

func main() {
    r := gin.Default()

    // 文件上传的路由
    r.POST("/upload", uploadAndUnzip)

    // 开始监听并服务
    r.Run()
}
