# gonan
A　Golang implementation of "Detective Conan".

```go
conan := gonan.NewConnan()
kogoro := gonan.NewKogoro()

watch := gonan.NewStunGunWristWatch(conan)
watch.Use(kogoro)

bowtie := gonan.NewVoiceChangingBowtie(conan)
bowtie.Use(kogoro)

culprit := gonan.NewCharacter("佐藤")
c := gonan.NewCase(
  gonan.WithVictim("山田"),
  gonan.WithCauseOfDeath("刺殺"),
  gonan.WithLocation("自宅"),
  gonan.AddCulprit(culprit),
  gonan.AddSuspect(gonan.NewCharacter("鈴木")),
  gonan.AddSuspect(gonan.NewCharacter("田中")),
)

t := gonan.NewTruth(
  gonan.WithCulprit()
)

if ok := c.Close(conan, t); ok {
  kogoro.WakeUp()
}

bowtie.Use(nil)

shinichi := gonan.NewShinichi()
gonan.Aptxize(shinichi)
shinichi.WhoAreYou() // 江戸川コナン、探偵さ。

gonan.Deaptxize(conan)
conan.WhoAreYou() // 工藤新一、探偵さ。

// ToDo
if shinichi == conan {
  fmt.Println("同一人物")
}

amuro := gonan.NewCharacter[gonan.Amuro]()
furuya, _ := amuro.(gonan.Furuya)
shinichi, _ = conan.(gonan.Shinichi)


```