package gonan

import (
	"fmt"
	"sync"
)

type Kogoro interface {
	Character
	Sleepable
}

type kogoro struct {
	sleeped bool
}

func (c *kogoro) Name() string {
	return "毛利小五郎"
}

func (c *kogoro) DisplayName() string {
	return "小五郎"
}

func (c *kogoro) sleep() error {
	if c.sleeped {
		return fmt.Errorf("%s is already sleeped", c.Name())
	} else {
		fmt.Printf("(%s) ふにゃ...\n", c.DisplayName())
		c.sleeped = true
		return nil
	}
}

func (c *kogoro) WakeUp() {
	c.sleeped = false
}

var (
	singletonKogoro *kogoro
	onceKogoro      sync.Once
)

func GetKogoro() Kogoro {
	onceKogoro.Do(func() {
		singletonKogoro = &kogoro{}
	})
	return singletonKogoro
}
