package trisoban

import (
	"bufio"
	"log"
	"os"
	"strings"

	tl "github.com/JoelOtter/termloop"
)

func (level *Level) Draw(screen *tl.Screen) {
	file, err := os.Open("util/levels.txt")
	if err != nil {
		log.Fatalln(file)
	}
	scanner := bufio.NewScanner(file)
	var startx = 7 // Determines at which coordinate the rendering starts
	var starty = 5 // Determines at which coordinate the rendering starts
Loop:
	for scanner.Scan() {
		xline := strings.Split(scanner.Text(), "")
		for i, v := range xline {
			switch v {
			case "#":
				screen.RenderCell(startx+i+1, starty, &tl.Cell{
					Fg: tl.ColorBlue,
					Ch: '▓',
				})
			case "C":
				screen.RenderCell(startx+i+1, starty, &tl.Cell{
					Fg: tl.ColorRed,
					Ch: '▒',
				})
			case "P":
				screen.RenderCell(startx+i+1, starty, &tl.Cell{
					Bg: tl.ColorYellow,
					Fg: tl.ColorWhite,
					Ch: '▓',
				})
			case "G":
				screen.RenderCell(startx+i+1, starty, &tl.Cell{
					Fg: tl.ColorYellow,
					Ch: '░',
				})
			case "-":
				break Loop
			}
		}
		starty++
	}
}
