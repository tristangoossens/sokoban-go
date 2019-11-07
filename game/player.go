package trisoban

import (
	"bufio"
	tl "github.com/JoelOtter/termloop"
	"log"
	"os"
	"strings"
)

func (player *Player) Draw(screen *tl.Screen) {
	player.Entity.Draw(screen)
}

func CheckPlayerPosition() Coordinates {
	file, err := os.Open("util/levels.txt")
	if err != nil {
		log.Fatalln(file)
	}
	scanner := bufio.NewScanner(file)
	var startx = 7 // Determines at which coordinate the rendering starts
	var starty = 5 // Determines at which coordinate the rendering starts
	var c Coordinates
	for scanner.Scan() {
		xline := strings.Split(scanner.Text(), "")
		for i, v := range xline {
			switch v {
			case "P":
				c = Coordinates{
					X: startx + i + 1,
					Y: starty,
				}
			}
		}
		starty++
	}
	file.Close()
	return c
}
