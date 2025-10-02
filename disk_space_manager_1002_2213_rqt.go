// 代码生成时间: 2025-10-02 22:13:44
It demonstrates how to create an HTTP server with endpoints for disk space management.
*/

package main

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "runtime"
    "github.com/gin-gonic/gin"
)

// DiskSpaceManager is a struct that will hold our disk space management functionalities.
type DiskSpaceManager struct{}

// NewDiskSpaceManager returns a new instance of DiskSpaceManager.
func NewDiskSpaceManager() *DiskSpaceManager {
    return &DiskSpaceManager{}
}

// GetDiskSpace returns the available disk space.
func (d *DiskSpaceManager) GetDiskSpace(c *gin.Context) {
    // Use the current directory as an example.
    path := "."
    
    // Get the disk usage statistics.
    fi, err := os.Stat(path)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to get disk space information",
        })
        return
    }

    if fi.IsDir() {
        var size int64
        var fileCount, dirCount int
        err = filepath.WalkDir(path, func(p string, d os.DirEntry, e error) error {
            if e != nil {
                return e
            }
            if d.Type().IsRegular() {
                size += d.Size()
                fileCount++
            }
            if d.IsDir() {
                dirCount++
            }
            return nil
        })
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to walk the directory",
            })
            return
        }
        
        c.JSON(http.StatusOK, gin.H{
            "total_size": size,
            "file_count": fileCount,
            "directory_count": dirCount,
        })
    } else {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "The provided path is not a directory",
        })
    }
}

func main() {
    r := gin.Default()
    diskManager := NewDiskSpaceManager()
    
    // Register the endpoint for getting disk space.
    r.GET("/disk-space", diskManager.GetDiskSpace)
    
    // Start the server.
    r.Run(fmt.Sprintf(":%d", 8080)) // listen and serve on 0.0.0.0:8080
}
