// 代码生成时间: 2025-08-31 15:08:50
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/tealeg/xlsx/v3"
)

// 确保这些包被安装了：
// go get github.com/gin-gonic/gin
// go get github.com/tealeg/xlsx/v3

// ExcelGenerator 结构体，包含生成Excel相关的数据和方法
type ExcelGenerator struct {}

// GenerateExcel 生成Excel文件
func (eg *ExcelGenerator) GenerateExcel(c *gin.Context) {
    // 响应格式为Excel文件
    c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
    c.Header("Content-Disposition", "attachment; filename=example.xlsx")

    // 创建一个新的Excel文件
    file := xlsx.NewFile()
    sheet, _ := file.AddSheet("Sheet1")

    // 添加一些示例数据
    sheet.AddRow(&xlsx.Row{
        Cells: []*xlsx.Cell{
            {Value: "Column 1"},
            {Value: "Column 2"},
            {Value: "Column 3"},
        },
    })
    sheet.AddRow(&xlsx.Row{
        Cells: []*xlsx.Cell{
            {Value: "Data 1"},
            {Value: "Data 2"},
            {Value: "Data 3"},
        },
    })

    // 创建文件并写入响应
    fileWriter, err := os.Create("example.xlsx")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file"})
        return
    }
    defer fileWriter.Close()
    err = file.Write(fileWriter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file"})
        return
    }

    // 读取文件内容并写入响应
    c.FileFromPath("example.xlsx")
}

func main() {
    // 创建一个Gin路由器
    router := gin.Default()

    // 路由到GenerateExcel方法
    router.GET("/generate", func(c *gin.Context) {
        eg := ExcelGenerator{}
        eg.GenerateExcel(c)
    })

    // 启动服务器
    log.Fatal(router.Run(":8080"))
}
