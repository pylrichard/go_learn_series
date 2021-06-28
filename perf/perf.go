package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

const row = 1000
const col = 2000

//对二维数组进行随机填充
func fillMatrix(m *[row][col]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[i][j] = s.Intn(10000)
		}
	}
}

//对二维数组的每一行进行加和操作
func calculate(m *[row][col]int) {
	for i := 0; i < row; i++ {
		sum := 0
		for j := 0; j < col; j++ {
			sum += m[i][j]
		}
	}
}

var cpuProfile = flag.String("cpu_profile", "", "write cpu profile to `file`")
var memProfile = flag.String("mem_profile", "", "write mem profile to `file`")
var goroutineProfile = flag.String("goroutine_profile", "", "write goroutine profile to `file`")

/*
	go build perf.go
	./perf -cpu_profile cpu.prof -mem_profile mem.prof -goroutine_profile goroutine.prof
	格式：go tool pprof <二进制文件名> <要查看的profile文件>
	go tool pprof perf cpu.prof
	top 查看cpu使用情况
	list funcName 查看指定函数耗时情况
	svg 将所有函数运行情况以svg格式输出
 */
func main() {
	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			fmt.Println("err create profile:", err)
		}
		_ = f.Close()
		//CPUProfile代码需要放在待测试代码的前面，开启pprof测试
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Println("err start cpu profile:", err)
		}
		defer pprof.StopCPUProfile()
	}

	//内存用量二维数组占用较多
	arr := [row][col]int{}
	fillMatrix(&arr)
	calculate(&arr)
	//查看二维数组变量在GC之后是否可以被成功回收
	//可以在输出mem.prof内容之前，手动触发一次GC
	runtime.GC()

	/*
		将pprof信息中的指标导出到profile文件，所以要放在待测试代码的后面执行
	 */
	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			fmt.Println("err create mem profile:", err)
		}
		_ = f.Close()
		if err = pprof.WriteHeapProfile(f); err != nil {
			fmt.Println("err write mem profile:", err)
		}
	}

	if *goroutineProfile != "" {
		f, err := os.Create(*goroutineProfile)
		if err != nil {
			fmt.Println("err create goroutine profile:", err)
		}
		//Lookup()传入不同性能flag，导出不同性能指标信息
		if goProf := pprof.Lookup("goroutine"); goProf == nil {
			fmt.Println("err write goroutine")
		} else {
			_ = goProf.WriteTo(f, 0)
		}
		_ = f.Close()
	}
}