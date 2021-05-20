package main

import (
	"fmt"
	"time"
)

/*
	阻塞式通信Channel要保证通信双方都在Channel上，通信才会完成
	否则先到Channel的一方需要阻塞等待另一方到Channel上，才会完成本次通信
	双方通信实体才会继续执行之后的任务
	asyncServiceTask中的子协程执行完任务之后(打印asyncServiceTask getResult)
	之后立即将结果放到Channel中
	但由于主协程(接收者)还在阻塞执行其他任务，不在Channel上
	所以子协程(发送者)会被一直阻塞等待在向Channel放入消息的代码retChan <- ret
	直到主协程通过<-retCh从Channel中获取到消息时(打印serviceTask done)之后
	子协程才被释放继续执行(打印asyncServiceTask done)
 */
func blockChannel() chan string {
	fmt.Println("blockChannel")

	return make(chan string)
}

/*
	asyncServiceTask在获得serviceTask的结果之后(打印asyncServiceTask getResult)
	由于使用了Buffer Channel，子协程(发送者)立即将结果放入到Channel中
	而此时主协程(接收者)还在执行otherTask，主协程不在Channel上但并没有造成子协程的阻塞等待
	当主协程完成otherTask之后，通过Channel获取到子协程早已放好的消息(打印serviceTask done)
 */
func bufferChannel() chan string {
	fmt.Println("bufferChannel")

	return make(chan string, 1)
}

func asyncServiceTask(isBlocked bool) chan string {
	var retChan chan string
	if isBlocked {
		retChan = blockChannel()
	} else {
		retChan = bufferChannel()
	}
	go func() {
		ret := serviceTask()
		fmt.Println("asyncServiceTask getResult")
		retChan <- ret
		fmt.Println("asyncServiceTask done")
	}()

	return retChan
}

func serviceTask() string {
	time.Sleep(time.Millisecond * 50)

	return "serviceTask done"
}

func otherTask() {
	fmt.Println("otherTask working...")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("otherTask done")
}

func testService() {
	//顺序串行执行
	fmt.Println(serviceTask())
	otherTask()
}

func testAsyncService() {
	//经过协程并发处理的两个Task执行的总时间缩短了
	retChan := asyncServiceTask(false)
	otherTask()
	fmt.Println(<-retChan)
}

func main() {
	testService()
	testAsyncService()
	time.Sleep(time.Millisecond * 1000)
}