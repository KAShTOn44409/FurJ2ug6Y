// 代码生成时间: 2025-08-28 23:21:40
package main

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "strings"

    "github.com/gin-gonic/gin"
)

// FileRenameRequest 定义批量重命名的请求数据结构
type FileRenameRequest struct {
    SourceFolder string   `json:"sourceFolder"`
# NOTE: 重要实现细节
    Mappings    []string `json:"mappings"` // 旧名称:新名称的映射列表
}

// FileRenameResponse 定义批量重命名的响应数据结构
type FileRenameResponse struct {
    Success bool        `json:"success"`
# FIXME: 处理边界情况
    Message string     `json:"message"`
    Files   []string   `json:"files"` // 成功重命名的文件列表
# 添加错误处理
    Errors  []string   `json: