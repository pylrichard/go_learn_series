package main

import "fmt"

func main() {
	//声明后各元素初始值为0
	var arr1 [3]int
	arr1[0] = 10
	fmt.Println(arr1)
	//初始化并赋值
	arr2 := [3]int{ 1, 2, 3 }
	fmt.Println(arr2)
	//初始化不指定数组长度
	arr3 := [...]int { 1, 2, 3 }
	fmt.Println(arr3)
	
	/*
		数组遍历
	 */
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%d:%d ", i, arr1[i])
	}
	fmt.Println()
	for index, item := range arr1 {
		fmt.Printf("%d:%d ", index, item)
	}
	fmt.Println()
	/*
		slice := array[startIndex:endIndex]
		包含startIndex，不包含endIndex
		startIndex和endIndex，不支持负数
	 */
	s1 := arr1[0:]
	fmt.Println(s1)
	//从0到1，不包括2
	s2 := arr1[:2]
	fmt.Println(s2)
	//从1(包括1)开始
	s3 := arr1[1:]
	fmt.Println(s3)
	//整个数组
	s4 := arr1[:]
	fmt.Println(s4)
}