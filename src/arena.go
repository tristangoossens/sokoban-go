package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

// Border inherits the entity making it a drawable. It also consists of the bordercoordinates in a map.
type Border struct {
	*tl.Entity
	bCoords map[Coordinates]int
}

// NewBorder creates an entity for the border and returns a pointer to a border
func NewBorder() *Border {
	border := new(Border)
	border.Entity = tl.NewEntity(1, 1, 1, 1)

	return border
}

// Draw is a method for border that will print out the border with the given coordinates
func (border *Border) Draw(screen *tl.Screen) {
	for v := range border.bCoords {
		screen.RenderCell(v.X, v.Y, &tl.Cell{
			Fg: tl.ColorBlue,
			Ch: '▓',
		})
	}
}
