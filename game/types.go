package trisoban

import tl "github.com/JoelOtter/termloop"

// Variables

var playercell1 = tl.Cell{
	Fg: tl.ColorGreen,
	Ch: '▓',
}

var gs *Gamescreen
var game *tl.Game
var b *LevelBorder

var prevX int
var prevY int

// Levels

type Titlescreen struct {
	tl.Level
	Logo *tl.Entity
	Text []*tl.Text
}

type Gamescreen struct {
	tl.Level
	PlayerEntity *Player
	BorderEntity *LevelBorder
}

// Entities

type LevelBorder struct {
	*tl.Entity
	aCoords map[Coordinates]int
}

type Player struct {
	*tl.Entity
	pCoords Coordinates
}

// Other

type Coordinates struct {
	X int
	Y int
}
