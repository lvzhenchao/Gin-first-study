package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {

	engine := gin.Default()

	engine.LoadHTMLGlob("./html/*")

	engine.Static("/img", "./img")
	engine.GET("/hellohtml", func(context *gin.Context) {
		fullpath := context.FullPath()
		fmt.Println(fullpath)

		context.HTML(http.StatusOK, "index.html", gin.H{
			"fullpath": fullpath,
			"title": "Gin 教程",
		})
	})

	engine.Run()
}

//加载静态文件, 需要设置html目录  engine.LoadHTMLGlob

//加载静态资源文件, engine.static() 将第二个参数路径映射到第一个参数上