package gonan

import (
	"fmt"
	"sync"
)

type Conan interface {
	Character
	Detective
	VoiceChangable
}

type conan struct {
	changeVoiceTarget Character
}

func (c *conan) Name() string {
	return "江戸川コナン"
}

func (c *conan) DisplayName() string {
	if c.changeVoiceTarget != nil {
		return fmt.Sprintf("変声機:%sの声", c.changeVoiceTarget.DisplayName())
	} else {
		return "コナン"
	}
}

func (c *conan) changeVoice(target Character) {
	c.changeVoiceTarget = target
}

func (c *conan) restoreVoice() {
	c.changeVoiceTarget = nil
}

func (c *conan) solve() bool {
	return true
}

var (
	singletonConan *conan
	onceConan      sync.Once
)

func GetConan() Conan {
	onceConan.Do(func() {
		singletonConan = &conan{}
	})
	return singletonConan
}
