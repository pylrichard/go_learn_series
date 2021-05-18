package main

import "fmt"

//Go不支持继承
type Person struct {
	Name string
	Age  int
}

//对象成员进行值复制
func (p Person) toString() string {
	return "Name: " + p.Name
}

//接收者为指针，避免对象成员进行内存拷贝
func (p *Person) String() string {
	return "Name: " + p.Name
}

func testStruct() {
	p1 := Person{
		Name:	"pyl",
		Age:	33,
	}
	fmt.Println(p1.toString())
	//返回对象指针，相当于p2 := &Person{}
	p2 := new(Person)
	p2.Name = "zy"
	fmt.Println(p2.String())
}