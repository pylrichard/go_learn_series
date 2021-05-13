package main

import "fmt"

func defineMap() {
	m1 := map[string]int{"one": 1, "two": 2}
	fmt.Println(m1["one"])
	m2 := map[string]int{}
	m2["one"] = 1
	m3 := make(map[string]int, 10)
	m3["two"] = 2
	
	for k, v := range m1 {
		fmt.Printf("%s:%d\n", k, v)
	}
	
	m4 := map[int]func(opt int) int{}
	m4[1] = func(opt int) int { return opt }
	m4[2] = func(opt int) int { return opt*opt }
	fmt.Println(m4[1](2))
	fmt.Println(m4[2](2))
}

func judgeMapZeroValue() {
	m := map[string]int{}

	if value, ok := m["one"]; ok {
		fmt.Printf("key one exist value: %d\n", value)
	} else {
		fmt.Println("key one does not exist")
	}
}

func implMapSet() {
	set := map[string]bool {}
	cd := "cheng du"
	if set[cd] {
		fmt.Printf("%s exist", cd)
	} else {
		fmt.Printf("%s does not exist", cd)
	}

	for item, _ := range set {
		fmt.Println("set item:", item)
	}

	fmt.Println(len(set))

	delete(set, cd)
}

func main() {
	defineMap()
	judgeMapZeroValue()
}