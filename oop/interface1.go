package main

import "fmt"

type Pet struct {}

func (p *Pet) Say(content string) {
	fmt.Println("pet say: ", content)
}

func (p *Pet) SpeakTo(name string) {
	fmt.Println("pet speak to: " + name)
}

/*
	组合相对于继承的优点在于：
	可以利用面向接口编程原则的一系列优点，封装性好，耦合性低
	相对于继承的编译期确定实现，组合的运行态指定实现，更加灵活
	组合是非侵入式的，继承是侵入式的
 */
type Dog struct {
	//使用匿名引入的方式来组合其他struct
	Pet
}

func (d *Dog) Say(content string) {
	fmt.Println("I'm a Dog")
}

func interface1() {
	d := new(Dog)
	//指定匿名struct调用内层同名方法
	d.Pet.Say("Hello")
	//默认调用外层同名方法
	d.Say("Hello")
	//调用Pet方法
	d.SpeakTo("Host")
}