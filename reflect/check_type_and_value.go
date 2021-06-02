package main

import (
	"fmt"
	"reflect"
)

func CheckTypeAndValue() {
	a := 12
	checkTypeAndValue(a)
}

func checkTypeAndValue(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Println("type:", t)
	//返回变量的值 (reflect.Value)
	value := reflect.ValueOf(v)
	fmt.Println("value:", value)
	t = value.Type()
	fmt.Println("type:", t)
}