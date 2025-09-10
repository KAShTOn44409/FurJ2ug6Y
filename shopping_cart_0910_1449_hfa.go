// 代码生成时间: 2025-09-10 14:49:55
package main

import (
    "encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
)

// ShoppingCart represents a shopping cart with a list of items.
type ShoppingCart struct {
    Items map[string]int `json:"items"`
}

// NewShoppingCart initializes a new shopping cart.
func NewShoppingCart() *ShoppingCart {
    return &ShoppingCart{
        Items: make(map[string]int),
    }
}

// AddItem adds an item to the shopping cart.
func (cart *ShoppingCart) AddItem(item string, quantity int) error {
    if quantity <= 0 {
        return errors.New("quantity must be greater than 0")
    }
    cart.Items[item] += quantity
    return nil
}

// RemoveItem removes an item from the shopping cart.
func (cart *ShoppingCart) RemoveItem(item string) error {
    if _, exists := cart.Items[item]; !exists {
        return errors.New("item does not exist in the cart")
    }
    delete(cart.Items, item)
    return nil
}

// GetCart returns the current state of the shopping cart.
func (cart *ShoppingCart) GetCart() ([]byte, error) {
    return json.Marshal(cart)
}

func main() {
    router := gin.Default()

    // Middleware that logs the request
    router.Use(gin.Logger())

    // Middleware that recovers from any panics and returns a 500 if necessary
    router.Use(gin.Recovery())

    var cart ShoppingCart

    // Endpoint to add an item to the cart
    router.POST("/add", func(c *gin.Context) {
        var item struct {
            Item   string `json:"item"`
            Quantity int    `json:"quantity"`
        }
        if err := c.ShouldBindJSON(&item); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }
        if err := cart.AddItem(item.Item, item.Quantity); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }
        c.Status(http.StatusOK)
    })

    // Endpoint to remove an item from the cart
    router.DELETE("/remove", func(c *gin.Context) {
        var item struct {
            Item string `json:"item"`
        }
        if err := c.ShouldBindJSON(&item); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }
        if err := cart.RemoveItem(item.Item); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }
        c.Status(http.StatusOK)
    })

    // Endpoint to get the current state of the cart
    router.GET("/cart", func(c *gin.Context) {
        cartData, err := cart.GetCart()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
            return
        }
        c.JSON(http.StatusOK, gin.H{
            "cart": cartData,
        })
    })

    // Run the server on port 8080
    router.Run(":8080")
}
