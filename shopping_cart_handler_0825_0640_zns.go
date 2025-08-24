// 代码生成时间: 2025-08-25 06:40:08
package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

// CartItem 购物车中的商品项
type CartItem struct {
    ID      int    "json:"id""
    Name    string "json:"name""
    Quantity int    "json:"quantity""
}

// Cart 购物车
type Cart struct {
    Items map[int]CartItem
}

// NewCart 创建一个新的购物车
func NewCart() *Cart {
    return &Cart{Items: make(map[int]CartItem)}
}

// AddItem 向购物车添加商品
func (c *Cart) AddItem(item CartItem) {
    if _, exists := c.Items[item.ID]; exists {
        c.Items[item.ID].Quantity = c.Items[item.ID].Quantity + item.Quantity
    } else {
        c.Items[item.ID] = item
    }
}

// RemoveItem 从购物车移除商品
func (c *Cart) RemoveItem(itemID int) {
    if _, exists := c.Items[itemID]; exists {
        delete(c.Items, itemID)
    }
}

// CartHandler 购物车处理函数
func CartHandler(c *gin.Context) {
    cart := NewCart()
    var item CartItem
    // 解析请求体
    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": fmt.Sprintf("invalid item: %v", err),
        })
        return
    }
    // 添加到购物车
    cart.AddItem(item)
    // 返回购物车内容
    c.JSON(http.StatusOK, cart.Items)
}

// RemoveCartItemHandler 从购物车移除商品的处理函数
func RemoveCartItemHandler(c *gin.Context) {
    itemID := c.Param("id")
    if _, err := strconv.Atoi(itemID); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": fmt.Sprintf("invalid item id: %v", err),
        })
        return
    }
    cart := NewCart() // 这里仅用于演示，实际应用中应从持久层获取购物车
    cart.RemoveItem(itemID)
    // 返回购物车内容
    c.JSON(http.StatusOK, cart.Items)
}

func main() {
    r := gin.Default()

    // 购物车路由
    r.POST("/cart", CartHandler)
    r.DELETE("/cart/:id", RemoveCartItemHandler)

    // 启动服务器
    r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
