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

func (p *Pythonista) Write(name string) {
	fmt.Println("Hello Pythonista: ", name)
}

func Write(name string, p Programer) {
	p.Write(name)
}

/*
	小功能接口定义
 */
type Reader interface {
	Read(buf []byte) (int, error)
}

type Writer interface {
	Write(buf []byte) (int, error)
}

//大接口使用小接口嵌套组装
type IO interface {
	Reader
	Writer
}

type Pet struct {}

func (p *Pet) Say(content string) {
	fmt.Println("pet say: ", content)
}

func (p *Pet) SpeakTo(name string) {
	fmt.Println("pet speak to: " + name)
}

//使用组合
type Dog struct {
	Pet
}

func (d *Dog) Say(content string) {
	//指定匿名struct调用内层同名方法
	d.Pet.Say("Hello")
	fmt.Println("I'm a Dog")
}

func main() {
	name := "pyl"
	//gp变量被声明之后，它有两部分：类型(Gopher)和数据(&Gopher{})
	g := new(Gopher)
	Write(name, g)
	p := new(Pythonista)
	Write(name, p)

	d := new(Dog)
	//默认调用外层同名方法
	d.Say("Hello")
	//调用Pet方法
	d.SpeakTo("Host")
}