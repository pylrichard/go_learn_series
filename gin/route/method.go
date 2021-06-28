package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func method() {
	r := gin.Default()
	r.Handle("GET", "/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	//支持所有类型的请求
	r.Any("/any", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "any",
		})
	})
	//从url获取参数
	//ht "http://127.0.0.1:8080/pyl/123"
	r.GET("/:name/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name": ctx.Param("name"),
			"id": ctx.Param("id"),
		})
	})
	//泛绑定可以让前缀相同的路由使用同一个处理逻辑
	//http://127.0.0.1:8080/user/"符合接口路由
	//http://127.0.0.1:8080/user"不符合接口路由
	//r.GET("/user/*action", func(ctx *gin.Context) {
	//	ctx.String(http.StatusOK, "hello gin")
	//})
	_ = r.Run()
}