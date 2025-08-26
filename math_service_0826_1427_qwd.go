// 代码生成时间: 2025-08-26 14:27:48
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Define the structure to encapsulate mathematical operations
// 定义数学运算的结构体
type MathService struct {
}

// Add adds two numbers and returns the result
// 加法运算
func (m *MathService) Add(a, b float64) (float64, error) {
    if a < 0 || b < 0 {
        return 0, nil // Assuming non-negative numbers for simplicity
    }
    return a + b, nil
}

// Subtract subtracts two numbers and returns the result
// 减法运算
func (m *MathService) Subtract(a, b float64) (float64, error) {
    if a < 0 || b < 0 {
        return 0, nil // Assuming non-negative numbers for simplicity
    }
    return a - b, nil
}

// Multiply multiplies two numbers and returns the result
// 乘法运算
func (m *MathService) Multiply(a, b float64) (float64, error) {
    if a < 0 || b < 0 {
        return 0, nil // Assuming non-negative numbers for simplicity
    }
    return a * b, nil
}

// Divide divides two numbers and returns the result
// 除法运算
func (m *MathService) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    r := gin.Default()

    // Define a mathService instance
    mathService := MathService{}

    // Define routes with their respective handlers
    // 定义路由和对应的处理器
    r.GET("/add", func(c *gin.Context) {
        a, _ := strconv.ParseFloat(c.Query("a"), 64)
        b, _ := strconv.ParseFloat(c.Query("b"), 64)
        result, err := mathService.Add(a, b)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input for addition"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"result": result})
    })

    r.GET("/subtract", func(c *gin.Context) {
        a, _ := strconv.ParseFloat(c.Query("a"), 64)
        b, _ := strconv.ParseFloat(c.Query("b"), 64)
        result, err := mathService.Subtract(a, b)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input for subtraction"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"result": result})
    })

    r.GET("/multiply", func(c *gin.Context) {
        a, _ := strconv.ParseFloat(c.Query("a"), 64)
        b, _ := strconv.ParseFloat(c.Query("b"), 64)
        result, err := mathService.Multiply(a, b)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input for multiplication"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"result": result})
    })

    r.GET("/divide", func(c *gin.Context) {
        a, _ := strconv.ParseFloat(c.Query("a"), 64)
        b, _ := strconv.ParseFloat(c.Query("b"), 64)
        result, err := mathService.Divide(a, b)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot divide by zero"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"result": result})
    })

    // Start the server
    // 启动服务器
    r.Run(":8080") // listening and serving on 0.0.0.0:8080
}
