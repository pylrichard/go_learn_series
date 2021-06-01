package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/*
	一个请求需要处理多个相对独立的任务，每个任务都有可能超时，需要将超时时间控制在6s内
	等所有任务返回之后，将结果返回给前端
	这个需求可以使用协程完成，每个任务使用协程处理，使用sync.WaitGroup保证所有任务完成之后再继续处理
	为了对单个任务进行超时控制，可以使用多路选择机制
 */
func timeout(duration int) string {
	t := 6 * time.Second
	ch := runTask(duration)
	select {
	case ret := <-ch:
		return ret
	case <-time.After(t):
		return "timeout in " + strconv.Itoa(duration) + "s"
	}
}

func runTask(duration int) chan string {
	ch := make(chan string)
	go func(duration int) {
		time.Sleep(time.Duration(duration) * time.Second)
		ch <- "success in " + strconv.Itoa(duration) + "s"
	}(duration)

	return ch
}

func goroutine2() {
	tasksTime := []int{3, 10, 4, 2, 1, 3}
	count := len(tasksTime)
	var wg sync.WaitGroup
	ch := make(chan string, len(tasksTime))
	for i := 0; i < len(tasksTime); i++ {
		wg.Add(1)
		go func(taskTime int, ch chan string) {
			result := timeout(taskTime)
			ch <- result
			wg.Done()
		}(tasksTime[i], ch)
	}
	wg.Wait()
	results := make([]string, count)
	for i := 0; i < count; i++ {
		ret := <-ch
		fmt.Println("result: " + ret)
		results[i] = ret
	}
	fmt.Println(results)
}