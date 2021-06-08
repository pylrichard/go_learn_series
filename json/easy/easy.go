package easy

import (
	"fmt"

	"go/go_learn_series/json/typedef"
)

/*
	标准库json解析是通过反射机制实现，就造成解析性能低下
	在解析本地配置文件的场景下使用较多，但在高并发场景下就显得不足

	使用easyjson包可以避免使用反射，只针对预先定义的结构体，对输入的json字符串进行字符串截取
	并将对应的json字段赋值给结构体字段
	生成解析代码时easyjson还是会使用反射机制
	所以需要为每个解析字段指定FieldTag，否则easyjson会将变量名当作解析字段生成解析代码
	解析代码生成之后，easyjson解析json数据就和反射机制没有关系
	easyjson parser and codegen based on reflection, so it won't works on package main files
	because they cant be imported by parser

	easyjson提供了代码生成工具，可以一键生成go文件中定义的结构体对应的解析代码：
	go get -u github.com/mailru/easyjson/...
	go mod tidy
	easyjson -all ./typedef.go
	为typedef.go中定义的所有struct生成对应的json解析规则
*/
func Json() {
	t := new(typedef.Teacher)
	err := t.UnmarshalJSON([]byte(typedef.JsonStr))
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("teacher: ", t)
	}
	jsonStr, err := t.MarshalJSON()
	if err != nil {
		fmt.Println("err: ®", err)
	} else {
		fmt.Println("json string: ", string(jsonStr))
	}
}