package trisoban

import (
	"bufio"
	"fmt"
	tl "github.com/JoelOtter/termloop"
	"log"
	"os"
	"strings"
)

func NewCrate() *Crate {
	crate := new(Crate)
	crate.Entity = tl.NewEntity(1, 1, 1, 1)

	return crate
}

func NewBorder() *Border {
	border := new(Border)
	border.Entity = tl.NewEntity(1, 1, 1, 1)

	return border
}

func NewPlayer() *Player {
	player := new(Player)
	player.Entity = tl.NewEntity(1, 1, 1, 1)

	return player
}

func NewGoal() *Goal {
	goal := new(Goal)
	goal.Entity = tl.NewEntity(1, 1, 1, 1)

	return goal
}

func MapLevel() {
	var startx = 7 // Determines at which coordinate the rendering starts
	var starty = 5 // Determines at which coordinate the rendering starts

	border.bCoords = make(map[Coordinates]int)

	file, err := os.Open(fmt.Sprintf("util/levels/level%d.txt", CurrentLevel))
	if err != nil {
		log.Fatalln(file)
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
