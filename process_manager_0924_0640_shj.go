// 代码生成时间: 2025-09-24 06:40:56
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "os/exec"
    "time"

    "github.com/gin-gonic/gin"
)

// ProcessManager provides functionalities to manage processes.
type ProcessManager struct {
    // Any fields needed for the process manager
}

// NewProcessManager creates a new instance of ProcessManager.
func NewProcessManager() *ProcessManager {
    return &ProcessManager{}
}

// StartProcess starts a new process with the given command.
func (pm *ProcessManager) StartProcess(c *gin.Context, command string) {
    // Split the command into parts if needed
    args := []string{command}
    // Execute the command
    process, err := os.StartProcess(command, args, &os.ProcAttr{
        Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("Failed to start process: %s", err),
        })
        return
    }
    // Wait for the process to finish and get its exit code
    state, err := process.Wait()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("Failed to wait for process: %s", err),
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "exit_code": state.ExitCode(),
    })
}

// KillProcess kills a running process by its PID.
func (pm *ProcessManager) KillProcess(c *gin.Context, pid int) {
    // Find the process with the given PID
    process, err := os.FindProcess(pid)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("Process with PID %d not found: %s", pid, err),
        })
        return
    }
    // Kill the process
    err = process.Kill()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("Failed to kill process: %s", err),
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "Process killed successfully",
    })
}

func main() {
    r := gin.Default()

    // Use middleware to handle logging and recovery
    r.Use(gin.Recovery())
    r.Use(gin.Logger())

    pm := NewProcessManager()

    // Define routes
    r.POST("/start", func(c *gin.Context) {
        command := c.PostForm("command")
        pm.StartProcess(c, command)
    })
    r.POST("/kill", func(c *gin.Context) {
        pid := c.PostForm("pid")
        pidInt, err := strconv.Atoi(pid)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": fmt.Sprintf("Invalid PID: %s", err),
            })
            return
        }
        pm.KillProcess(c, pidInt)
    })

    // Start the server
    log.Fatal(r.Run(":8080"))
}
