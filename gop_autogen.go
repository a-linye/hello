package main

import "github.com/goplus/spx"

type Calf struct {
	spx.Sprite
	*Game
}
type Game struct {
	spx.Game
	Calf Calf
}
//line tutorial/00-Hello/main.spx:5
func (this *Game) MainEntry() {
//line tutorial/00-Hello/main.spx:5:1
	spx.Gopt_Game_Run(this, "assets", &spx.Config{Title: "Hello (by Go+ spx engine)"})
}
func main() {
	spx.Gopt_Game_Main(new(Game))
}
//line tutorial/00-Hello/Calf.spx:1
func (this *Calf) Main() {
//line tutorial/00-Hello/Calf.spx:1:1
	this.OnStart(func() {
//line tutorial/00-Hello/Calf.spx:2:1
		this.Say("Hello Go+")
	})
}
