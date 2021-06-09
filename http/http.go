package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	/*
		访问http://127.0.0.1:8080/time，返回当前时间戳
		访问http://127.0.0.1:8080/time/xxx，由于/time是一个叶子路由，因此并不会路由到/time，而是路由到/ Handler
		Go的http包路由匹配采用最长匹配原则，如果有多个匹配，会使用匹配路径最长的进行处理
		如果没有找到任何匹配项，会返回404错误

		URL分为两种：
		末尾包含/：表示一个子树，后面可以跟其他子路径
		末尾不含/：表示一个叶子路径，固定的路由路径

		开发过程中对代码进行性能检测，而对于线上代码不可能一直输出测试文件，这样会影响到线上业务
		不过pprof同样提供了线上代码性能测试的方法，通过http请求对线上业务进行一定时长的采样

		使用pprof对线上业务进行性能采样的步骤：
		1. 导入"net/http/pprof"包，并启动http server
		2. 通过http://<host>:<port>/debug/pprof访问
		3. 进行采样获取cpu profile，seconds参数指定采样时长
			go tool pprof http://<host>:<port>/debug/pprof/profile?seconds=10
		4. 生成火焰图
			go tool pprof -http=:8081 ~/pprof/pprof.samples.cpu.001.pb.gz
	 */
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "hello http")
	})
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		str := fmt.Sprintf("{\"time\": %s}", t)
		_, _ = w.Write([]byte(str))
	})
	_ = http.ListenAndServe(":8080", nil)
}