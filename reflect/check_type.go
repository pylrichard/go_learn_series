package main

import (
	"fmt"
	"reflect"
)

func CheckType() {
	a := 12
	checkType(a)

	b := "hello go"
	checkType(b)

	var c float64 = 66
	checkType(c)
}

func checkType(v interface{}) {
	//返回变量类型 (reflect.Type枚举类型)
	t := reflect.TypeOf(v)
	//判断变量类型
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	default:
		fmt.Println("unknown type: ", t)
	}
}