package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"home_rent/web/controller"
	"home_rent/web/model"
	"home_rent/web/utils"
	"net/http"
	"time"
)

func LoginFilter(ctx *gin.Context) {
	//登录校验
	session := sessions.Default(ctx)
	userName := session.Get("userName")
	resp := make(map[string]interface{})
	if userName == nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		ctx.JSON(http.StatusOK, resp)
		ctx.Abort()
		return
	}

	//计算这个业务耗时
	fmt.Println("next之前打印", time.Now())

	//执行函数
	ctx.Next()

	fmt.Println("next之后打印....")
}

func main() {

	// 初始化Redis链接池
	model.InitRedis()

	// 初始化MySQL连接池
	_, err := model.InitDb()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	// 初始化容器
	store, _ := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte{'j', 'k'})

	// 使用容器
	router.Use(sessions.Sessions("mysession", store)) // 使用中间件，指定容器

	router.Static("/home", "./web/view")

	r1 := router.Group("api/v1.0")
	{
		//路由规范
		r1.GET("/areas", controller.GetArea)
		//r1.GET("/session",controller.GetSession)
		//传参方法,url传值,form表单传值,ajax传值,路径传值
		r1.GET("/imagecode/:uuid", controller.GetImageCd)
		r1.GET("/smscode/:mobile", controller.GetSmscd)
		r1.POST("/users", controller.PostRet)

		//登录业务   路由过滤器   中间件
		r1.Use(sessions.Sessions("mysession", store))
		r1.POST("/sessions", controller.PostLogin)
		//r1.GET("/session", controller.GetSession)
		//路由过滤器   登录的情况下才能执行一下路由请求
		r1.Use(LoginFilter)
		r1.DELETE("/session", controller.DeleteSession)
		r1.GET("/user", controller.GetUserInfo)
		r1.PUT("/user/name", controller.PutUserInfo)

		r1.POST("/user/avatar", controller.PostAvatar)
		r1.POST("/user/auth", controller.PutUserAuth)
		r1.GET("/user/auth", controller.GetUserInfo)
		//获取已发布房源信息
		r1.GET("/user/houses", controller.GetUserHouses)
		//发布房源
		r1.POST("/houses", controller.PostHouses)
		//添加房源图片
		r1.POST("/houses/:id/images", controller.PostHousesImage)
		//展示房屋详情
		r1.GET("/houses/:id", controller.GetHouseInfo)
		//展示首页轮播图
		r1.GET("/house/index", controller.GetIndex)
		//搜索房屋
		r1.GET("/houses", controller.GetHouses)
		//下订单
		r1.POST("/orders", controller.PostOrders)
		//获取订单
		r1.GET("/user/orders", controller.GetUserOrder)
		//同意/拒绝订单
		r1.PUT("/orders/:id/status", controller.PutOrders)
	}

	router.Run(":8084")
}
