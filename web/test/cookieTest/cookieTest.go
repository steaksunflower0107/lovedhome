package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/test", func(context *gin.Context) {
		// 设置 Cookie
		context.SetCookie("myTest", "value", 60*60, "", "", false, true)
		cookieVal, _ := context.Cookie("myTest")

		fmt.Println("获取的Cookie为", cookieVal)

		context.Writer.WriteString("测试 Cookie...")
	})

	router.Run(":9999")
}
