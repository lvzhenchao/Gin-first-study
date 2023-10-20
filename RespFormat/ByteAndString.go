package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	engine := gin.Default()

	//返回字节切片格式
	engine.GET("/hellobyte", func(context *gin.Context) {
		fullPath := "请求路径 " + context.FullPath()
		fmt.Println(fullPath)
		context.Writer.Write([]byte(fullPath))
	})

	//返回字符串格式
	engine.GET("/hellostring", func(context *gin.Context) {
		fullPath := "请求路径 " + context.FullPath()
		fmt.Println(fullPath)
		context.Writer.WriteString(fullPath)
	})


	engine.Run()
}

//context.Writer 是gin框架封装的一个ResponseWriter接口类型

//返回数据格式[]byte

