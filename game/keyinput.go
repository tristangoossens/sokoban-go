package trisoban

import tl "github.com/JoelOtter/termloop"

func (ts *Titlescreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyEnter:
			gs = NewGameScreen()
			game.Screen().SetLevel(gs)
		}
	}
}

// func (player *Player) Tick(event tl.Event) {
// 	if event.Type == tl.EventKey {
// 		x, y := player.Position()
// 		xx, yy := player.Position()
// 		switch event.Key {
// 		case tl.KeyArrowUp:
// 			player.MoveUp(player.pCoords)
// 		case tl.KeyArrowDown:

// 			player.SetPosition(player.pCoords.X, player.pCoords.Y)

// 			prevX = player.pCoords.X
// 			prevY = player.pCoords.Y - 1

// 		case tl.KeyArrowLeft:
// 			player.pCoords.X = player.pCoords.X - 1
// 			player.SetPosition(player.pCoords.X, player.pCoords.Y)

// 			prevX = player.pCoords.X + 1
// 			prevY = player.pCoords.Y

// 		case tl.KeyArrowRight:
// 			player.pCoords.X = player.pCoords.X + 1
// 			player.SetPosition(player.pCoords.X, player.pCoords.Y)

// 			prevX = player.pCoords.X - 1
// 			prevY = player.pCoords.Y

// 		}
// 	}
// }

// func (player *Player) MoveUp(c []Coordinates) {
// 	for i, v := range c {
// 		v.Y = v.Y - 1

// 		prevX = v.X
// 		prevY = v.Y + 1

// 		player.SetPosition(v.X, v.Y)
// 	}
// }
