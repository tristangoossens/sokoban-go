package trisoban

import (
	"bufio"
	"log"
	"os"
	"strings"

	tl "github.com/JoelOtter/termloop"
)

func MapBorder() map[Coordinates]int {
	var startx = 7 // Determines at which coordinate the rendering starts
	var starty = 5 // Determines at which coordinate the rendering starts

	var cMap = make(map[Coordinates]int)

	file, err := os.Open("util/levels.txt")
	if err != nil {
		log.Fatalln(file)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		xline := strings.Split(scanner.Text(), "")
		for i, v := range xline {
			switch v {
			case "#":
				cMap[Coordinates{
					X: startx + i,
					Y: starty,
				}] = 1
			}
		}
		starty++
	}
	file.Close()
	return cMap
}

func (border *Border) Draw(screen *tl.Screen) {
	for v := range border.bCoords {
		screen.RenderCell(v.X, v.Y, &tl.Cell{
			Fg: tl.ColorBlue,
			Ch: '▓',
		})
	}
}
