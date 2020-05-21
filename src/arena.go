package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

// Draw is a method for border that will print out the border with the given coordinates
func (border *Border) Draw(screen *tl.Screen) {
	for v := range border.bCoords {
		screen.RenderCell(v.X, v.Y, &tl.Cell{
			Fg: tl.ColorBlue,
			Ch: '▓',
		})
	}
}
