package trisobaneditor

import tl "github.com/JoelOtter/termloop"

// Goal goal object
type Goal struct {
	*tl.Entity
	Coordinates
}

// CreateGoal create a goal entity
func CreateGoal(c Coordinates) *Goal {
	goal := new(Goal)
	goal.Entity = tl.NewEntity(1, 1, 1, 1)
	goal.Coordinates = c

	return goal
}

// RemoveGoal remove a goal from the entity slice
func (g *Goal) RemoveGoal(slice []*Goal, i int) []*Goal {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// FindGoal find the goal entity on the given coordinates
func (g *Goal) FindGoal(c Coordinates) *Goal {
	if g.Coordinates == c {
		return g
	}
	return nil
}

// Draw draw the goal
func (g *Goal) Draw(screen *tl.Screen) {
	screen.RenderCell(g.x, g.y, &tl.Cell{
		Fg: tl.ColorYellow,
		Ch: '░',
	})
}
