// 代码生成时间: 2025-09-23 01:22:25
package main

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

// InventoryItem represents an item in the inventory.
type InventoryItem struct {
    ID         uint   `json:"id"`
    Name       string `json:"name"`
    StockCount int    `json:"stockCount"`
}

// InventoryManager is responsible for managing the inventory.
type InventoryManager struct {
    items map[uint]InventoryItem
}

// NewInventoryManager creates a new inventory manager with an empty inventory.
func NewInventoryManager() *InventoryManager {
    return &InventoryManager{
        items: make(map[uint]InventoryItem),
    }
}

// AddItem adds a new item to the inventory.
func (im *InventoryManager) AddItem(item InventoryItem) {
    im.items[item.ID] = item
}

// UpdateStock updates the stock count for an item.
func (im *InventoryManager) UpdateStock(id uint, newStock int) error {
    if _, exists := im.items[id]; !exists {
        return fmt.Errorf("item with ID %d not found", id)
    }
    im.items[id].StockCount = newStock
    return nil
}

// GetItem retrieves an item from the inventory by its ID.
func (im *InventoryManager) GetItem(id uint) (*InventoryItem, error) {
    item, exists := im.items[id]
    if !exists {
        return nil, fmt.Errorf("item with ID %d not found", id)
    }
    return &item, nil
}

// InventoryHandler handles inventory-related HTTP requests.
func InventoryHandler(im *InventoryManager) func(c *gin.Context) {
    return func(c *gin.Context) {
        idStr := c.Param("id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "invalid item ID",
            })
            return
        }
        item, err := im.GetItem(uint(id))
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{
                "error": err.Error(),
            })
            return
        }
        c.JSON(http.StatusOK, item)
    }
}

func main() {
    r := gin.Default()
    
    // Create a new inventory manager.
    im := NewInventoryManager()
    
    // Populate the inventory with some initial items.
    im.AddItem(InventoryItem{ID: 1, Name: "Widget", StockCount: 10})
    im.AddItem(InventoryItem{ID: 2, Name: "Gadget", StockCount: 5})
    
    // Define a route to handle inventory requests.
    r.GET("/inventory/:id", InventoryHandler(im))
    
    // Start the server.
    r.Run(":8080\)
}