package trisoban

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	tl "github.com/JoelOtter/termloop"
)

// NewEntityCollection creates a collection of all the entities in the game
func NewEntityCollection() *EntityCollection {
	col := new(EntityCollection)
	col.Crate = []*Crate{}
	col.Goal = []*Goal{}
	col.Player = NewPlayer()
	col.Border = NewBorder()

	return col
}

// NewCrate creates an entity for the crate and returns a pointer to a crate
func NewCrate() *Crate {
	crate := new(Crate)
	crate.Entity = tl.NewEntity(1, 1, 1, 1)

	return crate
}

// NewBorder creates an entity for the border and returns a pointer to a border
func NewBorder() *Border {
	border := new(Border)
	border.Entity = tl.NewEntity(1, 1, 1, 1)

	return border
}

// NewPlayer creates an entity for the player and returns a pointer to a player
func NewPlayer() *Player {
	player := new(Player)
	player.Entity = tl.NewEntity(1, 1, 1, 1)

	return player
}

// NewGoal creates an entity for the goal and returns a pointer to a goal
func NewGoal() *Goal {
	goal := new(Goal)
	goal.Entity = tl.NewEntity(1, 1, 1, 1)

	return goal
}

// MapLevel reads from a level file and maps the coordinates for all of the entities
func MapLevel() {
	var startx = 7 // Determines at which coordinate the rendering starts
	var starty = 5 // Determines at which coordinate the rendering starts

	border.bCoords = make(map[Coordinates]int)

	file, err := os.Open(fmt.Sprintf("util/levels/level%d.txt", CurrentLevel))
	if err != nil {
		log.Fatalf("Level with number %d was not found", CurrentLevel)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xline := strings.Split(scanner.Text(), "")
		for i, v := range xline {
			switch v {
			case "#":
				border.bCoords[Coordinates{
					X: startx + i,
					Y: starty,
				}] = 1
			case "G":
				goal.Coordinates = Coordinates{
					X: startx + i,
					Y: starty,
				}
			case "P":
				player.Coordinates = Coordinates{
					X: startx + i,
					Y: starty,
				}
				player.SetPosition(player.X, player.Y)
			case "C":
				crate.Coordinates = Coordinates{
					X: startx + i,
					Y: starty,
				}
				crate.SetPosition(crate.X, crate.Y)
			}
		}
		starty++
	}
	file.Close()
}
