// 代码生成时间: 2025-09-22 08:19:55
package main

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
)

// xssEscape escapes HTML special characters to prevent XSS attacks.
func xssEscape(str string) string {
    return strings.NewReplacer(
        "&", "&amp;",
        "<", "&lt;",
        ">", "&gt;",
        """", "&quot;", // Double quotes
        "'", "&#x27;", // Single quotes
    ).Replace(str)
}

// XssProtectionMiddleware is a Gin middleware function that prevents XSS attacks.
func XssProtectionMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Check if request method is GET, as we only need to protect against XSS on HTML responses
        if c.Request.Method == http.MethodGet {
            c.Next()
            return
        }

        // Wrap the original writer with a new ResponseWriter that will escape HTML
        originalWriter := c.Writer
        c.Writer = NewResponseWriter(originalWriter, xssEscape)

        // Continue the request chain
        c.Next()
    }
}

// NewResponseWriter creates a new ResponseWriter that escapes HTML special characters.
func NewResponseWriter(originalWriter gin.ResponseWriter, escapeFn func(string) string) gin.ResponseWriter {
    return &xssResponseWriter{
        originalWriter: originalWriter,
        escapeFn:      escapeFn,
    }
}

// xssResponseWriter is a ResponseWriter that escapes HTML special characters.
type xssResponseWriter struct {
    originalWriter gin.ResponseWriter
    escapeFn      func(string) string
}

// Write implements the ResponseWriter interface. It escapes HTML special characters before writing to the original writer.
func (w *xssResponseWriter) Write(b []byte) (int, error) {
    str := string(b)
    escapedStr := w.escapeFn(str)
    return w.originalWriter.Write([]byte(escapedStr))
}

// WriteString implements the ResponseWriter interface. It escapes HTML special characters before writing to the original writer.
func (w *xssResponseWriter) WriteString(s string) (int, error) {
    escapedStr := w.escapeFn(s)
    return w.originalWriter.WriteString(escapedStr)
}

// SetCookie implements the ResponseWriter interface.
func (w *xssResponseWriter) SetCookie(name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool) {
    w.originalWriter.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
}

// Header implements the ResponseWriter interface.
func (w *xssResponseWriter) Header() http.Header {
    return w.originalWriter.Header()
}

// Status implements the ResponseWriter interface.
func (w *xssResponseWriter) Status() int {
    return w.originalWriter.Status()
}

// Size implements the ResponseWriter interface.
func (w *xssResponseWriter) Size() int {
    return w.originalWriter.Size()
}

// Before implements the ResponseWriter interface.
func (w *xssResponseWriter) Before(c *gin.Context) {
    w.originalWriter.Before(c)
}

// After implements the ResponseWriter interface.
func (w *xssResponseWriter) After(c *gin.Context) {
    w.originalWriter.After(c)
}

func main() {
    router := gin.New()
    router.Use(XssProtectionMiddleware())

    // Define a simple route that could be vulnerable to XSS
    router.POST("/xss", func(c *gin.Context) {
        input := c.PostForm("input")
        // Normally, you would use template rendering and automatically escape HTML
        // For demonstration purposes, we are directly writing to the response
        c.String(http.StatusOK, "Received input: " + input)
    })

    router.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Page not found",
        })
    })

    router.Run()
}
