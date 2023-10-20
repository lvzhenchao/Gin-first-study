package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	engine := gin.Default()//里面用到了两个中间件Logger Recovery

	//统一引用
	//engine.Use(RequestInfos())

	engine.GET("/query", RequestInfos(),  func(context *gin.Context) {//也可以单独只用在当前请求内
		fmt.Println("中间件使用方法")
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"message": context.FullPath(),
		})
	})

	engine.GET("/middel", func(context *gin.Context) {

	})

	engine.Run()
}

//打印请求信息的中间件
func RequestInfos() gin.HandlerFunc  {
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method
		fmt.Println("请求路径", path)
		fmt.Println("请求method", method)

		context.Next()
		//获取处理之后的状态
		fmt.Println("状态码：", context.Writer.Status())
	}
}

//自定义中间件标准
//1：func函数
//2：返回类型为HandlerFunc

//作用：请求之前和请求之后处理相关内容

