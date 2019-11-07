package trisoban

import (
	"bufio"
	tl "github.com/JoelOtter/termloop"
	"log"
	"os"
	"strings"
)

func (player *Player) Draw(screen *tl.Screen) {
	// if player.BorderCollision() {
	// 	player.PreviousPosition()
	// }
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
	var coord Coordinates
	for scanner.Scan() {
		xline := strings.Split(scanner.Text(), "")
		for i, v := range xline {
			switch v {
			case "P":
				coord = Coordinates{
					X: startx + i,
					Y: starty,
				}
			}
		}
		starty++
	}
	file.Close()
	return coord
}

func (player *Player) PreviousPosition() {
	player.SetPosition(prevX, prevY)
}

func (player *Player) BorderCollision() bool {
	if b.Contains(player.pCoords) {
		return true
	}
	return false
}
