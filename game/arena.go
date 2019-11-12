package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

func (border *Border) Draw(screen *tl.Screen) {
	for v := range border.bCoords {
		screen.RenderCell(v.X, v.Y, &tl.Cell{
			Fg: tl.ColorBlue,
			Ch: '▓',
		})
	}
}
