// 代码生成时间: 2025-09-20 09:37:59
package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

// 定义一个错误处理函数
func errorHandle(c *gin.Context, message string, code int) {
    c.JSON(http.StatusOK, gin.H{
        "error": message,
        "code": code,
    })
}

// 定义一个示例数据模型
type ExampleData struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

// 示例处理器函数
func exampleHandler(c *gin.Context) {
    data := ExampleData{
        ID:   1,
        Name: "Example",
    }
    c.JSON(http.StatusOK, data)
}

// 实现GET方法的处理器
func getExample(c *gin.Context) {
    // 假设我们从数据库或其他地方获取数据，这里只是示例
    data := ExampleData{
        ID:   1,
        Name: "Example",
    }
    c.JSON(http.StatusOK, data)
}

// 实现POST方法的处理器
func postExample(c *gin.Context) {
    var data ExampleData
    if err := c.ShouldBindJSON(&data); err != nil {
        errorHandle(c, err.Error(), http.StatusBadRequest)
        return
    }
    c.JSON(http.StatusOK, data)
}

// 实现DELETE方法的处理器
func deleteExample(c *gin.Context) {
    id := c.Param("id")
    if id == "" {
        errorHandle(c, "ID not provided", http.StatusBadRequest)
        return
    }
    // 假设我们从数据库删除数据，这里只是示例
    c.JSON(http.StatusOK, gin.H{
        "message": "Deleted successfully",
        "id": id,
    })
}

func main() {
    router := gin.Default()

    // 使用中间件
    router.Use(gin.Recovery())

    // 路由定义
    router.GET("/example", getExample)
    router.POST("/example", postExample)
    router.DELETE("/example/:id", deleteExample)

    // 启动服务
    router.Run(":8080")
}
