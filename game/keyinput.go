package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

func (ts *Titlescreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyEnter:
			gs := NewGameScreen()
			game.Screen().SetLevel(gs)
		}
	}
}

func (player *Player) Tick(event tl.Event) {
	var oldX int
	var oldY int
	if event.Type == tl.EventKey {
		player.X, player.Y = player.Position()
		oldX, oldY = player.Position()
		switch event.Key {
		case tl.KeyArrowUp:
			if player.CheckBorderCollision(oldX, oldY-1) {
				player.SetPosition(player.X, player.Y)
			} else {
				player.Y = player.Y - 1
				player.SetPosition(player.X, player.Y)
			}
		case tl.KeyArrowDown:
			if player.CheckBorderCollision(oldX, oldY+1) {
				player.SetPosition(player.X, player.Y)
			} else {
				player.Y = player.Y + 1
				player.SetPosition(player.X, player.Y)
			}
		case tl.KeyArrowLeft:
			if player.CheckBorderCollision(oldX-1, oldY) {
				player.SetPosition(player.X, player.Y)
			} else {
				player.X = player.X - 1
				player.SetPosition(player.X, player.Y)
			}
		case tl.KeyArrowRight:
			if player.CheckBorderCollision(oldX+1, oldY) {
				player.SetPosition(player.X, player.Y)
			} else {
				player.X = player.X + 1
				player.SetPosition(player.X, player.Y)
			}
		}
	}
}
