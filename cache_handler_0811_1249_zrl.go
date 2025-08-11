// 代码生成时间: 2025-08-11 12:49:18
package main

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

// CacheConfig 缓存配置结构体
type CacheConfig struct {
    Duration time.Duration
}

// cacheData 缓存的数据结构
type cacheData struct {
    Data string `json:"data"`
}

// NewCacheConfig 创建一个新的缓存配置实例
func NewCacheConfig(duration time.Duration) *CacheConfig {
    return &CacheConfig{Duration: duration}
}

// CacheMiddleware Gin中间件，用于设置HTTP响应的缓存策略
func CacheMiddleware(config *CacheConfig) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 设置缓存过期时间
        c.Header("Cache-Control", "public, max-age="+strconv.Itoa(int(config.Duration.Seconds()))+", must-revalidate")
        c.Header("Expires", time.Now().Add(config.Duration).Format(http.TimeFormat))
        c.Next()
    }
}

// CacheHandler 缓存策略处理器
func CacheHandler(c *gin.Context) {
    // 尝试从缓存中获取数据，这里假设缓存的数据是静态的
    cachedData := cacheData{Data: "Cached data"}
    jsonData, err := json.Marshal(cachedData)
    if err != nil {
        // 错误处理
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to marshal cache data",
        })
        return
    }
    // 设置Content-Type为JSON
    c.Header("Content-Type", "application/json")
    // 返回缓存的数据
    c.Data(http.StatusOK, "application/json", jsonData)
}

func main() {
    router := gin.Default()
    // 创建缓存配置实例，这里设置缓存时间为5分钟
    cacheConfig := NewCacheConfig(5 * time.Minute)
    // 使用CacheMiddleware中间件设置缓存策略
    router.Use(CacheMiddleware(cacheConfig))
    // 路由到缓存策略处理器
    router.GET("/cache", CacheHandler)
    // 启动服务
    router.Run(":8080")
}