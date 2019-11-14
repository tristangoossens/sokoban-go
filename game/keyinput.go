package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

var oldX int
var oldY int
var oldcrateX int
var oldcrateY int

// Tick listens to keyinput on the titlescreen, so if the enter key is pressed, the game will start.
func (ts *Titlescreen) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyEnter:
			gs = NewGameScreen()
			game.Screen().SetLevel(gs)
		}
	}
}

// Tick listens to the player input and handle them accordingly.
func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		player.X, player.Y = player.Position()
		oldX, oldY = player.Position()
		crate.X, crate.Y = crate.Position()
		oldcrateX, oldcrateY = crate.Position()
		switch event.Key {
		case tl.KeyArrowUp:
			player.CheckCollisions("UP")
		case tl.KeyArrowDown:
			player.CheckCollisions("DOWN")
		case tl.KeyArrowLeft:
			player.CheckCollisions("LEFT")
		case tl.KeyArrowRight:
			player.CheckCollisions("RIGHT")
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
		}
	}
}

// CheckCollisions will calculate collisions and when there is a collision handle them accordingly.
//
// Check 1: Check if the player is colliding with a crate.
//
// Check 2: Check if the player is colliding with the border.
//
// Check 3: Check if the crate is colliding with the border.
//
// Check 4: Check if the crate is colliding with the goal.
func (player *Player) CheckCollisions(direction string) {
	switch direction {
	case "UP":
		if player.CheckCrateCollision(oldX, oldY-1) {
			if crate.CheckGoalCollision(oldcrateX, oldcrateY-1) {
				crate.Y = crate.Y - 1
				crate.SetPosition(crate.X, crate.Y)
				LevelCompleted()
			}
			if crate.CheckBorderCollision(oldcrateX, oldcrateY-1) {
				player.Y = player.Y + 1
				player.SetPosition(player.X, player.Y)
				crate.SetPosition(crate.X, crate.Y)
			} else {
				crate.Y = crate.Y - 1
				crate.SetPosition(crate.X, crate.Y)
			}
		}
		if player.CheckBorderCollision(oldX, oldY-1) {
			player.SetPosition(player.X, player.Y)
		} else {
			player.Y = player.Y - 1
			player.SetPosition(player.X, player.Y)
		}
	case "DOWN":
		if player.CheckCrateCollision(oldX, oldY+1) {
			if crate.CheckGoalCollision(oldcrateX, oldcrateY+1) {
				crate.Y = crate.Y + 1
				crate.SetPosition(crate.X, crate.Y)
				LevelCompleted()
			}
			if crate.CheckBorderCollision(oldcrateX, oldcrateY+1) {
				player.Y = player.Y - 1
				player.SetPosition(player.X, player.Y)
				crate.SetPosition(crate.X, crate.Y)
			} else {
				crate.Y = crate.Y + 1
				crate.SetPosition(crate.X, crate.Y)
			}
		}
		if player.CheckBorderCollision(oldX, oldY+1) {
			player.SetPosition(player.X, player.Y)
		} else {
			player.Y = player.Y + 1
			player.SetPosition(player.X, player.Y)
		}
	case "LEFT":
		if player.CheckCrateCollision(oldX-1, oldY) {
			if crate.CheckGoalCollision(oldcrateX-1, oldcrateY) {
				crate.X = crate.X - 1
				crate.SetPosition(crate.X, crate.Y)
				LevelCompleted()
			}
			if crate.CheckBorderCollision(oldcrateX-1, oldcrateY) {
				player.X = player.X + 1
				player.SetPosition(player.X, player.Y)
				crate.SetPosition(crate.X, crate.Y)
			} else {
				crate.X = crate.X - 1
				crate.SetPosition(crate.X, crate.Y)
			}
		}
		if player.CheckBorderCollision(oldX-1, oldY) {
			player.SetPosition(player.X, player.Y)
		} else {
			player.X = player.X - 1
			player.SetPosition(player.X, player.Y)
		}
	case "RIGHT":
		if player.CheckCrateCollision(oldX+1, oldY) {
			if crate.CheckGoalCollision(oldcrateX+1, oldcrateY) {
				crate.X = crate.X + 1
				crate.SetPosition(crate.X, crate.Y)
				LevelCompleted()
			}
			if crate.CheckBorderCollision(oldcrateX+1, oldcrateY) {
				player.X = player.X - 1
				player.SetPosition(player.X, player.Y)
				crate.SetPosition(crate.X, crate.Y)
			} else {
				crate.X = crate.X + 1
				crate.SetPosition(crate.X, crate.Y)
			}
		}
		if player.CheckBorderCollision(oldX+1, oldY) {
			player.SetPosition(player.X, player.Y)
		} else {
			player.X = player.X + 1
			player.SetPosition(player.X, player.Y)
		}
	}
}
