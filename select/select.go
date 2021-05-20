package main

import (
	"fmt"
	"time"
)

func testSelect() {
	retChan1 := retChannel(10, "one")
	retChan2 := retChannel(10, "two")
	select {
	/*
		执行select时要是所有的case都处于非阻塞状态，不能根据case的顺序判断会执行哪一个case
	 */
	case ret := <-retChan1:
		fmt.Println("One: " + ret)
	case ret := <-retChan2:
		fmt.Println("Two: " + ret)
	}
}

func testTimeout() {
	retChan := retChannel(500, "ok")
	select {
	case ret := <-retChan:
		fmt.Println("Result: " + ret)
	/*
		阻塞控制retChan的超时时间，当程序运行到select时，两个case都处于阻塞状态
		100ms之后第二个case的Channel会转为非阻塞状态，从而控制第一个case的Channel的超时时长
	 */
	case <-time.After(time.Millisecond * 100):
		fmt.Println("Timeout")
	}
}

func retChannel(t int, data string) chan string {
	retChan := make(chan string)
	go func() {
		time.Sleep(time.Millisecond * time.Duration(t))
		retChan <- data
	}()

	return retChan
}

func main() {
	testSelect()
	testTimeout()
}