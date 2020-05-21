package trisoban

import (
	"io/ioutil"
	"strconv"

	tl "github.com/JoelOtter/termloop"
)

// Tick listens to keyinput on the titlescreen, so if the enter key is pressed, the game will start.
func (ts *Titlescreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyEnter:
			CurrentLevel = 1
			gs = NewGameScreen()
			game.Screen().SetLevel(gs)

		case tl.KeyEsc:
			CurrentLevel = LoadLevel()
			gs = NewGameScreen()
			game.Screen().SetLevel(gs)
		}
	}
}

// Tick listens to the gamescreen input and handles it accordingly.
func (g *Gamescreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowUp:
			g.Move("up")
		case tl.KeyArrowDown:
			g.Move("down")
		case tl.KeyArrowLeft:
			g.Move("left")
		case tl.KeyArrowRight:
			g.Move("right")
		case tl.KeyF3:
			RestartLevel()
		case tl.KeyF1:
			if gs.CheckLevelCompletion() {
				gs.ChangeLevel("next")
			}
		case tl.KeyF2:
			if gs.CheckLevelCompletion() {
				gs.ChangeLevel("previous")
			}
		case tl.KeyInsert:
			ioutil.WriteFile("util/loadgame.txt", []byte(strconv.Itoa(CurrentLevel)), 0644)
		}
	}
}
