// 代码生成时间: 2025-08-30 04:51:52
package main

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
)

// ErrorResponse 定义了错误响应的结构
type ErrorResponse struct {
    Error string `json:"error"`
}

// User 定义了用户数据模型
type User struct {
    ID     string `json:"id"`
    Name   string `json:"name"`
    Email  string `json:"email"`
    Avatar string `json:"avatar"`
}

// CreateUserResponse 定义了创建用户后的响应
type CreateUserResponse struct {
    ID     string `json:"id"`
    Name   string `json:"name"`
    Email  string `json:"email"`
    Avatar string `json:"avatar"`
}

// hashPassword 用于生成密码的哈希值
func hashPassword(password string) string {
    // 使用MD5加密密码
    h := md5.New()
    h.Write([]byte(password))
    return hex.EncodeToString(h.Sum(nil))
}

// createUser 处理器，创建新用户
func createUser(c *gin.Context) {
    var newUser User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        // 如果绑定JSON失败，返回错误响应
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
        return
    }
    // 密码加密
    newUser.ID = hashPassword(newUser.Email)
    // 模拟数据库操作
    // 此处省略数据库代码...
    // 返回创建成功的响应
    c.JSON(http.StatusOK, CreateUserResponse{
        ID:     newUser.ID,
        Name:   newUser.Name,
        Email:  newUser.Email,
        Avatar: newUser.Avatar,
    })
}

// main 函数初始化Gin引擎并注册路由
func main() {
    r := gin.Default()

    // 可选：注册中间件
    r.Use(gin.Recovery())
    r.Use(gin.Logger())

    // 注册用户创建路由
    r.POST("/users", createUser)

    // 启动服务器
    r.Run() // 默认监听并在 0.0.0.0:8080 上启动服务
}
