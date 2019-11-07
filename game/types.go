package trisoban

import tl "github.com/JoelOtter/termloop"

type CurrentLevel struct {
	*tl.Entity
}

type Coordinates struct {
	X int
	Y int
}

type Player struct {
	*tl.Entity
	pCoords Coordinates
	pCanvas tl.Canvas
}

type Titlescreen struct {
	tl.Level
	Logo *tl.Entity
	Text []*tl.Text
}

type Gamescreen struct {
	tl.Level
}

var playercell1 = tl.Cell{
	Fg: tl.ColorGreen,
	Ch: '▓',
}
