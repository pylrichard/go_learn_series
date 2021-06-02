package main

import (
	"fmt"
	"reflect"
)

/*
	切片、Map、结构体是不能直接比较的，只能和nil进行比较
	reflect.DeepEqual()可以用来比较两个Map/切片/结构体是否相等
 */
func DeepEqual() {
	m1 := map[int]string{1: "one", 2:"two"}
	m2 := map[int]string{1: "three", 2:"two"}
	fmt.Println("m1 == m2 ?: ", reflect.DeepEqual(m1, m2))
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2 ?: ", reflect.DeepEqual(s1, s2))
}