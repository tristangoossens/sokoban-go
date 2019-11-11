package trisoban

import (
	"bufio"
	"log"
	"os"
	"strings"

	tl "github.com/JoelOtter/termloop"
)

func (player *Player) Draw(screen *tl.Screen) {
	screen.RenderCell(58, 15, &tl.Cell{
		Fg: tl.ColorRed,
		Ch: '▓',
	})
	screen.RenderCell(58, 16, &tl.Cell{
		Fg: tl.ColorYellow,
		Ch: '░',
	})
	screen.RenderCell(58, 17, &tl.Cell{
		Fg: tl.ColorWhite,
		Ch: '▓',
	})
	screen.RenderCell(58, 18, &tl.Cell{
		Fg: tl.ColorBlue,
		Ch: '▓',
	})
	screen.RenderCell(player.X, player.Y, &tl.Cell{
		Fg: tl.ColorRed,
		Ch: '▓',
	})
}

func GetPlayerPosition() (int, int) {
	var startx = 7 // Determines at which coordinate the rendering starts
	var starty = 5 // Determines at which coordinate the rendering starts
	var x int
	var y int

	file, err := os.Open("util/levels.txt")
	if err != nil {
		log.Fatalln(file)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		xline := strings.Split(scanner.Text(), "")
		for i, v := range xline {
			switch v {
			case "P":
				x, y = startx+i, starty
			}
		}
		starty++
	}
	file.Close()

	return x, y
}

func (player *Player) CheckBorderCollision(x int, y int) bool {
	c := Coordinates{
		X: x,
		Y: y,
	}
	_, exists := border.bCoords[c]

	if exists {
		return true
	}
	return false
}

func (player *Player) CheckCrateCollision(x int, y int) bool {
	c := Coordinates{
		X: x,
		Y: y,
	}
	if crate.Coordinates == c {
		return true
	}
	return false
}
