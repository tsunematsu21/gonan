package gadgets

import "fmt"

type StunGunWristwatch struct {
	bullets int
}

func (s *StunGunWristwatch) Shoot() bool {
	if s.bullets <= 0 {
		return false
	}
	fmt.Println("(麻酔銃の音) プシュッ")
	s.bullets -= 1
	return true
}

func NewStunGunWristwatch() *StunGunWristwatch {
	return &StunGunWristwatch{
		bullets: 1,
	}
}

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
