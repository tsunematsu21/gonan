package gonan

type Character interface {
	Name() string
	DisplayName() string
}

type Sleepable interface {
	sleep() error
	WakeUp()
}

type VoiceChangable interface {
	changeVoice(target Character)
	restoreVoice()
}

type AptxCapable interface {
	Aptxize()
	Deaptxize()
}

type mob struct {
	name        string
	displayName string
}

func (c *mob) Name() string {
	return c.name
}

func (c *mob) DisplayName() string {
	return c.displayName
}

func NewMob(name string, displayName string) Character {
	return &mob{
		name:        name,
		displayName: displayName,
	}
}
