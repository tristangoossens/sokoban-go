package trisoban

import (
	"fmt"
	"io/ioutil"

	tl "github.com/JoelOtter/termloop"
)

// NewTitleScreen will create a new titlescreen and read from the logo file to print out the ASCII art logo. This function will return a pointer to titlescreen.
func NewTitleScreen() *Titlescreen {
	ts := new(Titlescreen)
	ts.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	logofile, _ := ioutil.ReadFile("data/ui/mainmenu.text")
	ts.Mainmenu = tl.NewEntityFromCanvas(0, 0, tl.CanvasFromString(string(logofile)))

	ts.EntityText = []*tl.Text{
		tl.NewText(69, 21, "▓", tl.ColorRed, tl.ColorDefault),
		tl.NewText(69, 22, "▓", tl.ColorBlue, tl.ColorDefault),
		tl.NewText(69, 23, "▓", tl.ColorWhite, tl.ColorDefault),
		tl.NewText(69, 24, "░", tl.ColorYellow, tl.ColorDefault),
		tl.NewText(69, 25, "▓", tl.ColorGreen, tl.ColorDefault),
	}

	for _, v := range ts.EntityText {
		ts.AddEntity(v)
	}

	ts.AddEntity(ts.Mainmenu)
	return ts
}

// NewGameScreen will create a new gamescreen with the currentleveltext, instructions and the level itself and return it where it is called.
func NewGameScreen() *Gamescreen {
	gs := new(Gamescreen)
	gs.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	gs.LevelCompleted = tl.NewText(45, 27, "", tl.ColorWhite, tl.ColorBlack)

	lvlFiles, _ := ioutil.ReadDir("data/lvl")
	TotalLevels = len(lvlFiles)
	gs.CurrentLevelText = tl.NewText(44, 2, fmt.Sprintf("---| Current Level %d of %d |---", CurrentLevel, TotalLevels), tl.ColorWhite, tl.ColorBlack)

	uiFile, _ := ioutil.ReadFile("data/ui/gameui.txt")
	gs.UI = tl.NewEntityFromCanvas(0, 0, tl.CanvasFromString(string(uiFile)))

	gs.Instructions = []*tl.Text{
		tl.NewText(11, 29, "F1: Next level", tl.ColorWhite, tl.ColorBlack),
		tl.NewText(37, 29, "F2: Previous level", tl.ColorWhite, tl.ColorBlack),
	}

	gs.AddEntity(gs.UI)
	gs.AddEntity(gs.CurrentLevelText)
	gs.AddEntity(gs.LevelCompleted)

	for _, v := range gs.Instructions {
		gs.AddEntity(v)
	}

	gs.AddGameEntities()

	return gs
}
