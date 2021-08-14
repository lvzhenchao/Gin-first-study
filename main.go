package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main()  {
	//创建Engine
	engine := gin.Default()

	//处理get请求
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("请求路径：", context.FullPath())
		context.Writer.Write([]byte("hello, Gin\n"))
	})

	//启动运行
	if err := engine.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

//可以在命令行使用： curl http://localhost:8080/hello
//engine.Run()可以自定义端口号，默认8080

//engine代表gin框架的一个结构体定义，包含路由组，中间件、页面渲染接口、框架配置等
//两种方式
//gin.Default()
//gin.New()
//区别
//gin.Default也使用gin.New创建实例，但是会默认使用Logger和Recovery中间件
//logger负责打印输出日志的中间件，方便开发者进行程序调试
//recovery中间件作用是程序执行过程中遇到panic中断了服务，则recovery会恢复程序执行，并返回服务器500内部错误
//通常我们使用Default创建engine实例









