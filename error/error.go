package main

import (
	"errors"
	"fmt"
)

func deferRecover() {
	//recover()返回了panic()抛出的Error
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	defer deferRecover()

	//os.Exit(1)不会执行defer
	panic(errors.New("panic"))
}