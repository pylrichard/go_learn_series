package once

import (
	"fmt"
	"sync"
	"unsafe"
)

type Singleton struct {}

var once sync.Once
var instance *Singleton

func getSingletonObj() *Singleton {
	//多线程环境下只执行一次，Singleton对象只会创建一次
	once.Do(func() {
		fmt.Println("create instance")
		instance = new(Singleton)
	})

	return instance
}

func SyncOnce() {
	var wg sync.WaitGroup
	/*
		多个协程调用getSingletonObj()获取的是同一个对象
	 */
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := getSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}