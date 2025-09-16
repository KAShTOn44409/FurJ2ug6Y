// 代码生成时间: 2025-09-16 19:49:45
package main

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/gin-gonic/gin"
)

// Sortable represents a slice of integers that can be sorted.
type Sortable []int

// Len is the number of elements in the collection.
func (s Sortable) Len() int { return len(s) }

// Less reports whether the element with index i should sort before the element with index j.
func (s Sortable) Less(i, j int) bool { return s[i] < s[j] }

// Swap swaps the elements with indexes i and j.
func (s Sortable) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Sort sorts the slice using the provided algorithm.
func (s Sortable) Sort() {
    sort.Slice(s, func(i, j int) bool { return s.Less(i, j) })
}

// GenerateRandomSlice generates a random slice of integers.
func GenerateRandomSlice(size int) Sortable {
    rand.Seed(time.Now().UnixNano())
    return Sortable(rand.Perm(size))
}

// SortingHandler handles the sorting request.
func SortingHandler(c *gin.Context) {
    size := c.DefaultQuery("size", "10")
    sizeInt, err := strconv.Atoi(size)
    if err != nil {
        c.JSON(400, gin.H{
            "error": "Invalid size parameter",
        })
        return
    }

    slice := GenerateRandomSlice(sizeInt)
    slice.Sort()
    c.JSON(200, gin.H{
        "sorted_slice": slice,
    })
}

func main() {
    r := gin.Default()

    // Middleware to handle request logging.
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    r.GET("/sort", SortingHandler)

    // Start the server on port 8080.
    r.Run()
}
