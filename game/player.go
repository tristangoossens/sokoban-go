package trisoban

import tl "github.com/JoelOtter/termloop"

func NewPlayer() *Player {
	player := new(Player)
	player.Entity = tl.NewEntity(1, 1, 2, 1)
	player.P.X = 10
	player.P.Y = 10

	return player
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		player.P.X, player.P.Y = player.Position()
		switch event.Key {
		case tl.KeyArrowUp:
			player.P.Y = player.P.Y - 1
			player.SetPosition(player.P.X, player.P.Y)
		case tl.KeyArrowDown:
			player.P.Y = player.P.Y + 1
			player.SetPosition(player.P.X, player.P.Y)
		case tl.KeyArrowLeft:
			player.P.X = player.P.X - 1
			player.SetPosition(player.P.X, player.P.Y)
		case tl.KeyArrowRight:
			player.P.X = player.P.X + 1
			player.SetPosition(player.P.X, player.P.Y)
		}
	}
}

func (player *Player) Draw(screen *tl.Screen) {
	screen.RenderCell(player.P.X, player.P.Y, &tl.Cell{
		Bg: tl.ColorYellow,
		Fg: tl.ColorWhite,
		Ch: '▓',
	})
}
