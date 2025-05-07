package characters

import "fmt"

type Kogoro interface {
	Character
	SleepingDetective
}

type kogoro struct {
	isSleeping bool
}

func (k *kogoro) Name() string {
	return "毛利小五郎"
}

func (k *kogoro) Age() int {
	return 38
}

func (k *kogoro) DisplayName() string {
	return "小五郎"
}

func (k *kogoro) Speek(msg string) {
	fmt.Printf("(%s) %s\n", "小五郎", msg)
}

func (k *kogoro) sleep() {
	k.Speek("ふにゃ・・・。")
	k.isSleeping = true
}

func (k *kogoro) IsSleep() bool {
	return k.isSleeping
}

func (k *kogoro) WakeUp() {
	k.isSleeping = false
}

func NewKogoro() Kogoro {
	return &kogoro{}
}
