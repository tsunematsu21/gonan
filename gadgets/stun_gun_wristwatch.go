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
