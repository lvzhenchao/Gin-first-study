package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {

	engine := gin.Default()

	//1 通用请求
		//http://localhost:8080/hello?name=davie
		engine.Handle("GET", "/hello", func(context *gin.Context) {
			//获取请求路径
			path := context.FullPath()
			fmt.Println(path)

			//获取请求参数
			name := context.DefaultQuery("name", "test")
			fmt.Println(name)

			//输出
			context.Writer.Write([]byte("hello," + name))
		})

		//post请求
		//http://localhost:8080/login
		engine.Handle("POST", "/login", func(context *gin.Context) {
			//获取请求路径
			path := context.FullPath()
			fmt.Println(path)

			//获取请求参数
			username := context.PostForm("username")
			password := context.PostForm("password")
			fmt.Println(username)
			fmt.Println(password)

			context.Writer.Write([]byte(username + " 登录"))
		})

	//2 分类请求
		//get
		engine.GET("/hello1", func(context *gin.Context) {

		})

		//post
		engine.POST("/login1", func(context *gin.Context) {

		})



	engine.Run(":8080")
}


//Handle直接使用Handle方法进行http请求的处理

//context 上下文对象，简单说为我们提供操作环境
//context.Query和context.DefaultQuery获取get请求携带的参数
//context.Writer.write向请求发起端返回数据

//context.PostForm获取表单提交的数据字段
//context.getpostform获取表单数据（有两个返回参数 err）