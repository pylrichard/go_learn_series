package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	form用来指定结构体参数与请求参数的对应关系
	binding指定验证规则
	binding中可以指定多个条件，条件之间或关系使用|竖线分割，与关系使用逗号分割
 */
type Person struct {
	Age		int		`form:"age" binding:"required,gt=10"`
	Name	string	`form:"name" binding:"required"`
}

//ht "http://localhost:8080/user?name=pyl&age=11"
func standard() {
	r := gin.Default()
	r.GET("/user", func(ctx *gin.Context) {
		var p Person
		if err := ctx.ShouldBind(&p); err != nil {
			ctx.String(http.StatusBadRequest, "params err: %v", err)
			ctx.Abort()

			return
		}
		ctx.String(http.StatusOK, "params: %v", p)
	})
	//默认端口8080
	_ = r.Run()
}