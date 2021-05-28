package main

import (
	"go/go_learn_series/pool/obj"
	"go/go_learn_series/pool/sync"
)

func main() {
	obj.TestObjPool()
	sync.TestSyncPool()
}