package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

// Goal inherits the entity making it a drawable, it also inherits the coordinates struct.
type Goal struct {
	*tl.Entity
	Coordinates
	isActivated bool
}

// NewGoal creates an entity for the goal and returns a pointer to a goal
func NewGoal() *Goal {
	goal := new(Goal)
	goal.Entity = tl.NewEntity(1, 1, 1, 1)

	return goal
}

// Draw draws the goal entity. When the goal is not yet reached, the color will be yellow(When reached it updates to green).
func (goal *Goal) Draw(screen *tl.Screen) {
	if goal.isActivated {
		screen.RenderCell(goal.X, goal.Y, &tl.Cell{
			Fg: tl.ColorGreen,
			Ch: '▓',
		})
	} else {
		screen.RenderCell(goal.X, goal.Y, &tl.Cell{
			Fg: tl.ColorYellow,
			Ch: '░',
		})
	}
}
