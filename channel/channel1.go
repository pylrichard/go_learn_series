package main

import (
	"fmt"
	"sync"
)

/*
	dataProducer1和dataReceiver1通过channel发送数据时，需要事先约定好发送/接收的数量
	以此保证数据沟通的完整性
 */
func dataProducer1(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
		}
		wg.Done()
	}()
}

/*
	要进行解耦可以采取：
	1 可以约定一个特殊的数据标记(比如-1)，当数据写入完毕时，dataProducer就往channel中放入-1
	表示数据已经发送完毕。但要是有好几个数据接收者，那就需要放入多个结束标记，这样显然不友好
	2 最合适的方式是，当dataProducer发送完数据之后，关闭channel
	而dataReceiver在从channel中获取数据时，同时判断channel的状态
	要是channel已关闭，说明数据发送完毕。dataProducer在发送完数据之后关闭channel
 */
func dataProducer2(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 6; i++ {
			ch <- i
		}
		//发送完毕，关闭channel
		close(ch)
		wg.Done()
	}()
}

/*
	对于消息接收者，要是channel被关闭，而消息接收者还在channel上接收数据，不会抛出panic
	那么channel会立即返回消息类型的零值
 */
func dataReceiver1(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 8; i++ {
			//从已经关闭的channel读取数据并不会抛出panic
			//向已经关闭的channel放入消息，会抛出panic
			data := <-ch
			fmt.Println(data)
		}
		wg.Done()
	}()
}

func channel1() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	//dataProducer1(ch, &wg)
	dataProducer2(ch, &wg)
	wg.Add(1)
	dataReceiver1(ch, &wg)
	wg.Wait()
}