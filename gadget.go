package gonan

import (
	"fmt"
)

type Gadget interface {
	Use(target Character) error
}

type stunGunWristWatch struct {
	owner    Character
	magazine int
}

func (s *stunGunWristWatch) Use(target Character) error {
	if s.magazine < 1 {
		return fmt.Errorf("magazine is empty")
	}

	PlaySound("腕時計型麻酔銃の音", "パシュッ")

	s.magazine -= 1

	t, ok := target.(Sleepable)
	if ok {
		if err := t.sleep(); err != nil {
			return fmt.Errorf("failed to sleep %s: %v", target.Name(), err)
		}
	} else {
		return fmt.Errorf("target %s is not sleepable", target.Name())
	}

	return nil
}

func (s *stunGunWristWatch) Reload() error {
	if s.magazine >= 1 {
		return fmt.Errorf("magazine is full")
	}

	s.magazine += 1

	return nil
}

func NewStunGunWristWatch(user Character) Gadget {
	return &stunGunWristWatch{
		owner:    user,
		magazine: 1,
	}
}

type voiceChangingBowtie struct {
	owner VoiceChangable
}

func (v *voiceChangingBowtie) Use(target Character) error {
	if target == nil {
		v.owner.restoreVoice()
	} else {
		v.owner.changeVoice(target)
	}

	return nil
}

func NewVoiceChangingBowtie(user VoiceChangable) Gadget {
	return &voiceChangingBowtie{
		owner: user,
	}
}
