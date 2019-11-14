package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

// CheckBorderCollision will check if the crate has a collision with the border. If so, this will return true.
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

// CheckGoalCollision check if the crate is colliding with the goal, if so it will return true.
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

// Draw draws the crate given the coordinates, this will update every tick.
func (crate *Crate) Draw(screen *tl.Screen) {
	screen.RenderCell(crate.X, crate.Y, &tl.Cell{
		Fg: tl.ColorWhite,
		Ch: '▓',
	})
}
