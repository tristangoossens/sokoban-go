package trisoban

import tl "github.com/JoelOtter/termloop"

type Level struct {
	*tl.Entity
}

type Coordinates struct {
	X int
	Y int
}

type Player struct {
	*tl.Entity
	P Coordinates
}
