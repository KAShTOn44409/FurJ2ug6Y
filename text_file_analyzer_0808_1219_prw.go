// 代码生成时间: 2025-08-08 12:19:01
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

// TextFileAnalyzerHandler is a HTTP handler that analyzes the content of a text file
func TextFileAnalyzerHandler(c *gin.Context) {
    // Check if a file path is provided in the query string
    filePath := c.Query("path")
    if filePath == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "No file path provided",
        })
        return
    }

    // Check if the file exists
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "File not found",
        })
        return
    } else if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Error checking file status",
        })
        return
    }

    // Read the file content
    fileContent, err := ioutil.ReadFile(filePath)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Error reading file",
        })
        return
    }

    // Simple analysis: count the number of lines and words in the file content
    lines := 0
    words := 0
    for _, line := range strings.Fields(string(fileContent)) {
        lines++
        words += len(strings.Fields(line))
    }

    // Return the analysis result
    c.JSON(http.StatusOK, gin.H{
        "filename": filePath,
        "lines": lines,
        "words": words,
    })
}

func main() {
    r := gin.Default()

    // Register the TextFileAnalyzerHandler
    r.GET("/analyze", TextFileAnalyzerHandler)

    // Start the server
    log.Fatal(r.Run(":8080"))
}
