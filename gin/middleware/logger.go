package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

//gin.Default()默认使用了两个中间件Logger和Recovery
//Recovery中间件的作用是保证程序中出现了panic之后，保证服务进程不因一个错误而整个挂掉，否则导致后续请求无法响应
func logger() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.New()
	r.Use(gin.Logger())
	r.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
	_ = r.Run()
}