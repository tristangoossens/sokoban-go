package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

// Player inherits the entity making it a drawable, it also inherits the border and coordinates struct.
type Player struct {
	*tl.Entity
	Border
	Coordinates
}

// NewPlayer creates an entity for the player and returns a pointer to a player
func NewPlayer() *Player {
	player := new(Player)
	player.Entity = tl.NewEntity(1, 1, 1, 1)

	return player
}

// Draw will print the entity of the player and update it every tick. It also prints the cells for the instructions.
func (player *Player) Draw(screen *tl.Screen) {
	screen.RenderCell(player.X, player.Y, &tl.Cell{
		Fg: tl.ColorRed,
		Ch: '▓',
	})
}

// CheckBorderCollision will check if the player is colliding with the border.
func (player *Player) CheckBorderCollision(dir string) bool {
	_, exists := col.Border.bCoords[player.CalculatePlayerCoordinates(dir)]

	if exists {
		return true
	}
	return false
}

// CheckCrateCollision will check if the player is colliding with the crate.
func (player *Player) CheckCrateCollision(c *Crate, dir string) bool {
	if player.CalculatePlayerCoordinates(dir) == c.Coordinates {
		return true
	}
	return false
}

// CheckActivatedGoalCollision will check if the player is colliding an activated goal.
func (player *Player) CheckActivatedGoalCollision(g *Goal, dir string) bool {
	if g.isActivated && player.CalculatePlayerCoordinates(dir) == g.Coordinates {
		return true
	}
	return false
}

// CalculatePlayerCoordinates calculates the coordinates from the player for current move
func (player *Player) CalculatePlayerCoordinates(dir string) Coordinates {
	var co Coordinates
	switch dir {
	case "up":
		co = Coordinates{
			X: player.X,
			Y: player.Y - 1,
		}
	case "down":
		co = Coordinates{
			X: player.X,
			Y: player.Y + 1,
		}
	case "left":
		co = Coordinates{
			X: player.X - 1,
			Y: player.Y,
		}
	case "right":
		co = Coordinates{
			X: player.X + 1,
			Y: player.Y,
		}
	}
	return co
}

// MovePlayer move the player if there is no collision, else it will take a step back(entity doesn't move)
func (player *Player) MovePlayer(dir string, colliding bool) {
	if colliding {
		switch dir {
		case "up":
			player.Y = col.Player.Y + 1
		case "down":
			player.Y = col.Player.Y - 1
		case "left":
			player.X = col.Player.X + 1
		case "right":
			player.X = col.Player.X - 1
		}
	} else {
		switch dir {
		case "up":
			player.Y = col.Player.Y - 1
		case "down":
			player.Y = col.Player.Y + 1
		case "left":
			player.X = col.Player.X - 1
		case "right":
			player.X = col.Player.X + 1
		}
	}
	player.SetPosition(player.X, player.Y)
}
