package main

import "fmt"

func main() {
	//不用声明长度
	var s1 []int
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, 1)
	fmt.Println(len(s1), cap(s1))
	s2 := []int{}
	s3 := []int{ 1, 2, 3 }
	fmt.Println(s2, s3)
	s4 := make([]int, 3, 8)
	fmt.Println(s4)
	
	/*
		q2和summer共享内存，它们的ptr指向同一个底层数组
		变
	 */
	months := []string{ 
		"Jan", "Feb", "Mar", 
		"Apr", "May", "Jun", 
		"Jul", "Aug", "Sep", 
		"Oct", "Nov", "Dec",
	}
	//ptr指向"Apr"
	q2 := months[3:6]
	fmt.Println(q2)
	fmt.Println(len(q2), cap(q2))
	//ptr指向"Jun"
	summer := months[5:8]
	fmt.Println(summer)
	fmt.Println(len(summer), cap(summer))
	//修改summer的元素，那么q2[2]，months[5]也发生了改变
	summer[0] = "UnKnown"
	fmt.Println(summer)
	fmt.Println(q2)
	fmt.Println(months)
}