package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	使用不同的协程对counter进行加操作
	运行会发现打印的counter小于5000，这是因为不同协程之间具有竞争关系，counter是协程不安全的
*/
func lock1() {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(counter)
}

func lock2() {
	/*
		要想保证数据在不同的协程被安全操作，需要使用Go锁机制
	 */
	counter := 0
	var m sync.Mutex
	for i := 0; i < 5000; i++ {
		go func() {
			//协程运行完成后释放锁
			defer m.Unlock()
			//使用sync包的Mutex对数据操作进行了加锁
			m.Lock()
			counter++
		}()
	}
	/*
		加一个Sleep使当前协程等待一定时长
		这是因为counter变量也被主协程持有，主协程可能会更早运行完毕
		导致counter在没有完成5000次操作前，就被主协程打印输出
	 */
	time.Sleep(time.Second)
	fmt.Println(counter)
}

func lock3() {
	counter := 0
	/*
		RWMutex将数据读取/写入的锁分开，多个协程读取同一个数据不是互斥的，可以同时进行数据读取
		只有当多个协程写同一个数据，才是互斥的，才会启动锁机制，所以RWMutex比Mutex的性能更好
	 */
	var m sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		//WaitGroup添加一个协程任务
		wg.Add(1)
		go func() {
			defer func() {
				m.Unlock()
				//协程运行完毕，通知WaitGroup
				wg.Done()
			}()
			m.Lock()
			counter++
		}()
	}
	//WaitGroup等待所有协程运行完毕
	wg.Wait()
	fmt.Println(counter)
}