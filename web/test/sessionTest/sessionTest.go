package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 初始化容器
	store, _ := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte{'j', 'k'})
	// 使用容器
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/test", func(context *gin.Context) {
		// 调用session, 设置session数据
		session := sessions.Default(context)
		// 设置session
		session.Set("sessionTest", "rose")
		// 修改session时，需要执行Save，否则不会生效
		session.Save()
	})

	router.Run(":9999")
}
