// 代码生成时间: 2025-08-01 14:06:56
package main

import (
    "fmt"
    "log"
    "net/http"
    "sort"

    "github.com/gin-gonic/gin"
)

// SortingAlgorithmHandler 结构体，用于处理排序算法的请求
type SortingAlgorithmHandler struct {
}

// Sort 排序函数，接受一个整数切片，并返回排序后的切片
func (h *SortingAlgorithmHandler) Sort(c *gin.Context) {
    // 从请求体中获取整数切片
    var numbers []int
    if err := c.ShouldBindJSON(&numbers); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": fmt.Sprintf("invalid input: %v", err),
        })
        return
    }

    // 使用 sort 包中的 Sort 函数进行排序
    sort.Ints(numbers)

    // 返回排序后的数组
    c.JSON(http.StatusOK, gin.H{
        "sorted_numbers": numbers,
    })
}

func main() {
    router := gin.Default()

    // 创建 SortingAlgorithmHandler 实例
    handler := &SortingAlgorithmHandler{}

    // 注册路由和处理器
    router.POST("/sort", handler.Sort)

    // 启动服务器
    log.Fatal(router.Run(":8080"))
}
