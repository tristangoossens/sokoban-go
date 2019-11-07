package trisoban

import (
	tl "github.com/JoelOtter/termloop"
	"io/ioutil"
)

func NewGameScreen() *Gamescreen {
	gs := new(Gamescreen)
	gs.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	leveltje := new(CurrentLevel)
	leveltje.Entity = tl.NewEntity(1, 1, 1, 1)

	gs.Level.AddEntity(leveltje)

	return gs
}

func NewTitleScreen() *Titlescreen {
	ts := new(Titlescreen)
	ts.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	logofile, _ := ioutil.ReadFile("util/logo.text")
	logoEntity := tl.NewEntityFromCanvas(10, 3, tl.CanvasFromString(string(logofile)))

	ts.Text = []*tl.Text{
		tl.NewText(10, 20, "Press ENTER to start!", tl.ColorWhite, tl.ColorBlack),
		tl.NewText(10, 22, "Press BACKSPACE to start from last saved level!", tl.ColorWhite, tl.ColorBlack),
	}

	for _, v := range ts.Text {
		ts.Level.AddEntity(v)
	}

	ts.Level.AddEntity(logoEntity)
	return ts
}

func NewPlayer() *Player {
	player := new(Player)
	player.pCanvas = tl.NewCanvas(2, 1)
	player.pCanvas[0][0] = playercell1
	player.pCanvas[1][0] = playercell1
	player.pCoords = CheckPlayerPosition()
	player.Entity = tl.NewEntityFromCanvas(player.pCoords.X, player.pCoords.Y, player.pCanvas)

	return player
}
