package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type Booking struct {
	CheckIn		int `form:"check_in" binding:"required,gt=0,dateCheck"`
	//check_out必须大于check_in，binding规则中自定义验证指定验证器名称
	CheckOut	int	`form:"check_out" binding:"required,dateCheck,gtfield=CheckIn"`
}

//更多验证规则见https://pkg.go.dev/github.com/go-playground/validator/v10
func custom() {
	r := gin.Default()
	r.GET("/book", func(ctx *gin.Context) {
		var b Booking
		if err := ctx.ShouldBind(&b); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()

			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg":	"ok",
			"in":	b.CheckIn,
			"out":	b.CheckOut,
		})
	})

	var v, ok = binding.Validator.Engine().(*validator.Validate)
	if ok {
		//验证器和验证方法关联
		_ = v.RegisterValidation("dateCheck", customDateCheck)
	}
	_ = r.Run()
}

//需要同时满足结构体验证规则和自定义验证规则，才能够获取到正常的数据返回
//ht "http://127.0.0.1:8080/book?check_in=0&check_out=1"
//ht "http://127.0.0.1:8080/book?check_in=9999999997&check_out=9999999999"
func customDateCheck(field validator.FieldLevel) bool {
	now := int(time.Now().Unix())
	v := field.Field().Interface().(int)
	if v >= now {
		return true
	}

	return false
}