package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main()  {
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
