// 代码生成时间: 2025-10-08 19:12:45
package main

import (
    "fmt"
    "net/http"
    "os"
    "log"
    "github.com/gin-gonic/gin"
)

// SyntaxHighlighter 结构体，包含高亮处理的方法
type SyntaxHighlighter struct {
    // 可以添加更多字段来增强功能，比如支持的语言、样式表等
}

// Highlight 接口方法，接受源码并返回高亮后的HTML代码
func (h *SyntaxHighlighter) Highlight(sourceCode string, language string) (string, error) {
    // 这里应该是复杂的高亮逻辑，但为了示例简单，我们直接返回源码
    // 实际应用中可能需要调用外部库来实现代码高亮
    return "<pre><span style='color:blue'>" + sourceCode + "</span></pre>", nil
}

func main() {
    router := gin.Default()

    // 错误处理中间件
    router.Use(func(c *gin.Context) {
        c.Next()
        if len(c.Errors) > 0 {
            // 处理Gin框架产生的所有错误
            for _, err := range c.Errors {
                c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
                    "error": err.Err().Error(),
                })
            break
            }
        }
    })

    highlighter := &SyntaxHighlighter{}

    // 路由：接收POST请求，处理代码高亮
    router.POST("/highlight", func(c *gin.Context) {
        var req struct {
            SourceCode string `json:"source_code"`
            Language   string `json:"language"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid request parameters",
            })
            return
        }

        // 调用高亮处理方法
        highlightedCode, err := highlighter.Highlight(req.SourceCode, req.Language)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to highlight code",
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "highlighted_code": highlightedCode,
        })
    })

    // 启动服务器
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
