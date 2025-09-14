// 代码生成时间: 2025-09-15 03:52:46
package main

import (
    "archive/tar"
    "archive/zip"
    "bytes"
    "compress/gzip"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strings"

    "github.com/gin-gonic/gin"
)

// DecompressHandler is a Gin handler that decompresses uploaded files.
func DecompressHandler(c *gin.Context) {
    // Check if the file is in the request
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "No file in the request",
        })
        return
    }

    // Open the file
    src, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Error opening file",
        })
        return
    }
    defer src.Close()

    // Create a buffer to store the decompressed file
    buffer := new(bytes.Buffer)

    // Determine the file extension and decompress accordingly
    switch strings.ToLower(filepath.Ext(file.Filename)) {
    case ".zip":
        // Decompress zip files
        if err := decompressZip(src, buffer); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
            return
        }
    case ".tar.gz", ".tgz":
        // Decompress tar.gz files
        if err := decompressTarGz(src, buffer); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
            return
        }
    case ".gz":
        // Decompress gzip files
        if err := decompressGz(src, buffer); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
            return
        }
    default:
        // If the file is not a recognized compression format
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Unsupported file format",
        })
        return
    }

    // Write the decompressed content to a file (or handle it as needed)
    // For demonstration purposes, we're simply writing to the response
    c.DataFromReader(http.StatusOK, int64(buffer.Len()), "application/octet-stream", buffer, nil)
}

// decompressZip decompresses a zip file.
func decompressZip(reader io.Reader, writer io.Writer) error {
    z, err := zip.NewReader(reader, int64(src.Size))
    if err != nil {
        return err
    }
    for _, f := range z.File {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer rc.Close()
        _, err = io.Copy(writer, rc)
        if err != nil {
            return err
        }
    }
    return nil
}

// decompressTarGz decompresses a tar.gz file.
func decompressTarGz(reader io.Reader, writer io.Writer) error {
    gz, err := gzip.NewReader(reader)
    if err != nil {
        return err
    }
    defer gz.Close()
    tr := tar.NewReader(gz)
    for {
        hdr, err := tr.Next()
        if err == io.EOF {
            break // End of archive
        }
        if err != nil {
            return err
        }
        if !strings.HasPrefix(hdr.Name, "./") {
            hdr.Name = "./" + hdr.Name
        }
        if err := writeHeader(hdr, writer); err != nil {
            return err
        }
        if _, err := io.Copy(writer, tr); err != nil {
            return err
        }
    }
    return nil
}

// decompressGz decompresses a gzip file.
func decompressGz(reader io.Reader, writer io.Writer) error {
    gz, err := gzip.NewReader(reader)
    if err != nil {
        return err
    }
    defer gz.Close()
    _, err = io.Copy(writer, gz)
    return err
}

// writeHeader writes a tar header to the writer.
func writeHeader(hdr *tar.Header, writer io.Writer) error {
    fmt.Fprintf(writer, "%s %d %d %o %d %d %d %d %d %s %s %s
",
        hdr.Name,
        hdr.Size,
        hdr.Mode,
        hdr.Uid,
        hdr.Gid,
        hdr.Uname,
        hdr.Gname,
        dereferenceTime(hdr.ModTime),
        dereferenceTime(hdr.AccessTime),
        dereferenceTime(hdr.ChangeTime),
        "",
        "",
    )
    return nil
}

// dereferenceTime converts a time.Time to an int, representing seconds since the Unix epoch.
func dereferenceTime(t time.Time) int {
    return int(t.Unix())
}

func main() {
    r := gin.Default()
    r.POST("/decompress", DecompressHandler)
    r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
