package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

func (crate *Crate) CheckBorderCollision(x, y int) bool {
	c := Coordinates{
		X: x,
		Y: y,
	}

	_, exists := border.bCoords[c]

	if exists {
		return true
	}
	return false
}

func (crate *Crate) CheckGoalCollision(x, y int) bool {
	c := Coordinates{
		X: x,
		Y: y,
	}
	if goal.Coordinates == c {
		return true
	}
	return false
}

func (crate *Crate) Draw(screen *tl.Screen) {
	screen.RenderCell(crate.X, crate.Y, &tl.Cell{
		Fg: tl.ColorWhite,
		Ch: '▓',
	})
}
