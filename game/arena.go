package trisoban

import (
	"bufio"
	"log"
	"os"
	"strings"

	tl "github.com/JoelOtter/termloop"
)

type Level struct {
	*tl.Entity
	BorderCoordinates map[Coordinates]int
}

type Coordinates struct {
	X int
	Y int
}

// Draw draw
func (level *Level) Draw(screen *tl.Screen) {
	file, err := os.Open("levels.csv")
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
			case "C":
				screen.RenderCell(startx+i+1, starty, &tl.Cell{
					Bg: tl.ColorBlack,
					Ch: '©',
				})
			case "P":
				screen.RenderCell(startx+i+1, starty, &tl.Cell{
					Bg: tl.ColorBlack,
					Fg: tl.ColorWhite,
					Ch: '■',
				})
			case "G":
				screen.RenderCell(startx+i+1, starty, &tl.Cell{
					Fg: tl.ColorYellow,
					Ch: '░',
				})
			}
		}
		starty++
	}
}
