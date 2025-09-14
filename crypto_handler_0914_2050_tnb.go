// 代码生成时间: 2025-09-14 20:50:04
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"

    "github.com/gin-gonic/gin"
)

// CryptoHandler 处理密码加密和解密请求
type CryptoHandler struct{}

// NewCryptoHandler 创建一个新的CryptoHandler实例
func NewCryptoHandler() *CryptoHandler {
    return &CryptoHandler{}
}

// EncryptHandler 处理加密请求
func (h *CryptoHandler) EncryptHandler(c *gin.Context) {
    var req struct{
        Password string `json:"password" binding:"required"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    encrypted, err := encrypt(req.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Encryption failed"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"encrypted": hex.EncodeToString(encrypted)})
}

// DecryptHandler 处理解密请求
func (h *CryptoHandler) DecryptHandler(c *gin.Context) {
    var req struct{
        Encrypted string `json:"encrypted" binding:"required"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    encrypted, err := hex.DecodeString(req.Encrypted)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid encrypted string"})
        return
    }
    password, err := decrypt(encrypted)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Decryption failed"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"password": password})
}

// encrypt 执行加密操作
func encrypt(plaintext string) ([]byte, error) {
    block, err := aes.NewCipher([]byte("your-very-secret-key-32"))
    if err != nil {
        return nil, err
    }
    
    plaintextBytes := []byte(plaintext)
    blockSize := block.BlockSize()
    padding := blockSize - len(plaintextBytes)%blockSize
    plaintextBytes = append(plaintextBytes, bytes.Repeat([]byte{byte(padding)}, padding)...)
    
    cipherText := make([]byte, aes.BlockSize+len(plaintextBytes))
    iv := cipherText[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }
    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(cipherText[aes.BlockSize:], plaintextBytes)
    return cipherText, nil
}

// decrypt 执行解密操作
func decrypt(cipherText []byte) (string, error) {
    block, err := aes.NewCipher([]byte("your-very-secret-key-32"))
    if err != nil {
        return "", err
    }
    
    if len(cipherText) < aes.BlockSize {
        return "", errors.New("ciphertext too short")
    }
    
    iv := cipherText[:aes.BlockSize]
    cipherText = cipherText[aes.BlockSize:]
    mode := cipher.NewCBCDecrypter(block, iv)
    mode.CryptBlocks(cipherText, cipherText)
    
    // Unpad
    padding := int(cipherText[len(cipherText)-1])
    if padding < 1 || padding > aes.BlockSize {
        return "", errors.New("invalid padding")
    }
    cipherText = cipherText[:len(cipherText)-padding]
    return string(cipherText), nil
}

func main() {
    r := gin.Default()
    
    // 创建CryptoHandler实例
    cryptoHandler := NewCryptoHandler()
    
    // 注册加密处理器
    r.POST("/encrypt", cryptoHandler.EncryptHandler)
    
    // 注册解密处理器
    r.POST("/decrypt", cryptoHandler.DecryptHandler)
    
    // 启动服务器
    r.Run() // listen and serve on 0.0.0.0:8080
}
