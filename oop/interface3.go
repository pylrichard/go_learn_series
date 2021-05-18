package main

import "fmt"

type Programer interface {
	Write(name string)
}

//接口有非入侵性，实现不依赖接口定义
type Gopher struct {}

func (gp *Gopher) Write(name string) {
	fmt.Println("Hello Gopher: ", name)
}

type Pythonista struct {}

func (py *Pythonista) Write(name string) {
	fmt.Println("Hello Pythonista: ", name)
}

func Write(name string, p Programer) {
	p.Write(name)
}

func interface3() {
	name := "pyl"
	//gp变量被声明之后，它有两部分：类型(Gopher)和数据(&Gopher{})
	gp := new(Gopher)
	//多态
	Write(name, gp)
	py := new(Pythonista)
	Write(name, py)
}