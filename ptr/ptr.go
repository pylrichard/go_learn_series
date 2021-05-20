package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(p, *p)
	*p = 6
	fmt.Println(p, *p)
}