package trisoban

import tl "github.com/JoelOtter/termloop"

// Variables
var game *tl.Game
var border *Border

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
}

// Entities

type Player struct {
	*tl.Entity
	Border
	Coordinates
}

type Border struct {
	*tl.Entity
	bCoords map[Coordinates]int
}

// Other

type Coordinates struct {
	X int
	Y int
}
