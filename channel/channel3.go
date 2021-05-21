package main

import (
	"fmt"
	"time"
)

func channel3() {
	cancelCh := make(chan struct{})
	for i := 0; i < 6; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCanceled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 8)
			}
			fmt.Println("canceled: ", i)
		}(i, cancelCh)
	}
	//通过channel通知协程取消任务
	cancel1(cancelCh)
	// cancel2(cancelCh)
	time.Sleep(time.Second * 2)
}

//多路选择机制从channel中读取取消任务的消息
func isCanceled(cancelCh chan struct{}) bool {
	select {
	//cancelCh一直处于阻塞状态，一旦有消息返回true
	case <-cancelCh:
		return true
	default:
		return false
	}
}

//发送取消任务的消息
func cancel1(cancelCh chan struct{}) {
	/*
		向cancelCh发送一个消息，只有一个协程被取消
		要通知所有协程取消任务，就需要知道协程总数，然后发送对应条数消息，这会导致代码耦合
	 */
	for i := 0; i < 2; i++ {
		cancelCh <- struct{}{}
	}
}

func cancel2(cancelCh chan struct{}) {
	//channel关闭时，阻塞等待channel的所有接收者会立即返回，通知所有协程取消任务
	close(cancelCh)
}