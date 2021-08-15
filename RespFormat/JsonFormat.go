package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int
	Message string
	Data interface{}
}

func main()  {
	engine := gin.Default()

	//返回json格式（map类型）
	engine.GET("/hellojson", func(context *gin.Context) {
		fullpath := "请求地址：" + context.FullPath()
		fmt.Println(fullpath)
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"message": "ok",
			"data": fullpath,
		})
	})

	//返回json格式（结构体）
	engine.GET("/hellostruct", func(context *gin.Context) {
		fullpath := "请求地址：" + context.FullPath()
		fmt.Println(fullpath)

		resp := Response{Code: 1, Message: "ok", Data: fullpath}

		context.JSON(200, resp)
	})


	engine.Run()
}

//注意code码需是200，这个码是network 中Headers的Status Code
