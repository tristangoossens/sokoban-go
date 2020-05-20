package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

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
