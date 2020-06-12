package trisobaneditor

import (
	tl "github.com/JoelOtter/termloop"
)

// Cursor cursor
type Cursor struct {
	*tl.Entity
	Coordinates
}

// CreateCursor create a cursor object.
func CreateCursor() *Cursor {
	c := new(Cursor)
	c.Entity = tl.NewEntity(1, 1, 1, 1)
	c.Coordinates = Coordinates{
		x: 60,
		y: 15,
	}

	return c
}

// DrawPlayer draw a player on the current cursor location
func (c *Cursor) DrawPlayer() {
	if !sc.df.CheckDuplicateEntity(c.Coordinates) {
		if sc.p != nil {
			sc.RemoveEntity(sc.p)
			delete(sc.df.fm, c.Coordinates)
			sc.df.fm[sc.p.Coordinates] = ""
		}
		sc.p = CreatePlayer(c.Coordinates)
		sc.df.fm[c.Coordinates] = "Player"
		sc.requirements[0].SetColor(tl.ColorGreen, tl.ColorBlack)
		sc.AddEntity(sc.p)
	}
}

// DrawBorder draw a border on the current cursor location
func (c *Cursor) DrawBorder() {
	if !sc.df.CheckDuplicateEntity(c.Coordinates) {
		b := CreateBorder(c.Coordinates)
		sc.b = append(sc.b, b)
		sc.df.fm[c.Coordinates] = "Border"
		sc.AddEntity(b)
	}
}

// DrawGoal draw a goal on the current cursor location
func (c *Cursor) DrawGoal() {
	if !sc.df.CheckDuplicateEntity(c.Coordinates) {
		g := CreateGoal(c.Coordinates)
		sc.g = append(sc.g, g)
		sc.df.fm[c.Coordinates] = "Goal"
		sc.AddEntity(g)
	}
}

// DrawCrate draw a crate on the current cursor location
func (c *Cursor) DrawCrate() {
	if !sc.df.CheckDuplicateEntity(c.Coordinates) {
		crate := CreateCrate(c.Coordinates)
		sc.cs = append(sc.cs, crate)
		sc.df.fm[c.Coordinates] = "Crate"
		sc.AddEntity(crate)
	}
}

// Move move the cursor
func (c *Cursor) Move(dir string) {
	switch dir {
	case "up":
		if !c.CheckAreaCollision(dir) {
			c.y = c.y - 1
			c.SetPosition(c.x, c.y)
		}
	case "down":
		if !c.CheckAreaCollision(dir) {
			c.y = c.y + 1
			c.SetPosition(c.x, c.y)
		}
	case "left":
		if !c.CheckAreaCollision(dir) {
			c.x = c.x - 1
			c.SetPosition(c.x, c.y)
		}
	case "right":
		if !c.CheckAreaCollision(dir) {
			c.x = c.x + 1
			c.SetPosition(c.x, c.y)
		}
	}
}

// CalculateCursorCoordinates calculates the coordinates for the cursor for its next move.
func (c *Cursor) CalculateCursorCoordinates(dir string) Coordinates {
	var co Coordinates
	switch dir {
	case "up":
		co = Coordinates{
			x: c.x,
			y: c.y - 1,
		}
	case "down":
		co = Coordinates{
			x: c.x,
			y: c.y + 1,
		}
	case "left":
		co = Coordinates{
			x: c.x - 1,
			y: c.y,
		}
	case "right":
		co = Coordinates{
			x: c.x + 1,
			y: c.y,
		}
	}
	return co
}

// CheckAreaCollision check a collision with the area border
func (c *Cursor) CheckAreaCollision(dir string) bool {
	_, exists := sc.a.area[c.CalculateCursorCoordinates(dir)]

	if exists {
		return true
	}
	return false
}

// DeleteEntity delete an entity using the eraser
func (c *Cursor) DeleteEntity(co Coordinates) {
	v, exists := sc.df.fm[co]
	if exists {
		switch v {
		case "Player":
			if sc.p != nil {
				sc.df.fm[sc.p.Coordinates] = ""
				if sc.p.Coordinates == c.Coordinates {
					sc.RemoveEntity(sc.p)
				}
			}
		case "Border":
			sc.df.fm[c.Coordinates] = ""
			for i, v := range sc.b {
				border := v.FindBorder(c.Coordinates)
				if border != nil {
					sc.b = border.RemoveBorder(sc.b, i)
					sc.RemoveEntity(border)
				}
			}
		case "Goal":
			sc.df.fm[c.Coordinates] = ""
			for i, v := range sc.g {
				goal := v.FindGoal(c.Coordinates)
				if goal != nil {
					sc.g = goal.RemoveGoal(sc.g, i)
					sc.RemoveEntity(goal)
				}
			}
		case "Crate":
			sc.df.fm[c.Coordinates] = ""
			for i, v := range sc.cs {
				crate := v.FindCrate(c.Coordinates)
				if crate != nil {
					sc.cs = crate.RemoveCrate(sc.cs, i)
					sc.RemoveEntity(crate)
				}
			}
		}
	}
}

// Draw draw a cursor
func (c *Cursor) Draw(screen *tl.Screen) {
	screen.RenderCell(c.x, c.y, &tl.Cell{
		Fg: tl.ColorGreen,
		Bg: tl.ColorWhite,
		Ch: '▒',
	})

	if len(sc.cs) == len(sc.g) && len(sc.g) != 0 && len(sc.cs) != 0 {
		sc.requirements[1].SetColor(tl.ColorGreen, tl.ColorBlack)
	} else {
		sc.requirements[1].SetColor(tl.ColorRed, tl.ColorBlack)
	}
}
