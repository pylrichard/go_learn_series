package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name	string
	Age		int		`format:"student_age"`
	Grade	int
}

func (s *Student) UpdateGrade(grade int) {
	s.Grade = grade
}

func CallMethod() {
	s := Student{ Name: "pyl", Age: 33, Grade: 1 }
	//获得reflect.Value类型，属性值
	name := reflect.ValueOf(s).FieldByName("Name")
	fmt.Println("name:", name)
	//获得reflect.StructField类型，属性特征，包含属性数据类型，Tag信息等
	field, ok := reflect.TypeOf(s).FieldByName("Age")
	if ok {
		//reflect.StructField.Tag在Go内置的json解析中，被用来解析数据结果和Struct属性名之间的映射关系
		fmt.Println("type of Age:", field.Tag.Get("format"))
	} else {
		fmt.Println("failed")
	}
	fmt.Println("before:", s)
	//注意传入&s
	reflect.ValueOf(&s).MethodByName("UpdateGrade").Call([]reflect.Value{reflect.ValueOf(2)})
	fmt.Println("after:", s)
}