package main

import (
	"github.com/gin-gonic/gin"
	enLocales "github.com/go-playground/locales/en"
	zhLocales "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	v "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
)

type Teacher struct {
	Name 	string	`form:"name" validate:"required"`
	Age		int		`form:"age" validate:"required,gt=10"`
}

//ht "http://127.0.0.1:8080/user?name=pyl"
//ht "http://127.0.0.1:8080/user?name=pyl&age=2"
//ht "http://127.0.0.1:8080/user?name=pyl&age=12"
func translation() {
	var validator = v.New()
	zh := zhLocales.New()
	en := enLocales.New()
	var translator = ut.New(zh, en)
	r := gin.Default()
	r.GET("/user", func(ctx *gin.Context) {
		//客户端传递来的语言类型
		locale := ctx.DefaultQuery("locale", "zh")
		t, _ := translator.GetTranslator(locale)
		switch locale {
		case "zh":
			_ = zhTranslations.RegisterDefaultTranslations(validator, t)
			break
		case "en":
			_ = enTranslations.RegisterDefaultTranslations(validator, t)
			break
		default:
			_ = zhTranslations.RegisterDefaultTranslations(validator, t)
			break
		}

		var teacher Teacher
		if err := ctx.ShouldBind(&teacher); err != nil {
			ctx.String(http.StatusBadRequest, "err bind: %v", err)
			ctx.Abort()

			return
		}
		if err := validator.Struct(teacher); err != nil {
			errs := err.(v.ValidationErrors)
			var errSlice []string
			for _, e := range errs {
				errSlice = append(errSlice, e.Translate(t))
			}
			ctx.String(http.StatusBadRequest, "reason: %v", errSlice)
			ctx.Abort()

			return
		}
		ctx.String(http.StatusOK, "teacher: %v", teacher)
	})
	_ =  r.Run()
}