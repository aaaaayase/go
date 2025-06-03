package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 中间件-拦截器
func myHandler() (handlerFunc gin.HandlerFunc) {
	return func(c *gin.Context) {
		c.Set("name", "yun")
		// 放行
		c.Next()
		// 拦截
		//c.Abort()
	}
}

func main() {
	// 创建服务
	server := gin.Default()

	// 拦截器全局
	server.Use(myHandler())

	// 加载静态页面
	server.LoadHTMLFiles("templates/index.html")
	// 加载资源
	server.Static("/static", "./static")

	// 接口
	server.GET("/hell0", myHandler(), func(context *gin.Context) {
		name := context.MustGet("name")
		context.JSON(200, gin.H{"message": "hello world", "name": name})
	})

	server.POST("/user", func(context *gin.Context) {

		context.JSON(200, gin.H{"message": "user go!!!"})
	})

	// 响应一个前端页面
	server.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{"msg": "后台传的数据"})
	})

	// 从查询字符串获取参数
	server.GET("/user/login", func(context *gin.Context) {

		username := context.Query("username")
		password := context.Query("password")

		context.JSON(http.StatusOK, gin.H{"username": username, "password": password})

	})

	server.GET("user/logout/:username/:password", func(context *gin.Context) {
		username := context.Params.ByName("username")
		password := context.Params.ByName("password")
		context.JSON(http.StatusOK, gin.H{"username": username, "password": password})
	})

	// 前端向后端传递json
	server.POST("/json", func(context *gin.Context) {
		data, _ := context.GetRawData()
		var m map[string]interface{}
		json.Unmarshal(data, &m)

		context.JSON(http.StatusOK, m)

	})

	// 传递表单
	server.POST("form", func(context *gin.Context) {

		username := context.PostForm("username")
		context.JSON(http.StatusOK, gin.H{"username": username})

	})

	// 路由
	server.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusTemporaryRedirect, "http://www.baidu.com")
	})

	// 404
	server.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.html", nil)

	})

	// 路由组 /user/add
	userGroup := server.Group("/user")
	{
		userGroup.GET("/add")
		userGroup.GET("/logout")
		userGroup.GET("/delete")

	}

	// order路由组
	orderGroup := server.Group("order")
	{
		orderGroup.GET("/add")
		orderGroup.GET("/login")
	}

	// 端口
	server.Run(":8080")
}
