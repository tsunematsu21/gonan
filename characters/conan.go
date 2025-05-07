package characters

import (
	"fmt"

	"github.com/tsunematsu21/gonan/gadgets"
)

type Conan interface {
	Aptx4869User
	Character
	Detective
	DetectiveBoys
	RealFace
	Shoot(*gadgets.StunGunWristwatch, SleepingDetective) bool
	StartChangingVoice(*gadgets.VoiceChangingBowtie, Character)
	StopChangingVoice()
}

type conan struct {
	age          int
	name         string
	displayName  string
	voiceChanger *gadgets.VoiceChangingBowtie
	isAptxized   bool
}

func (c *conan) Name() string {
	if c.isAptxized {
		return c.RealName()
	}
	return "江戸川コナン"
}

func (c *conan) Age() int {
	if c.isAptxized {
		return 17
	}
	return 7
}

func (c *conan) DisplayName() string {
	if c.isAptxized {
		return "新一"
	}
	return "コナン"
}

func (c *conan) Speek(msg string) {
	if c.voiceChanger != nil {
		c.voiceChanger.Speek(msg)
	} else {
		fmt.Printf("(%s) %s\n", c.displayName, msg)
	}
}

func (c *conan) RealName() string {
	return "工藤新一"
}

func (c *conan) Minimize() {
	c.isAptxized = true
}

func (c *conan) Maximize() {
	c.isAptxized = false
}

func (c *conan) WhoAreYou() {
	c.Speek(fmt.Sprintf("%s、探偵さ。", c.Name()))
}

func (c *conan) Shoot(sgw *gadgets.StunGunWristwatch, sd SleepingDetective) bool {
	if ok := sgw.Shoot(); ok {
		sd.sleep()
		return true
	} else {
		return false
	}
}

func (c *conan) StartChangingVoice(vcb *gadgets.VoiceChangingBowtie, ch Character) {
	vcb.SetVoice(ch.DisplayName())
	c.voiceChanger = vcb
}

func (c *conan) StopChangingVoice() {
	c.voiceChanger = nil
}

func NewConan() Conan {
	return &conan{
		name:        "江戸川コナン",
		age:         10,
		displayName: "コナン",
	}
}
