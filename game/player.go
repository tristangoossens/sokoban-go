package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

func (player *Player) Draw(screen *tl.Screen) {
	screen.RenderCell(58, 15, &tl.Cell{
		Fg: tl.ColorRed,
		Ch: '▓',
	})
	screen.RenderCell(58, 16, &tl.Cell{
		Fg: tl.ColorYellow,
		Ch: '░',
	})
	screen.RenderCell(58, 17, &tl.Cell{
		Fg: tl.ColorWhite,
		Ch: '▓',
	})
	screen.RenderCell(58, 18, &tl.Cell{
		Fg: tl.ColorBlue,
		Ch: '▓',
	})
	screen.RenderCell(player.X, player.Y, &tl.Cell{
		Fg: tl.ColorRed,
		Ch: '▓',
	})
}

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
