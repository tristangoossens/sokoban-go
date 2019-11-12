package trisoban

import (
	"fmt"
	"io/ioutil"

	tl "github.com/JoelOtter/termloop"
)

func StartGame() {
	game = tl.NewGame()

	ts := NewTitleScreen()

	game.Screen().SetLevel(ts)
	game.Screen().SetFps(40)
	game.Start()
}

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

func NewTitleScreen() *Titlescreen {
	ts := new(Titlescreen)
	ts.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	logofile, _ := ioutil.ReadFile("util/logo.text")
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

func NewGameScreen() *Gamescreen {
	gs := new(Gamescreen)
	gs.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	gs.CurrentLevel = tl.NewText(69, 3, fmt.Sprintf("---| Current Level %d of %d |---", CurrentLevel, TotalLevels), tl.ColorWhite, tl.ColorBlack)

	instructionsFile, _ := ioutil.ReadFile("util/instructions.txt")
	instructionsEntity := tl.NewEntityFromCanvas(55, 5, tl.CanvasFromString(string(instructionsFile)))

	border = NewBorder()
	crate = NewCrate()
	goal = NewGoal()
	player = NewPlayer()

	MapLevel()

	gs.AddEntity(gs.CurrentLevel)
	gs.AddEntity(instructionsEntity)
	gs.AddEntity(border)
	gs.AddEntity(goal)
	gs.AddEntity(crate)
	gs.AddEntity(player)

	return gs
}

func UpdateLevelText() {
	gs.CurrentLevel.SetText(fmt.Sprintf("---| Current Level %d of %d |---", CurrentLevel, TotalLevels))
}

func LevelCompleted() {
	crate.reachedGoal = true
	gs.BeatLevel = tl.NewText(7, 20, "Congratulations!, You have beaten this level!", tl.ColorGreen, tl.ColorBlack)
	gs.RemoveEntity(crate)
	gs.AddEntity(gs.BeatLevel)

}

func NextLevel() {
	crate.reachedGoal = false
	CurrentLevel += 1
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
