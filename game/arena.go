package trisoban

import (
	"bufio"
	"log"
	"os"
	"strings"

	tl "github.com/JoelOtter/termloop"
)

func (border *LevelBorder) Draw(screen *tl.Screen) {
	file, err := os.Open("util/levels.txt")
	if err != nil {
		log.Fatalln(file)
	}
	scanner := bufio.NewScanner(file)
	var startx = 7 // Determines at which coordinate the rendering starts
	var starty = 5 // Determines at which coordinate the rendering starts

	for scanner.Scan() {
		xline := strings.Split(scanner.Text(), "")
		for i, v := range xline {
			switch v {
			case "#":
				border.aCoords[Coordinates{startx + i, starty}] = 1
				for v := range border.aCoords {
					screen.RenderCell(v.X, v.Y, &tl.Cell{
						Fg: tl.ColorBlue,
						Ch: '▓',
					})
				}
			}
		}
		starty++
	}
	file.Close()
}

func (border *LevelBorder) Contains(c Coordinates) bool {
	_, exists := border.aCoords[c]
	return exists
}
