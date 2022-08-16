package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"test/socket/utils"
)

func main() {
	engine := gin.Default()
	us := utils.NewUnixSocket("/data/socket/testAAA", 1024)
	//设置服务端接收处理
	us.SetContextHandler(func(context string) string {
		defer func() {
			//捕获test抛出的panic
			if err := recover(); err != nil {
				fmt.Println("IP-MAC绑定发生错误", err)
			}
		}()
		fmt.Println("testAAA收到string================:" + context)
		fmt.Println(context)
		testAAA(context)
		return "ok"
	})
	//开始服务
	go us.StartServer()
	engine.Run()
}

func testAAA(context string) {
	split := strings.Split(context, ",")
	s := split[10]
	fmt.Println(s)
}
