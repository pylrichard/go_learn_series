package main

/*
	组合和接口结合使用
 */
import "fmt"

type Animal interface {
	getName()
}

type Tiger struct {}

func (t *Tiger) getName() {
	fmt.Println("Tiger")
}

type Parrot struct {}

func (p *Parrot) getName() {
	fmt.Println("Parrot")
}

type Pigeon struct {}

func (p *Pigeon) getName() {
	fmt.Println("Pigeon")
}

type CatFamily struct {
	Animal
}

func (cf *CatFamily) getName() {
	fmt.Println("CatFamily")
}

type Bird struct {
	Animal
}

func interface2() {
		//Tiger实现Animal接口
		t := Tiger{}
		t.getName()
		//CatFamily嵌入Animal接口
		cf := CatFamily{&Tiger{}}
		cf.getName()
		cf.Animal.getName()
		//运行时动态指定实现
		b := Bird{&Parrot{}}
		b.getName()
		b.Animal = &Pigeon{}
		b.getName()
}