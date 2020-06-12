package trisobaneditor

import (
	tl "github.com/JoelOtter/termloop"
)

// Player player object
type Player struct {
	*tl.Entity
	Coordinates
}

// CreatePlayer create a player entity
func CreatePlayer(c Coordinates) *Player {
	player := new(Player)
	player.Entity = tl.NewEntity(1, 1, 1, 1)
	player.Coordinates = c

	return player
}

// Draw draw the player
func (p *Player) Draw(screen *tl.Screen) {
	screen.RenderCell(p.x, p.y, &tl.Cell{
		Fg: tl.ColorRed,
		Ch: '▓',
	})
}
