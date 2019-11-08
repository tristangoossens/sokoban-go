package trisoban

import (
	"io/ioutil"

	tl "github.com/JoelOtter/termloop"
)

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

func NewGameScreen() *Gamescreen {
	gs := new(Gamescreen)
	gs.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	border = NewBorder()
	player := NewPlayer()
	gs.AddEntity(border)
	gs.AddEntity(player)

	return gs
}

func NewBorder() *Border {
	border := new(Border)
	border.Entity = tl.NewEntity(1, 1, 1, 1)
	cMap := MapBorder()
	border.bCoords = cMap

	return border
}

func NewPlayer() *Player {
	player := new(Player)
	player.Entity = tl.NewEntity(1, 1, 1, 1)
	x, y := GetPlayerPosition()
	player.X = x
	player.Y = y
	player.SetPosition(player.X, player.Y)

	return player
}