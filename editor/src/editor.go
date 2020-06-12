package trisobaneditor

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	tl "github.com/JoelOtter/termloop"
)

var sc *Scene
var g *tl.Game

// Scene sruct that contains the editor level with all its entities.
type Scene struct {
	tl.Level
	ui           *tl.Entity
	currentFile  *tl.Text
	savetext     *tl.Text
	requirements []*tl.Text

	df *FieldMap

	a  *Area
	c  *Cursor
	p  *Player
	b  []*Border
	cs []*Crate
	g  []*Goal
}

// Start the editor tool
func Start() {
	g = tl.NewGame()

	sc = NewScene()

	g.Screen().SetLevel(sc)
	g.Screen().SetFps(40)
	g.Start()
}

// NewScene a new scene for level editor
func NewScene() *Scene {
	sc = new(Scene)
	sc.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	sc.c = CreateCursor()
	sc.a = CreateArea()
	sc.df = MapDrawingField()

	sc.MapLevel()

	sc.currentFile = tl.NewText(7, 5, "Current file: level"+strconv.Itoa(GetTotalLevels()+1)+".txt", tl.ColorWhite, tl.ColorBlack)
	sc.requirements = []*tl.Text{
		tl.NewText(61, 26, "The level contains a player", tl.ColorWhite, tl.ColorBlack),
		tl.NewText(61, 27, "The level has an equal amount of crates and goals", tl.ColorWhite, tl.ColorBlack),
	}

	sc.savetext = tl.NewText(7, 27, "To save your level, press [Ins]", tl.ColorWhite, tl.ColorBlack)

	instructionsFile, _ := ioutil.ReadFile("../data/ui/editorui.txt")
	sc.ui = tl.NewEntityFromCanvas(0, 0, tl.CanvasFromString(string(instructionsFile)))

	sc.AddEntity(sc.ui)

	for _, v := range sc.requirements {
		sc.AddEntity(v)
		v.SetColor(tl.ColorRed, tl.ColorBlack)
	}

	sc.AddEntity(sc.savetext)
	sc.AddEntity(sc.currentFile)

	sc.AddEntity(sc.a)
	sc.AddEntity(sc.df)
	sc.AddEntity(sc.c)

	return sc
}

// GetTotalLevels get the total amount of levels in the game
func GetTotalLevels() int {
	lvlFiles, _ := ioutil.ReadDir("../data/lvl")
	return len(lvlFiles)
}

// MapLevel open the ui and create editor border
func (sc *Scene) MapLevel() {

	var starty = 0

	sc.a.area = make(map[Coordinates]int)

	file, err := os.Open("../data/ui/editorui.txt")
	if err != nil {
		log.Fatalf("ui not found")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xline := strings.Split(scanner.Text(), "")
		for i, v := range xline {
			switch v {
			case "#":
				sc.a.area[Coordinates{
					x: i,
					y: starty,
				}] = 0
			}
		}
		starty++
	}
	file.Close()
}
