package collector

import (
	"context"
	"errors"
	"time"

	"go/go_learn_series/micro_kernel/agent"
)

type Collector struct {
	receiver	agent.EventReceiver
	ctx			context.Context
	stopChan 	chan struct{}
	name		string
	content		string
}

func New(name string, content string) *Collector {
	return &Collector{
		stopChan: 	make(chan struct{}),
		name:		name,
		content:	content,
	}
}

func (c *Collector) Init(receiver agent.EventReceiver) error {
	c.receiver = receiver

	return nil
}

func (c *Collector) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			c.stopChan <- struct{}{}
			break
		default:
			time.Sleep(60 * time.Millisecond)
			c.receiver.OnEvent(agent.Event(c.name + ":" + c.content))
		}
	}
}

func (c *Collector) Stop() error {
	select {
	case <-c.stopChan:
		return nil
	case <-time.After(1 * time.Second):
		return errors.New("timeout")
	}
}

func (c *Collector) Destroy() error {
	return nil
}