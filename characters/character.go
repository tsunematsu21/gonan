package characters

type Character interface {
	Name() string
	Age() int
	DisplayName() string
	Speek(string)
}

type DetectiveBoys interface{}

type Detective interface {
	WhoAreYou()
}

type SleepingDetective interface {
	sleep()
	WakeUp()
	IsSleep() bool
}

type Aptx4869User interface {
	Minimize()
	Maximize()
}

type BlackOrganization interface {
	CodeName() string
}

type RealFace interface {
	RealName() string
}
