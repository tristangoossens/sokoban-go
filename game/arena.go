package trisoban

import (
	"bufio"
	"log"
	"os"
	"strings"

	tl "github.com/JoelOtter/termloop"
)

func (level *CurrentLevel) Draw(screen *tl.Screen) {
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
				screen.RenderCell(startx+i+1, starty, &tl.Cell{
					Fg: tl.ColorBlue,
					Ch: '▓',
				})
			}
		}
		starty++
	}
	file.Close()
}
