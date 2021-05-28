package sync

import (
	"fmt"
	"runtime"
	"sync"
)

func TestSyncPool() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("New")
			return 6
		},
	}
	//首次从sync.Pool获取对象通过New()创建
	v1 := pool.Get().(int)
	fmt.Println(v1)
	pool.Put(8)
	//执行GC，会将sync.Pool中的缓存对象清空，再次获取时仍然通过New()重新创建
	//不执行GC，再次获取到的是放回的缓存对象数据
	runtime.GC()
	v2 := pool.Get().(int)
	fmt.Println(v2)

	for i := 0; i < 3; i++ {
		pool.Put(8)
	}
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get().(int))
			wg.Done()
		}(i)
	}
	wg.Wait()
}