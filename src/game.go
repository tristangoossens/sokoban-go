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
	gs.RemoveEntity(border)
	gs.RemoveEntity(crate)
	gs.RemoveEntity(goal)
	gs.RemoveEntity(player)
	gs.RemoveEntity(gs.BeatLevel)

	border = NewBorder()
	crate = NewCrate()
	goal = NewGoal()
	player = NewPlayer()

	MapLevel()

	gs.AddEntity(border)
	gs.AddEntity(goal)
	gs.AddEntity(crate)
	gs.AddEntity(player)
}

// NewTitleScreen will create a new titlescreen and read from the logo file to print out the ASCII art logo. This function will return a pointer to titlescreen.
func NewTitleScreen() *Titlescreen {
	ts := new(Titlescreen)
	ts.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	logofile, _ := ioutil.ReadFile("data/ui/logo.text")
	logoEntity := tl.NewEntityFromCanvas(10, 3, tl.CanvasFromString(string(logofile)))

	ts.Text = []*tl.Text{
		tl.NewText(10, 20, "Press ENTER to start!", tl.ColorWhite, tl.ColorBlack),
		tl.NewText(10, 22, "Press BACKSPACE to start from last saved level!", tl.ColorWhite, tl.ColorBlack),
	}

	for _, v := range ts.Text {
		ts.Level.AddEntity(v)
	}

	ts.Level.AddEntity(logoEntity)
	return ts
}

// NewGameScreen will create a new gamescreen with the currentleveltext, instructions and the level itself and return it where it is called.
func NewGameScreen() *Gamescreen {
	gs := new(Gamescreen)
	gs.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	gs.CurrentLevel = tl.NewText(69, 3, fmt.Sprintf("---| Current Level %d of %d |---", CurrentLevel, TotalLevels), tl.ColorWhite, tl.ColorBlack)

	instructionsFile, _ := ioutil.ReadFile("data/ui/instructions.txt")
	instructionsEntity := tl.NewEntityFromCanvas(55, 5, tl.CanvasFromString(string(instructionsFile)))

	col = NewEntityCollection()
	MapLevel()

	gs.AddEntity(gs.CurrentLevel)
	gs.AddEntity(instructionsEntity)
	gs.AddEntity(col.Border)

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
	crate.reachedGoal = true
	gs.BeatLevel = tl.NewText(7, 20, "Congratulations!, You have beaten this level!", tl.ColorGreen, tl.ColorBlack)
	gs.RemoveEntity(crate)
	gs.AddEntity(gs.BeatLevel)

}

// ChangeLevel will change the level given the player input(F1: "NEXT", F2: "PREVIOUS").
func ChangeLevel(s string) {
	crate.reachedGoal = false

	switch s {
	case "NEXT":
		CurrentLevel += 1
	case "PREVIOUS":
		CurrentLevel -= 1
	}

	gs.RemoveEntity(border)
	gs.RemoveEntity(player)
	gs.RemoveEntity(goal)
	gs.RemoveEntity(gs.BeatLevel)

	border = NewBorder()
	crate = NewCrate()
	goal = NewGoal()
	player = NewPlayer()

	MapLevel()

	gs.AddEntity(border)
	gs.AddEntity(goal)
	gs.AddEntity(crate)
	gs.AddEntity(player)

	UpdateLevelText()
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
