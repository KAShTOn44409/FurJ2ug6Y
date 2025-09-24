// 代码生成时间: 2025-09-24 17:23:22
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// ShoppingCart 购物车结构体
type ShoppingCart struct {
    ID     string   `json:"id"`
    Items  []string `json:"items"`
}

// NewShoppingCart 创建一个新的购物车实例
func NewShoppingCart() *ShoppingCart {
    return &ShoppingCart{
        ID:    "1",
        Items: make([]string, 0),
    }
}

// AddItem 向购物车添加商品
func (sc *ShoppingCart) AddItem(item string) {
    sc.Items = append(sc.Items, item)
}

// RemoveItem 从购物车移除商品
func (sc *ShoppingCart) RemoveItem(item string) {
    for i, v := range sc.Items {
        if v == item {
            sc.Items = append(sc.Items[:i], sc.Items[i+1:]...)
            break
        }
    }
}

// GetCart 获取购物车内容
func (sc *ShoppingCart) GetCart() []string {
    return sc.Items
}

// CartHandler 处理购物车请求的Gin处理器
func CartHandler(c *gin.Context) {
    // 创建购物车实例
    cart := NewShoppingCart()

    // 处理添加商品请求
    if len(c.PostForm("item")) > 0 {
        item := c.PostForm("item")
        cart.AddItem(item)
        c.JSON(http.StatusOK, gin.H{
            "status":  "success",
            "message": "Item added to cart",
            "cart":    cart.GetCart(),
        })
        return
    }

    // 处理移除商品请求
    if len(c.Query("remove")) > 0 {
        item := c.Query("remove")
        cart.RemoveItem(item)
        c.JSON(http.StatusOK, gin.H{
            "status":  "success",
            "message": "Item removed from cart",
            "cart":    cart.GetCart(),
        })
        return
    }

    // 默认返回购物车内容
    c.JSON(http.StatusOK, gin.H{
        "status":  "success",
        "message": "Cart retrieved",
        "cart":    cart.GetCart(),
    })
}

func main() {
    r := gin.Default()

    // 使用内置的Logger和Recovery中间件
    // Logger中间件会记录请求日志
    // Recovery中间件会捕获任何发生的panic，并返回500状态码
    r.Use(gin.Logger(), gin.Recovery())

    // 注册购物车处理器
    r.GET("/cart", CartHandler)
    r.POST("/cart", CartHandler)

    // 启动服务器
    r.Run(":8080")
}
