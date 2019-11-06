package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

var game *tl.Game

func StartGame() {
	game = tl.NewGame()

	ts := NewTitleScreen()
	// player := NewPlayer()

	// leveltje := new(Level)
	// leveltje.Entity = tl.NewEntity(1, 1, 1, 1)

	game.Screen().SetLevel(ts)
	game.Screen().SetFps(40)
	game.Start()
}

func (ts *Titlescreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyEnter:
			gs := NewGameScreen()
			game.Screen().SetLevel(gs)
		}
	}
}


