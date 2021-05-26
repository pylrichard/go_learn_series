package main

import (
	"go/go_learn_series/concurrency/once"
	"go/go_learn_series/concurrency/resp"
)

func main() {
	once.SyncOnce()
	resp.FirstResp()
	resp.AllResp()
}