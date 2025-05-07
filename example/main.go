package main

import (
	"github.com/tsunematsu21/gonan/characters"
	"github.com/tsunematsu21/gonan/gadgets"
)

func main() {
	conan := characters.NewConan()

	conan.Speek("真実はいつも一つ！")

	conan.WhoAreYou() // (コナン) 江戸川コナン、探偵さ。
	conan.Maximize()

	conan.WhoAreYou() // (新一) 工藤新一、探偵さ。
	conan.Minimize()

	kogoro := characters.NewKogoro()
	watch := gadgets.NewStunGunWristwatch()
	bowtie := gadgets.NewVoiceChangingBowtie()

	conan.Speek("この腕時計型麻酔銃でおっちゃんを眠らせて...")
	if ok := conan.Shoot(watch, kogoro); ok { // (麻酔銃の音) プシュッ
		conan.StartChangeVoice(bowtie, kogoro)
		conan.Speek("犯人はお前だ！")
		conan.StopChangingVoice()
	}

	conan.Speek("いつものようにおっちゃんを眠らせて...")
	if ok := conan.Shoot(watch, kogoro); !ok {
		conan.Speek("しまった！弾切れだ！")
	}
}
