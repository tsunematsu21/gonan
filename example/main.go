package main

import (
	"fmt"

	"github.com/tsunematsu21/gonan"
)

func main() {
	conan := gonan.GetConan()
	kogoro := gonan.GetKogoro()

	watch := gonan.NewStunGunWristWatch(conan)
	if err := watch.Use(kogoro); err != nil {
		fmt.Printf("failed to use stun gun wrist watch: %v\n", err)
	}
	// (小五郎) ふにゃ...

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
	// (変声機:小五郎の声) やっと分かったんですよ...
	// (変声機:小五郎の声) 寝室で次郎さんを殺害した犯人がね...
	// (変声機:小五郎の声) 犯人は一郎さん、あなただ！
	// (変声機:小五郎の声) 犯行の動機はおそらく恨んでいたからでしょう

	bowtie.Use(nil)
	gonan.Speek(conan, "真実はいつも一つ！")
	// (コナン) 真実はいつも一つ！
}
