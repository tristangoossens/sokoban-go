package trisobaneditor

import (
	tl "github.com/JoelOtter/termloop"
)

// Coordinates struct that is inherited by every entity which contains the x and y coordinates.
type Coordinates struct {
	x, y int
}

// Area entity for the border of the bounds of the editor
type Area struct {
	*tl.Entity
	area map[Coordinates]int
}

// CreateArea create the border entity and return it.
func CreateArea() *Area {
	b := new(Area)
	b.Entity = tl.NewEntity(1, 1, 1, 1)

	return b
}

// Draw is a function that will draw the area border
func (a *Area) Draw(screen *tl.Screen) {
	// This for loop will range ArenaBorder containing the coordinates of the arenaborder and will print them out on the screen.
	for v := range a.area {
		screen.RenderCell(v.x, v.y, &tl.Cell{
			Fg: tl.ColorWhite,
			Ch: '░',
		})
	}
}
