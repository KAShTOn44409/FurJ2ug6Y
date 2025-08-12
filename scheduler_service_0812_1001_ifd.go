// 代码生成时间: 2025-08-12 10:01:11
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/gin-gonic/gin"
)

// SchedulerService 定时任务调度器服务
type SchedulerService struct {
    jobDuration time.Duration
    jobInterval time.Duration
}

// NewSchedulerService 创建一个新的调度器服务
func NewSchedulerService(jobDuration, jobInterval time.Duration) *SchedulerService {
    return &SchedulerService{
        jobDuration: jobDuration,
        jobInterval: jobInterval,
    }
}

// Start 开始执行定时任务
func (s *SchedulerService) Start() {
    ticker := time.NewTicker(s.jobInterval)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            err := s.ExecuteJob()
            if err != nil {
                log.Printf("Error executing job: %v", err)
            }
        }
    }
}

// ExecuteJob 执行定时任务
func (s *SchedulerService) ExecuteJob() error {
    // 模拟任务执行，实际项目中应替换为具体业务逻辑
    fmt.Println("Executing job...")
    time.Sleep(s.jobDuration)
    fmt.Println("Job executed.")

    return nil
}

func main() {
    // 创建定时任务调度器
    scheduler := NewSchedulerService(5*time.Second, 10*time.Second)
    go scheduler.Start()

    // 设置Gin中间件
    r := gin.Default()
    r.Use(gin.Recovery())

    // 定义定时任务调度器的路由
    r.GET("/start-scheduler", func(c *gin.Context) {
        // 启动定时任务调度器
        scheduler.Start()
        c.JSON(200, gin.H{
            "message": "Scheduler started",
        })
    })

    // 启动Gin服务器
    log.Fatal(r.Run(":8080"))
}
