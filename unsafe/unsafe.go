package main

/*
	unsafe包通常用在和外部编程语言交互的场景中。在Go内部用到的情况比较少
	Go不支持强制类型转换，通过unsafe包可以将变量指针转化成任意类型的变量指针
 */
import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

func unsafe1() {
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))
	fmt.Println(unsafe.Pointer(&i))
	fmt.Println(f)
}

/*
	别名变量还是可以使用unsafe进行转换
 */
type CustomInt int

func unsafe2() {
	a := []int{1, 2, 3, 4}
	b := *(*[]CustomInt)(unsafe.Pointer(&a))
	fmt.Println(a)
	fmt.Println(b)
}

/*
	为了读写大数据安全，会在写数据时先将新数据写在一块新内存中
	等数据写完之后，使用unsafe包将原数据的指针指向新数据的内存地址，从而实现读写数据的安全性

	对同一个数据内存区sharedBufPtr进行读写操作，读操作和写操作同时在不同协程中进行
	写操作导致sharedBufPtr指针不停变化，但是读操作读到的sharedBufPtr数据始终是完整的，并没有出现脏数据
	这就是因为写操作时，在数据没有写完之前，是在另外一块内存进行写操作
	在数据全部写完之后，使用atomic原子操作将sharedBufPtr的地址进行修改
	这样做的好处是既保证了数据被不同协程操作的安全性，又提高了读写效率，无需在写操作时将读操作暂停
 */
func unsafe3() {
	var sharedBufPtr unsafe.Pointer
	writeFn := func() {
		var data []int
		for i := 0; i < 15; i++ {
			data = append(data, i)
		}
		//数据写完，修改指针指向
		atomic.StorePointer(&sharedBufPtr, unsafe.Pointer(&data))
	}
	readFn := func() {
		data := atomic.LoadPointer(&sharedBufPtr)
		fmt.Println(data, *(*[]int)(data))
	}
	var wg sync.WaitGroup
	writeFn()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 5; j++ {
				writeFn()
				time.Sleep(100 * time.Millisecond)
			}
			wg.Done()
		}()
	}
	/*
		写数据的同时读数据
	 */
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 20; j++ {
				readFn()
				time.Sleep(100 * time.Millisecond)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	unsafe1()
	unsafe2()
	unsafe3()
}