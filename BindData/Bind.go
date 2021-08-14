package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Student struct {
	Name string `form:"name"`
	Classes string `form:classes`
}

type Register struct {
	Name string `form:"name"`
	Phone string `form:"phone"`
}


func main()  {

	engine := gin.Default()

	//http://localhost:8080/hello?name=dave&classes=软件工程
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var student Student
		err := context.ShouldBindQuery(&student)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(student.Name)
		fmt.Println(student.Classes)

		context.Writer.Write([]byte("hello "+student.Name))
	})


	engine.POST("/register", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var register  Register
		err := context.ShouldBind(&register)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(register.Name)
		fmt.Println(register.Phone)

		context.Writer.Write([]byte(register.Name + " 注册"))
	})

	engine.Run()
}

//表单实体绑定
//ShouldBind
//通过tag标签
//context.BindJSON绑定json格式的数据
//context.BindXML绑定xml格式的数据
