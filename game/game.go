package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

func StartGame() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	leveltje := new(Level)
	leveltje.Entity = tl.NewEntity(1, 1, 1, 1)

	level.AddEntity(leveltje)
	game.Screen().SetLevel(level)
	game.Start()
}