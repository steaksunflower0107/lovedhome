package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	// 1.链接数据库
	Conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis Dial 错误", err)
	}
	defer Conn.Close()
	// 操作数据库
	reply, err := Conn.Do("set", "married", "false")
	// 回复助手函数 ---> 确定成具体的数据类型
	ret, e := redis.String(reply, err)
	fmt.Println(ret, e)
}
