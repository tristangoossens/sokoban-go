package trisobaneditor

import tl "github.com/JoelOtter/termloop"

// Crate crate object
type Crate struct {
	*tl.Entity
	Coordinates
}

// CreateCrate create a crate entity
func CreateCrate(c Coordinates) *Crate {
	crate := new(Crate)
	crate.Entity = tl.NewEntity(1, 1, 1, 1)
	crate.Coordinates = c

	return crate
}

// RemoveCrate remove a crate from the entity slice
func (c *Crate) RemoveCrate(slice []*Crate, i int) []*Crate {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// FindCrate find the crate with the given coordinates
func (c *Crate) FindCrate(co Coordinates) *Crate {
	if c.Coordinates == co {
		return c
	}
	return nil
}

// Draw draw the border
func (c *Crate) Draw(screen *tl.Screen) {
	screen.RenderCell(c.x, c.y, &tl.Cell{
		Fg: tl.ColorWhite,
		Ch: '▓',
	})
}
