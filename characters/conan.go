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
	StartChangeVoice(*gadgets.VoiceChangingBowtie, Character)
	StopChangingVoice()
}

type conan struct {
	age          int
	name         string
	displayName  string
	voiceChanger *gadgets.VoiceChangingBowtie
}

func (c *conan) Name() string {
	return c.name
}

func (c *conan) Age() int {
	return c.age
}

func (c *conan) DisplayName() string {
	return c.displayName
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
	c.age = 7
	c.name = "江戸川コナン"
	c.displayName = "コナン"
}

func (c *conan) Maximize() {
	c.age = 17
	c.name = "工藤新一"
	c.displayName = "新一"
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

func (c *conan) StartChangeVoice(vcb *gadgets.VoiceChangingBowtie, ch Character) {
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
