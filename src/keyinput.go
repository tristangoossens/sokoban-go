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

		case tl.KeyBackspace:
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
			if crate.reachedGoal {
				ChangeLevel("NEXT")
			}
		case tl.KeyF2:
			if crate.reachedGoal {
				ChangeLevel("PREVIOUS")
			}
		case tl.KeyInsert:
			ioutil.WriteFile("util/loadgame.txt", []byte(strconv.Itoa(CurrentLevel)), 0644)
		}
	}
}

// Move calculate and check collisions for the current move
func (g *Gamescreen) Move(dir string) {
	for ci, cv := range col.Crates {
		if col.Player.CheckCrateCollision(cv, dir) {
			col.Player.SetPosition(col.Player.X, col.Player.Y)
			cv.MoveCrate(dir, false)
			crateslicecopy := cv.CopyCrateSlice()
			if cv.CheckBorderCollision() {
				col.Player.MovePlayer(dir, true)
				cv.MoveCrate(dir, true)
			}
			if cv.CheckCrateCollision(cv.RemoveCrate(crateslicecopy, ci)) {
				cv.MoveCrate(dir, true)
				col.Player.MovePlayer(dir, true)
			}
			for _, gv := range col.Goals {
				if cv.CheckGoalCollision(gv) {
					if cv.CheckActivatedGoalCollision(gv) {
						col.Player.MovePlayer(dir, true)
						cv.MoveCrate(dir, true)
					} else {
						gv.isActivated = true
						gs.RemoveEntity(cv)
						col.Crates = cv.RemoveCrate(col.Crates, ci)
					}
				}
			}
		}
	}
	if col.Player.CheckBorderCollision(dir) {
		col.Player.MovePlayer(dir, true)
	}
	for _, gv := range col.Goals {
		if col.Player.CheckActivatedGoalCollision(gv, dir) {
			col.Player.MovePlayer(dir, true)
		}
	}
	col.Player.MovePlayer(dir, false)
}
