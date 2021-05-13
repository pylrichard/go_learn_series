package main

import (
	"fmt"
	"time"
)

type f func(opt int) int

func slowFunc(opt int) int {
	time.Sleep(time.Duration(opt) * time.Second)

	return opt
}

func calcFuncRunTime(inner f) f {
	return func(opt int) int {
		start := time.Now()
		ret := inner(opt)
		fmt.Println("func time:", time.Since(start).Seconds())

		return ret 
	}
}

func useFuncAsParam(opt int) {
	sf := calcFuncRunTime(slowFunc)
	sf(opt)
}

func variableLenParam() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(4, 5))
}

func sum(ops ...int) int {
	result := 0
	for _, v := range ops {
		result += v
	}

	return result
}

func clear() {
	fmt.Println("func finished")
}

func useDefer() {
	defer clear()
	fmt.Println("func is running")
	panic("panic")
}

func main() {
	// useFuncAsParam(2)
	// variableLenParam()
	useDefer()
}