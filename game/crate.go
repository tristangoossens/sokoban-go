package trisoban

import (
	"bufio"
	"log"
	"os"
	"strings"

	tl "github.com/JoelOtter/termloop"
)

func MapCrate() Coordinates {
	var startx = 7 // Determines at which coordinate the rendering starts
	var starty = 5 // Determines at which coordinate the rendering starts

	var coords Coordinates

	file, err := os.Open("util/levels.txt")
	if err != nil {
		log.Fatalln(file)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		xline := strings.Split(scanner.Text(), "")
		for i, v := range xline {
			switch v {
			case "C":
				coords = Coordinates{
					X: startx + i,
					Y: starty,
				}
			}
		}
		starty++
	}
	file.Close()
	return coords
}

func MapGoal() Coordinates {
	var startx = 7 // Determines at which coordinate the rendering starts
	var starty = 5 // Determines at which coordinate the rendering starts

	var coords Coordinates

	file, err := os.Open("util/levels.txt")
	if err != nil {
		log.Fatalln(file)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		xline := strings.Split(scanner.Text(), "")
		for i, v := range xline {
			switch v {
			case "G":
				coords = Coordinates{
					X: startx + i,
					Y: starty,
				}
			}
		}
		starty++
	}
	file.Close()
	return coords
}

func (crate *Crate) CheckBorderCollision(x, y int) bool {
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

func (crate *Crate) CheckGoalCollision(x, y int) bool {
	c := Coordinates{
		X: x,
		Y: y,
	}
	if goal.Coordinates == c {
		return true
	}
	return false
}

func (goal *Goal) Draw(screen *tl.Screen) {
	if crate.reachedGoal {
		screen.RenderCell(goal.X, goal.Y, &tl.Cell{
			Fg: tl.ColorGreen,
			Ch: '▓',
		})
	} else {
		screen.RenderCell(goal.X, goal.Y, &tl.Cell{
			Fg: tl.ColorYellow,
			Ch: '░',
		})
	}
}

func (crate *Crate) Draw(screen *tl.Screen) {
	screen.RenderCell(crate.X, crate.Y, &tl.Cell{
		Fg: tl.ColorMagenta,
		Ch: '▓',
	})
}
