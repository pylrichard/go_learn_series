package main

import (
	"fmt"
	"time"
)

func print(i int) {
	fmt.Println(i)
}

func goroutine1() {
	for i := 0; i < 10; i++ {
		/*
			打印结果并不是按照0-9的顺序进行，因为创建的协程并不是按照创建顺序被调度的
			这和协程、内核对象之间的竞争相关，每次打印的结果顺序是随机的
		 */
		go print(i)
	}
	//添加此行代码，避免主goroutine先结束
	time.Sleep(time.Duration(3) * time.Second)
}
