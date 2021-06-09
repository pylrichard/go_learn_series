package main

import (
	"fmt"
	"net/http"
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