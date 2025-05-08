package gonan

import (
	"fmt"
	"sync"
)

type Conan interface {
	Character
	Detective
	VoiceChangable
	AptxCapable
}

type conan struct {
	changeVoiceTarget Character
	aptxized          bool
}

func (c *conan) Name() string {
	if c.aptxized {
		return "江戸川コナン"
	} else {
		return "工藤新一"
	}
}

func (c *conan) DisplayName() string {
	if c.changeVoiceTarget != nil {
		return fmt.Sprintf("変声機:%sの声", c.changeVoiceTarget.DisplayName())
	} else if c.aptxized {
		return "コナン"
	} else {
		return "新一"
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

func (c *conan) setAptxized(b bool) {
	c.aptxized = b
}

func (c *conan) Aptxize() {
	c.aptxized = true
}

func (c *conan) Deaptxize() {
	c.aptxized = false
}

var (
	singletonConan *conan
	onceConan      sync.Once
)

func GetConan() Conan {
	onceConan.Do(func() {
		singletonConan = &conan{
			aptxized: true,
		}
	})
	singletonConan.aptxized = true
	return singletonConan
}

func GetShinichi() Conan {
	c := GetConan()
	c.Deaptxize()
	return c
}
