package trisobaneditor

import tl "github.com/JoelOtter/termloop"

// Border border entity
type Border struct {
	*tl.Entity
	Coordinates
}

// CreateBorder create a border entity
func CreateBorder(c Coordinates) *Border {
	border := new(Border)
	border.Entity = tl.NewEntity(1, 1, 1, 1)
	border.Coordinates = c
	return border
}

// FindBorder find the border entity on the given coordinates
func (b *Border) FindBorder(c Coordinates) *Border {
	if b.Coordinates == c {
		return b
	}
	return nil
}

// RemoveBorder remove a border from the entity slice
func (b *Border) RemoveBorder(slice []*Border, i int) []*Border {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// Draw draw the border
func (b *Border) Draw(screen *tl.Screen) {
	screen.RenderCell(b.x, b.y, &tl.Cell{
		Fg: tl.ColorBlue,
		Ch: '▓',
	})
}
