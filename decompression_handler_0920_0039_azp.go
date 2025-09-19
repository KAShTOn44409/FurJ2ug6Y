// 代码生成时间: 2025-09-20 00:39:00
package main

import (
    "archive/zip"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "strings"

    "github.com/gin-gonic/gin"
)

// DecompressFile decompresses a zip file to a given directory.
func DecompressFile(dst string, archive *zip.ReadCloser) error {
    defer archive.Close()
    reader, err := archive.Reader.File()
    if err != nil {
        return err
    }
    for _, file := range reader.File {
        rc, err := file.Open()
        if err != nil {
            return err
        }
        defer rc.Close()

        // Store the file content in the destination directory.
        path := filepath.Join(dst, file.Name)
        if file.FileInfo().IsDir() {
            os.MkdirAll(path, os.ModePerm)
        } else {
            f, err := os.OpenFile(
                path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
            if err != nil {
                return err
            }
            defer f.Close()
            _, err = io.Copy(f, rc)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

// HandleDecompression is a Gin.HandlerFunc that handles file decompression requests.
func HandleDecompression(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "No file part in the request",
        })
        return
    }
    src, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Could not open file",
        })
        return
    }
    defer src.Close()

    // Create a buffer to store the archive
    archive, err := zip.NewReader(src, file.Size)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid zip file",
        })
        return
    }

    destination := "./decompressed" // Define the destination directory
    if err := DecompressFile(destination, archive); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to decompress file",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "File decompressed successfully",
    })
}

func main() {
    r := gin.Default()

    // Middlewares
    r.Use(gin.Recovery())
    r.Use(gin.Logger())

    r.POST("/decompress", HandleDecompression)

    r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
