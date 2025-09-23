// 代码生成时间: 2025-09-23 10:09:22
package main

import (
    "fmt"
    "math"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

// MathService 定义数学服务结构体
type MathService struct{}

// AddHandler 处理加法请求
func (s *MathService) AddHandler(c *gin.Context) {
    // 解析参数
    var a, b float64
    var err error
    if a, err = strconv.ParseFloat(c.Query("a"), 64); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for 'a'"})
        return
    }
    if b, err = strconv.ParseFloat(c.Query("b"), 64); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for 'b'"})
        return
    }

    // 计算结果
    result := a + b

    // 返回结果
    c.JSON(http.StatusOK, gin.H{"result": result})
}

// SubtractHandler 处理减法请求
func (s *MathService) SubtractHandler(c *gin.Context) {
    var a, b, result float64
    var err error
    if a, err = strconv.ParseFloat(c.Query("a"), 64); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for 'a'"})
        return
    }
    if b, err = strconv.ParseFloat(c.Query("b"), 64); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for 'b'"})
        return
    }
    result = a - b
    c.JSON(http.StatusOK, gin.H{"result": result})
}

// MultiplyHandler 处理乘法请求
func (s *MathService) MultiplyHandler(c *gin.Context) {
    var a, b, result float64
    var err error
    if a, err = strconv.ParseFloat(c.Query("a"), 64); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for 'a'"})
        return
    }
    if b, err = strconv.ParseFloat(c.Query("b"), 64); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for 'b'"})
        return
    }
    result = a * b
    c.JSON(http.StatusOK, gin.H{"result": result})
}

// DivideHandler 处理除法请求
func (s *MathService) DivideHandler(c *gin.Context) {
    var a, b, result float64
    var err error
    if a, err = strconv.ParseFloat(c.Query("a"), 64); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for 'a'"})
        return
    }
    if b, err = strconv.ParseFloat(c.Query("b"), 64); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for 'b'"})
        return
    }
    if b == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Division by zero"})
        return
    }
    result = a / b
    c.JSON(http.StatusOK, gin.H{"result": result})
}

// PowerHandler 处理幂运算请求
func (s *MathService) PowerHandler(c *gin.Context) {
    var base, exponent, result float64
    var err error
    if base, err = strconv.ParseFloat(c.Query("base"), 64); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for 'base'"})
        return
    }
    if exponent, err = strconv.ParseFloat(c.Query("exponent"), 64); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for 'exponent'"})
        return
    }
    result = math.Pow(base, exponent)
    c.JSON(http.StatusOK, gin.H{"result": result})
}

func main() {
    r := gin.Default()

    // 注册数学计算处理器
    mathService := MathService{}
    r.GET("/add", mathService.AddHandler)
    r.GET("/subtract", mathService.SubtractHandler)
    r.GET="/multiply", mathService.MultiplyHandler)
    r.GET("/div", mathService.DivideHandler)
    r.GET("/power", mathService.PowerHandler)

    // 启动服务
    r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
