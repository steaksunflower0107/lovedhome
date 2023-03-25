package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // "_"表示代码不直接使用包，但是底层链接要使用
	"github.com/jinzhu/gorm"
)

// Student 创建全局结构体(表)
type Student struct {
	Id   int // 默认主键，主键的引入是为了提升查询速度
	Name string
	Age  int
}

// 创建全局连接池句柄
var GlobalConn *gorm.DB

func main() {
	// 连接数据库 -- 格式：用户名：密码@协议(IP:port)/数据库名
	GlobalConn, err := gorm.Open("mysql", "root:sf20020107@tcp(127.0.0.1:3306)/ihome?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库链接失败", err)
		return
	}

	GlobalConn.DB().SetMaxIdleConns(10)
	GlobalConn.DB().SetMaxOpenConns(100)

	GlobalConn.SingularTable(true)

	// 借助gorm创建数据表
	fmt.Println(GlobalConn.AutoMigrate(new(Student)).Error)
}

func InsertData() {
	// 先创建数据 ---> 创建对象
	var stu Student
	stu.Name = "jack"
	stu.Age = 100

	// 插入数据
	fmt.Println(GlobalConn.Create(&stu).Error)
}

func SearchData() {
	var stu Student
	GlobalConn.First(&stu)
	fmt.Println(stu)
}
