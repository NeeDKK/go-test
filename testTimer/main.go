package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	queue = make(chan string, 5)
)

func main() {
	e := gin.Default()
	parseStart, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-07-29 10:56:20", time.Local)
	parseEnd, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-07-29 19:00:00", time.Local)
	subStart := parseStart.Sub(time.Now())
	subEnd := parseEnd.Sub(time.Now())
	t1 := time.NewTicker(subStart)
	t2 := time.NewTicker(subEnd)
	go func() {
		for {
			select {
			case <-t1.C:
				fmt.Println("计时器开始")
			case <-t2.C:
				fmt.Println("计时器结束")
			}
		}
	}()
	go func() {
		for {
			select {
			case <-queue:
				fmt.Println("收到停止指令")
				t1.Stop()
				t2.Stop()
			}
		}
	}()
	e.GET("/test", func(context *gin.Context) {
		queue <- "停止"
		context.JSON(200, gin.H{
			"message": "hello",
		})
	})
	e.Run(":8080")
}
