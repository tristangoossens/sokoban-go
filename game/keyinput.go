package trisoban

import tl "github.com/JoelOtter/termloop"

func (ts *Titlescreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyEnter:
			player := NewPlayer()
			gs := NewGameScreen()
			gs.AddEntity(player)
			game.Screen().SetLevel(gs)
		}
	}
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		player.pCoords.X, player.pCoords.Y = player.Position()
		switch event.Key {
		case tl.KeyArrowUp:
			player.pCoords.Y = player.pCoords.Y - 1
			player.SetPosition(player.pCoords.X, player.pCoords.Y)

		case tl.KeyArrowDown:
			player.pCoords.Y = player.pCoords.Y + 1
			player.SetPosition(player.pCoords.X, player.pCoords.Y)

		case tl.KeyArrowLeft:
			player.pCoords.X = player.pCoords.X - 1
			player.SetPosition(player.pCoords.X, player.pCoords.Y)

		case tl.KeyArrowRight:
			player.pCoords.X = player.pCoords.X + 1
			player.SetPosition(player.pCoords.X, player.pCoords.Y)

		}
	}
}
