package trisoban

import (
	"io/ioutil"

	tl "github.com/JoelOtter/termloop"
)

func NewGameScreen() *Gamescreen {
	gs := new(Gamescreen)
	gs.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	player := NewPlayer()
	b = NewLevelBorder()

	gs.AddEntity(b)
	gs.AddEntity(player)

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
	player.pCoords = CheckPlayerPosition()
	canvas := tl.NewCanvas(2, 1)
	canvas[0][0] = playercell1
	canvas[1][0] = playercell1
	player.Entity = tl.NewEntityFromCanvas(0, 0, canvas)

	// player.Entity = tl.NewEntityFromCanvas(player.pCoords.X1, player.pCoords.Y1, player.pCanvas)

	return player
}

func NewLevelBorder() *LevelBorder {
	levelborder := new(LevelBorder)
	levelborder.Entity = tl.NewEntity(1, 1, 1, 1)
	levelborder.aCoords = map[Coordinates]int{}

	return levelborder
}
