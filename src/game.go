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
