package agent

import (
	"context"
	"errors"
	"fmt"
)

/*
	微内核模式(micro-kernel)在IDE中非常常见，比如通过安装插件支持不同语言。微内核的特点是易于扩展、错误隔离等
	在微内核模式中，Core System会运行多个Plugin，每个Plugin的运行相对独立，一个Plugin出错，不会影响其他Plugin的运行
 */
import (
	"sync"
)

type Event string

var (
	Running = 0
	Waiting = 1
)

var WrongStateError = errors.New("error agent state")

type EventReceiver interface {
	OnEvent(event Event)
}

//Collector 相当于Plugin
type Collector interface {
	Init(receiver EventReceiver) error
	Start(ctx context.Context) error
	Stop() error
	Destroy() error
}

//Agent 相当于Core System
type Agent struct {
	collectors	map[string]Collector
	eventBuf	chan Event
	cancel		context.CancelFunc
	ctx			context.Context
	state		int
}

func New(size int) *Agent {
	agent := Agent{
		collectors: make(map[string]Collector),
		eventBuf:	make(chan Event, size),
		state:		Waiting,
	}

	return &agent
}

//OnEvent 提供给Collector向Agent传递消息
func (a *Agent) OnEvent(event Event) {
	a.eventBuf <- event
}

func (a *Agent) Start() error {
	if a.state != Waiting {
		return WrongStateError
	}
	a.state = Running
	a.ctx, a.cancel = context.WithCancel(context.Background())
	go a.EventProcessRoutine()

	return a.startCollectors()
}

func (a *Agent) Stop() error {
	if a.state != Running {
		return WrongStateError
	}
	a.state = Waiting
	a.cancel()

	return a.stopCollectors()
}

func (a *Agent) startCollectors() error {
	var err error
	var errs []error
	var mutex sync.Mutex
	for name, collector := range a.collectors {
		go func(name string, collector Collector, ctx context.Context) {
			defer func() {
				mutex.Unlock()
			}()
			err = collector.Start(ctx)
			mutex.Lock()
			if err != nil {
				errs = append(errs, errors.New(name + ":" + err.Error()))
			}
		}(name, collector, a.ctx)
	}
	if errs == nil {
		return nil
	}

	return errs[0]
}

func (a *Agent) stopCollectors() error {
	var err error
	var errs []error
	for name, collector := range a.collectors {
		if err = collector.Stop(); err != nil {
			errs = append(errs, errors.New(name + ":" + err.Error()))
		}
	}
	if errs == nil {
		return nil
	}

	return errs[0]
}

func (a *Agent) destroyCollectors() error {
	var err error
	var errs []error
	for name, collector := range a.collectors {
		if err = collector.Destroy(); err != nil {
			errs = append(errs, errors.New(name + ":" + err.Error()))
		}
	}
	if errs == nil {
		return nil
	}

	return errs[0]
}

/*
	将Collector注册到Agent中
	并通过调用Collector的Init()，将EventReceiver传递给Collector，以便于Collector向Agent上报信息
 */
func (a *Agent) RegisterCollector(name string, collector Collector) error {
	if a .state != Waiting {
		return WrongStateError
	}
	a.collectors[name] = collector

	return collector.Init(a)
}

//EventProcessRoutine 获取从Collector上报的数据，每8条进行打印
func (a *Agent) EventProcessRoutine() {
	var count = 8
	var eventBuf = make([]Event, count)
	for {
		for i := 0; i < count; i++ {
			select {
			case eventBuf[i] = <-a.eventBuf:
			case <-a.ctx.Done():
				return
			}
		}
		fmt.Println(eventBuf)
	}
}