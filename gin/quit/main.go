package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/user", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello user")
	})

	s := &http.Server{
		Addr: "127.0.0.1:8080",
		Handler: r,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil &&
			err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("shutdown...")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancelFunc()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Println("server exit")
}