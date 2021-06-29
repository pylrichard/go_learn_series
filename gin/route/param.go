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

/*
	无论request中的body还是response中的body，都是io.ReadCloser类型
	意味着一旦通过ioutil.ReadAll()全部读取完成，就无法进行第二次读取
	因为在io.ReadCloser内部会有一个标记，记录读取位置，因此一旦读到尾部，就不能再从头读取了
	由于ReadCloser不能Seek，因此解决方法就是把body读出来之后，再重新包装成io.ReadCloser，然后再绑定回body
 */
func printBody(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
	}
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
	ctx.String(http.StatusOK, string(body))
}

func param() {
	r := gin.Default()
	r.Use(printBody)
	//ht -v GET "http://127.0.0.1:8080/order?name=pyl&id=123"
	r.GET("/order", func(ctx *gin.Context) {
		id := ctx.Query("id")
		name := ctx.DefaultQuery("name", "zy")
		ctx.JSON(http.StatusOK, gin.H{
			"id": id,
			"name": name,
		})
	})
	//ht -v POST "http://127.0.0.1:8080/pay" name=pyl id=123
	r.POST("/pay", func(ctx *gin.Context) {
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
		}
		ctx.String(http.StatusOK, string(body))
	})
	/*
		ht -v GET "http://127.0.0.1:8080/user/info?name=pyl&address=chengdu"
		ht -v POST "http://127.0.0.1:8080/user/info" name=pyl address=chengdu
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