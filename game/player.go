package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

// Draw will print the entity of the player and update it every tick. It also prints the cells for the instructions.
func (player *Player) Draw(screen *tl.Screen) {
	// Instructions:

	/// Player
	screen.RenderCell(58, 15, &tl.Cell{
		Fg: tl.ColorRed,
		Ch: '▓',
	})
	/// Goal
	screen.RenderCell(58, 16, &tl.Cell{
		Fg: tl.ColorYellow,
		Ch: '░',
	})
	/// Crate
	screen.RenderCell(58, 17, &tl.Cell{
		Fg: tl.ColorWhite,
		Ch: '▓',
	})
	/// Border
	screen.RenderCell(58, 18, &tl.Cell{
		Fg: tl.ColorBlue,
		Ch: '▓',
	})

	// Player entity
	screen.RenderCell(player.X, player.Y, &tl.Cell{
		Fg: tl.ColorRed,
		Ch: '▓',
	})
}

// CheckBorderCollision will check if the player is colliding with the border.
func (player *Player) CheckBorderCollision(x int, y int) bool {
	c := Coordinates{
		X: x,
		Y: y,
	}
	_, exists := border.bCoords[c]

	if exists {
		return true
	}
	return false
}

// CheckCrateCollision will check if the player is colliding with the crate.
func (player *Player) CheckCrateCollision(x int, y int) bool {
	c := Coordinates{
		X: x,
		Y: y,
	}
	if crate.Coordinates == c {
		return true
	}
	return false
}
