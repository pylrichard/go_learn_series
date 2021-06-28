package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Name	string `form:"name"`
	Address string `form:"address"`
}

func param() {
	r := gin.Default()
	r.GET("/order", func(ctx *gin.Context) {
		id := ctx.Query("id")
		name := ctx.DefaultQuery("name", "zy")
		ctx.JSON(http.StatusOK, gin.H{
			"id": id,
			"name": name,
		})
	})
	//获取body
	//ht -v POST 'http://127.0.0.1:8080/pay' name=pyl id=123
	r.POST("/pay", func(ctx *gin.Context) {
		//通过ioutil.ReadAll()读取Request的Body之后，通过PostForm()无法获取POST参数
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
		}
		ctx.String(http.StatusOK, string(body))
		//要是想改变这种情况可以再把数据写回到Body中
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		id := ctx.PostForm("id")
		name := ctx.DefaultPostForm("name", "unknown")
		ctx.String(http.StatusOK, "%s,%s", id, name)
	})
	/*
		bind获取参数
		ht GET "http://127.0.0.1:8080/user/info?name=pyl&addr=chengdu"
		ht -v POST "http://127.0.0.1:8080/user/info" name=pyl addr=chengdu
	 */
	r.GET("/user/info", bindUserInfo)
	r.POST("/user/info", bindUserInfo)
	_ = r.Run()
}

func bindUserInfo(ctx *gin.Context) {
	//bind根据当前请求的context-type区分该如何获取请求中的参数
	var p Person
	if err := ctx.ShouldBind(&p); err != nil {
		ctx.String(http.StatusBadRequest, "bind err: %v", err)
		ctx.Abort()
	}
	ctx.String(http.StatusOK, "bind user info: %v", p)
}