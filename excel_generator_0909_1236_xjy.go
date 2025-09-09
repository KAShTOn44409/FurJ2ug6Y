// 代码生成时间: 2025-09-09 12:36:19
package main

import (
    "encoding/json"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/xuri/excelize/v2"
)

// ExcelGeneratorHandler 处理生成Excel表格的请求
func ExcelGeneratorHandler(c *gin.Context) {
    // 尝试生成Excel文件
    err := generateExcel()
    if err != nil {
        // 如果发生错误，返回错误信息
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    // 成功生成Excel文件
    c.JSON(http.StatusOK, gin.H{
        "message": "Excel file generated successfully",
    })
}

// generateExcel 是生成Excel文件的内部函数
func generateExcel() error {
    // 创建一个新的Excel文件
    f := excelize.NewFile()
    // 创建一个名为"Sheet1"的工作表
    index := f.NewSheet("Sheet1")
    // 设置工作表的名称
    f.SetActiveSheet(index)
    // 添加一些示例数据
    f.SetCellValue("Sheet1", "A2", "Name")
    f.SetCellValue("Sheet1", "B2", "Age")
    f.SetCellValue("Sheet1", "A3", "John Doe")
    f.SetCellValue("Sheet1", "B3", 30)
    
    // 保存文件到磁盘
    if err := f.SaveAs("example.xlsx"); err != nil {
        return err
    }
    // 返回nil表示成功
    return nil
}

func main() {
    r := gin.Default()

    // 注册Excel生成器处理器
    r.GET("/generate_excel", ExcelGeneratorHandler)

    // 启动Gin服务器
    r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
}
