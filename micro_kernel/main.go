package main

import (
	"go/go_learn_series/micro_kernel/agent"
	"go/go_learn_series/micro_kernel/collector"
	"time"
)

func main() {
	a := agent.New(100)
	c1 := collector.New("c1", "1")
	c2 := collector.New("c2", "2")
	_ = a.RegisterCollector("c1", c1)
	_ = a.RegisterCollector("c2", c2)
	_ = a.Start()
	time.Sleep(1 * time.Second)
	_ = a.Stop()
}