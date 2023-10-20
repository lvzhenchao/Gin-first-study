package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	engine := gin.Default()

	routerGroup := engine.Group("/user")

	routerGroup.POST("/register", registerHandle)

	routerGroup.DELETE("/:id", loginHandle)

	engine.Run()
}

func registerHandle(context *gin.Context)  {
	fullpath := "用户注册：" + context.FullPath()
	fmt.Println(fullpath)
	context.Writer.WriteString(fullpath)
}

func loginHandle(context *gin.Context)  {
	fullpath := "用户删除：" + context.FullPath()
	fmt.Println(fullpath)
	userID := context.Param("id")
	context.Writer.WriteString(fullpath + "" +userID)
}

//多路由分组开发及相关开发技巧