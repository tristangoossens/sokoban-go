package trisoban

import (
	tl "github.com/JoelOtter/termloop"
	"io/ioutil"
)

func StartGame() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	player := NewPlayer()

	// leveltje := new(Level)
	// leveltje.Entity = tl.NewEntity(1, 1, 1, 1)

	logofile, _ := ioutil.ReadFile("util/logo.text")
	logoEntity := tl.NewEntityFromCanvas(10, 3, tl.CanvasFromString(string(logofile)))
	level.AddEntity(logoEntity)

	level.AddEntity(player)
	game.Screen().SetLevel(level)
	game.Screen().SetFps(40)
	game.Start()
}
