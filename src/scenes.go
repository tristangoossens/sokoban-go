package trisoban

import (
	"fmt"
	"io/ioutil"

	tl "github.com/JoelOtter/termloop"
)

// Titlescreen is the level for the titlescreen.
type Titlescreen struct {
	tl.Level
	Mainmenu   *tl.Entity
	EntityText []*tl.Text
}

// Gamescreen is the level for the gamescreen.
type Gamescreen struct {
	tl.Level
	UI               *tl.Entity
	CurrentLevelText *tl.Text
	Time             *tl.Text
	Instructions     []*tl.Text
	SaveConfirmation *tl.Text
	LevelCompleted   *tl.Text
}

// GameCompletionScreen is the level of the victory screen, displaying when you finish all levels.
type GameCompletionScreen struct {
	tl.Level
	UI            *tl.Entity
	FinalTimeText *tl.Text
	FinalTime     string
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

	gs.LevelCompleted = tl.NewText(45, 27, "", tl.ColorWhite, tl.ColorBlack)
	gs.SaveConfirmation = tl.NewText(45, 26, "", tl.ColorWhite, tl.ColorBlack)
	gs.Time = tl.NewText(52, 4, "", tl.ColorWhite, tl.ColorBlack)

	GetTotalLevels()
	gs.CurrentLevelText = tl.NewText(44, 2, fmt.Sprintf("---| Current Level %d of %d |---", CurrentLevel, TotalLevels), tl.ColorWhite, tl.ColorBlack)

	uiFile, _ := ioutil.ReadFile("data/ui/gameui.txt")
	gs.UI = tl.NewEntityFromCanvas(0, 0, tl.CanvasFromString(string(uiFile)))

	gs.Instructions = []*tl.Text{
		tl.NewText(11, 29, "F1: Next level", tl.ColorWhite, tl.ColorBlack),
		tl.NewText(37, 29, "F2: Previous level", tl.ColorWhite, tl.ColorBlack),
		tl.NewText(96, 29, "Ins: Save game", tl.ColorWhite, tl.ColorBlack),
	}

	gs.AddEntity(gs.UI)
	gs.AddEntity(gs.CurrentLevelText)
	gs.AddEntity(gs.Time)
	gs.AddEntity(gs.LevelCompleted)
	gs.AddEntity(gs.SaveConfirmation)

	for _, v := range gs.Instructions {
		gs.AddEntity(v)
	}

	gs.AddGameEntities()

	return gs
}

// NewGameCompletionScreen will create a new gamecompletionscreen for when all levels are beaten.
func NewGameCompletionScreen() *GameCompletionScreen {
	gcs := new(GameCompletionScreen)
	gcs.Level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})

	UIFile, _ := ioutil.ReadFile("data/ui/victoryui.txt")
	gcs.UI = tl.NewEntityFromCanvas(0, 0, tl.CanvasFromString(string(UIFile)))

	gcs.AddEntity(gcs.UI)

	if sw != nil {
		sw.Stop()
		gcs.FinalTimeText = tl.NewText(48, 14, "Final Time: "+sw.FinishTime(), tl.ColorWhite, tl.ColorBlack)
		gcs.FinalTime = sw.FinishTime()
		gcs.AddEntity(gcs.FinalTimeText)
	}

	return gcs
}
