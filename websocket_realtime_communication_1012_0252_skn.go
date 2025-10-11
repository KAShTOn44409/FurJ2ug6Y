// 代码生成时间: 2025-10-12 02:52:24
package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "net/http"
)

// upgrader 用于将HTTP连接升级为WebSocket连接
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

// WebSocketMessage 定义了WebSocket消息的结构
type WebSocketMessage struct {
    Message string `json:"message"`
}

// wsHandle 处理WebSocket请求的函数
func wsHandle(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        // 如果升级失败，返回错误信息
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    defer conn.Close()

    // 读取WebSocket消息
    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            // 如果读取消息失败，返回错误并关闭连接
            fmt.Println("read error: ", err)
            break
        }
        // 打印收到的消息
        fmt.Printf("recv: %s", message)

        // 回写消息到客户端
        if err := conn.WriteMessage(messageType, message); err != nil {
            fmt.Println("write error: ", err)
            break
        }
    }
}

func main() {
    r := gin.Default()

    // 将`/ws`路由设置为处理WebSocket请求的函数
    r.GET("/ws", wsHandle)

    // 启动服务
    r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
