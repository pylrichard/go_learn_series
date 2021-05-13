package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func rangeForStr() {
	s := "Hello Go语言"
	for _, c := range s {
		fmt.Println(c)
	}
}

func getStrCharLen() {
	s := "Hello Go语言"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	runeStr := []rune(s)
	fmt.Println(len(runeStr))
}

func procStr() {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	fmt.Println(parts)
	result := strings.Join(parts, "-")
	fmt.Println(result)

	fmt.Println(strconv.Itoa(66))
	if i, err := strconv.Atoi("88"); err == nil {
		fmt.Println(i)	
	} else {
		fmt.Println("convert error")
	}
}

func main() {
	// rangeForStr()
	// getStrCharLen()
	procStr()
}