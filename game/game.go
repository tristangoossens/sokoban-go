package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

var game *tl.Game

func StartGame() {
	game = tl.NewGame()

	ts := NewTitleScreen()

	game.Screen().SetLevel(ts)
	game.Screen().SetFps(40)
	game.Start()
}
