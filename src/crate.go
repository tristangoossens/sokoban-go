package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

// Crate inherits the entity making it a drawable, it also inherits the coordinates struct, and has a bool for the rachedgoal check.
type Crate struct {
	*tl.Entity
	Coordinates
}

// NewCrate creates an entity for the crate and returns a pointer to a crate
func NewCrate() *Crate {
	crate := new(Crate)
	crate.Entity = tl.NewEntity(1, 1, 1, 1)

	return crate
}

// Draw draws the crate given the coordinates, this will update every tick.
func (crate *Crate) Draw(screen *tl.Screen) {
	screen.RenderCell(crate.X, crate.Y, &tl.Cell{
		Fg: tl.ColorWhite,
		Ch: '▓',
	})
}

// CheckBorderCollision will check if the crate has a collision with the border. If so, this will return true.
func (crate *Crate) CheckBorderCollision() bool {
	_, exists := col.Border.bCoords[crate.Coordinates]

	if exists {
		return true
	}
	return false
}

//CheckCrateCollision check whether the given crate is colliding with any other crate
func (crate *Crate) CheckCrateCollision(crates []*Crate) bool {
	for _, v := range crates {
		if crate.Coordinates == v.Coordinates {
			return true
		}
	}
	return false
}

// CheckGoalCollision check if the crate is colliding with the goal, if so it will return true.
func (crate *Crate) CheckGoalCollision(g *Goal) bool {
	if crate.Coordinates == g.Coordinates {
		return true
	}
	return false
}

// CheckActivatedGoalCollision check if the crate is colliding with an activated goal, if so it will return true.
func (crate *Crate) CheckActivatedGoalCollision(g *Goal) bool {
	if g.isActivated && crate.Coordinates == g.Coordinates {
		return true
	}
	return false
}

// MoveCrate move the crate if there is no collision, else it will take a step back(entity doesn't move)
func (crate *Crate) MoveCrate(dir string, colliding bool) {
	if colliding {
		switch dir {
		case "up":
			crate.Y = crate.Y + 1
		case "down":
			crate.Y = crate.Y - 1
		case "left":
			crate.X = crate.X + 1
		case "right":
			crate.X = crate.X - 1
		}
	} else {
		switch dir {
		case "up":
			crate.Y = crate.Y - 1
		case "down":
			crate.Y = crate.Y + 1
		case "left":
			crate.X = crate.X - 1
		case "right":
			crate.X = crate.X + 1
		}
	}
	crate.SetPosition(crate.X, crate.Y)
}

// RemoveCrate remove a crate from the entity slice
func (crate *Crate) RemoveCrate(slice []*Crate, i int) []*Crate {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// CopyCrateSlice create a duplicate of the crate entity slice
func (crate *Crate) CopyCrateSlice() []*Crate {
	slicecopy := make([]*Crate, len(col.Crates))
	copy(slicecopy, col.Crates)

	return slicecopy
}
