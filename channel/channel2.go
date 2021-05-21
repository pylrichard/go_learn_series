package main

import (
	"fmt"
	"sync"
)

func dataReceiver2(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			//所有channel接收者从channel中阻塞读取消息时，channel会返回两个数据
			//第二个数据可以用来判断channel当前是否处于关闭状态
			//channel关闭时ok为false，data为消息类型的零值
			//这种机制常被用来向多个订阅者发送关闭消息
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				fmt.Println("channel closed")
				break
			}
		}
		wg.Done()
	}()
}

func channel2() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer2(ch, &wg)
	wg.Add(1)
	dataReceiver2(ch, &wg)
	wg.Wait()
}