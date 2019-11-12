package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

func (goal *Goal) Draw(screen *tl.Screen) {
	if crate.reachedGoal {
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
