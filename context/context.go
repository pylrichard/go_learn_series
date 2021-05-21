package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	/*
		利用Context取消关联任务，Context是任务上下文
		根Context：通过context.Background()创建
		子Context：通过context.WithCancel(parentContent)创建
		当前Context被取消时，基于它的子Context都会被取消
	 */
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 6; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCanceled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 6)
			}
			fmt.Println("canceled: ", i)
		}(i, ctx)
	}
	//当前任务取消回调函数，子任务(5个子协程)中通过ctx.Done()获得父任务被取消的通知消息
	cancel()
	time.Sleep(time.Second * 2)
}

func isCanceled(ctx context.Context) bool {
	//多路选择判断子任务是否该被取消
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}