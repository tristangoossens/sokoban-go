package trisoban

import (
	tl "github.com/JoelOtter/termloop"
)

func NewCrate() *Crate {
	crate := new(Crate)
	crate.Entity = tl.NewEntity(1, 1, 1, 1)
	cCoords := MapCrate()
	crate.Coordinates = cCoords
	crate.SetPosition(crate.X, crate.Y)

	return crate
}

func NewBorder() *Border {
	border := new(Border)
	border.Entity = tl.NewEntity(1, 1, 1, 1)
	cMap := MapBorder()
	border.bCoords = cMap

	return border
}

func NewPlayer() *Player {
	player := new(Player)
	player.Entity = tl.NewEntity(1, 1, 1, 1)
	x, y := GetPlayerPosition()
	player.X = x
	player.Y = y
	player.SetPosition(player.X, player.Y)

	return player
}

func NewGoal() *Goal {
	goal := new(Goal)
	goal.Entity = tl.NewEntity(1, 1, 1, 1)
	gCoords := MapGoal()
	goal.Coordinates = gCoords

	return goal
}
