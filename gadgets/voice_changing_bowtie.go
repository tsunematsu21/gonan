package gadgets

import "fmt"

type VoiceChangingBowtie struct {
	voice string
}

func (v *VoiceChangingBowtie) SetVoice(name string) {
	v.voice = name
}

func (v *VoiceChangingBowtie) Speek(msg string) {
	fmt.Printf("(変声機:%sの声) %s\n", v.voice, msg)
}

func NewVoiceChangingBowtie() *VoiceChangingBowtie {
	return &VoiceChangingBowtie{}
}
