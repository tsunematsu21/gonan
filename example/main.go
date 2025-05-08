package main

import (
	"fmt"

	"github.com/tsunematsu21/gonan"
)

func main() {
	conan := gonan.GetConan()
	kogoro := gonan.GetKogoro()

	watch := gonan.NewStunGunWristWatch(conan)

	gonan.Say(conan, "この腕時計型麻酔銃でおっちゃんを眠らせて...")
	if err := watch.Use(kogoro); err != nil {
		fmt.Printf("failed to use stun gun wrist watch: %v\n", err)
	}

	bowtie := gonan.NewVoiceChangingBowtie(conan)
	bowtie.Use(kogoro)

	culprit := gonan.NewMob("山田", "一郎")
	c, err := gonan.NewCase(
		gonan.WithVictim(gonan.NewMob("山田次郎", "次郎")),
		gonan.WithLocation("寝室"),
		gonan.AddSuspect(culprit, true),
		gonan.AddSuspect(gonan.NewMob("山田三郎", "三郎"), false),
	)
	if err != nil {
		fmt.Println("failed to open case:", err)
	}

	t, err := gonan.NewTruth(culprit, "恨んでいたから")
	if err != nil {
		fmt.Println("failed to create truth:", err)
	}

	if err := c.Close(conan, t); err != nil {
		fmt.Println("failed to close case:", err)
	}

	bowtie.Use(nil)
	gonan.Say(conan, "真実はいつも一つ！")

	conan.Deaptxize()
	gonan.Say(conan, "真実はいつも一つ！")

	conan.Aptxize()
	gonan.Say(conan, "真実はいつも一つ！")
}
