package trisoban

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	tl "github.com/JoelOtter/termloop"
)

// StartGame starts the game by creating a new game and adding a titlescreen.
func StartGame() {
	game = tl.NewGame()

	ts := NewTitleScreen()

	game.Screen().SetLevel(ts)
	game.Screen().SetFps(40)
	game.Start()
}

// RestartLevel is a function that will reset the level by setting all of the entites to their default positions.
func RestartLevel() {

}

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

	gs.CurrentLevel = tl.NewText(44, 2, fmt.Sprintf("---| Current Level %d of %d |---", CurrentLevel, TotalLevels), tl.ColorWhite, tl.ColorBlack)

	uiFile, _ := ioutil.ReadFile("data/ui/gameui.txt")
	gs.UI = tl.NewEntityFromCanvas(0, 0, tl.CanvasFromString(string(uiFile)))

	gs.Instructions = []*tl.Text{
		tl.NewText(11, 29, "F1: Next level", tl.ColorWhite, tl.ColorBlack),
		tl.NewText(37, 29, "F2: Previous level", tl.ColorWhite, tl.ColorBlack),
	}

	col = NewEntityCollection()
	MapLevel()
	gs.AddEntity(gs.UI)
	gs.AddEntity(gs.CurrentLevel)
	gs.AddEntity(col.Border)

	for _, v := range gs.Instructions {
		gs.AddEntity(v)
	}

	for _, v := range col.Goals {
		gs.AddEntity(v)
	}

	for _, v := range col.Crates {
		gs.AddEntity(v)
	}

	gs.AddEntity(col.Player)

	return gs
}

// UpdateLevelText will update the currentlevel text when the ChangeLevel function is called.
func UpdateLevelText() {
	gs.CurrentLevel.SetText(fmt.Sprintf("---| Current Level %d of %d |---", CurrentLevel, TotalLevels))
}

// LevelCompleted will remove the crate entity and print a text when the level is beaten.
func LevelCompleted() {

}

// ChangeLevel will change the level given the player input(F1: "NEXT", F2: "PREVIOUS").
func ChangeLevel(s string) {

}

//LoadLevel load level from last saved lvl
func LoadLevel() int {
	dat, err := ioutil.ReadFile("util/loadgame.txt")
	if err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}
	byteToInt, _ := strconv.Atoi(string(dat))
	return byteToInt
}
