package resp

import (
	"fmt"
	"runtime"
	"time"
)

var num = 8

func FirstResp() {
	fmt.Println("before: ", runtime.NumGoroutine())

	ch := runMultipleTask(num)
	fmt.Println(<-ch)

	time.Sleep(time.Second * 2)
	fmt.Println("after: ", runtime.NumGoroutine())
}

func AllResp() {
	fmt.Println("before: ", runtime.NumGoroutine())

	ch := runMultipleTask(num)
	result := ""
	/*
		由于按顺序启动的协程并不意味着会按照顺序完成任务，由于CPU资源抢占，协程完成任务的顺序是随机的
		那么消息在channel中的顺序也是随机的，按顺序读取channel中的结果，就是协程完成任务的顺序
	 */
	for i := 0; i < num; i++ {
		result += <-ch + "\n"
	}
	fmt.Println(result)

	time.Sleep(time.Second * 2)
	fmt.Println("after: ", runtime.NumGoroutine())
}

func runMultipleTask(num int) chan string {
	/*
		创建一个非Buffer Channel，然后启动8个协程，每个协程独立处理任务
		任务处理完毕之后，将结果放置到channel中
		runMultipleTask()会阻塞读取channel中的内容，从channel中只读取首个消息
		后面7个协程会阻塞等待消息接收者回到channel上，才会发送消息，从而解除协程的阻塞状态
		因此在程序运行后会有7个协程一直被挂起，无法释放资源
		在高并发情况下会导致大量资源被占用

		使用Buffer Channel，并将capacity设置为协程的数量
		这样就可以避免只读取首个消息之后，其他协程阻塞等待
	 */
	ch := make(chan string, num)
	for i := 0; i < num; i++ {
		go func(id int, ch chan string) {
			ret := runTask(id)
			ch <-ret
		}(i, ch)
	}

	return ch
}

func runTask(id int) string {
	time.Sleep(time.Millisecond * 6)

	return fmt.Sprintf("task %d", id)
}