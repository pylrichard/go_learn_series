package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ipAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ipList := []string {
			"127.0.0.1",
		}
		f := false
		clientIP := ctx.ClientIP()
		for _, ip := range ipList {
			if ip == clientIP {
				f = true
				break
			}
		}
		if !f {
			ctx.String(http.StatusBadRequest, "client no permission")
			ctx.Abort()
		}
	}
}

func ipAuth() {
	r := gin.Default()
	r.Use(ipAuthMiddleware())
	r.GET("/user", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ip auth")
	})
	_ = r.Run()
}