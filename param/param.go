package main

import "fmt"

//值传递将参数进行拷贝，函数内操作的是副本
func modify1(i int) {
	i = i + 1
	fmt.Println("in modify1 ", i)
}

//引用传递
func modify2(p *int, tag string) {
	fmt.Println(p, tag)
	*p = *p + 1
}

func main() {
	i := 1
	//p是i的内存地址
	p := &i
	//&p是p的内存地址
	fmt.Println(p, &p)
	modify1(i)
	fmt.Println("after modify1 ", i)
	fmt.Println("before modify2 ", i)
	modify2(&i, "&i")
	fmt.Println("after modify2 ", i)
	modify2(p, "p")
}