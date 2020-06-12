package trisoban

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	tl "github.com/JoelOtter/termloop"
)

var game *tl.Game
var gs *Gamescreen

// CurrentLevel this integer will determine the level you are currently on.
var CurrentLevel int

// TotalLevels this integer represents the total amount of levels.
var TotalLevels = GetTotalLevels()

// StartGame starts the game by creating a new game and adding a titlescreen.
func StartGame() {
	game = tl.NewGame()

	ts := NewTitleScreen()

	game.Screen().SetLevel(ts)
	game.Screen().SetFps(40)
	game.Start()
}

// GetTotalLevels get total levels in lvl folder
func GetTotalLevels() int {
	lvlFiles, _ := ioutil.ReadDir("data/lvl")
	return len(lvlFiles)
}

//LoadLevel load level from last saved lvl
func LoadLevel() int {
	dat, err := ioutil.ReadFile("data/ui/loadgame.txt")
	if err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}
	byteToInt, _ := strconv.Atoi(string(dat))
	return byteToInt
}

// AddGameEntities create all entities on the game screen and map level
func (gs *Gamescreen) AddGameEntities() {
	col = NewEntityCollection()

	gs.MapLevel()

	for _, v := range col.Goals {
		gs.AddEntity(v)
	}

	for _, v := range col.Crates {
		gs.AddEntity(v)
	}

	gs.AddEntity(col.Border)
	gs.AddEntity(col.Player)
}

// MapLevel reads from a level file and maps the coordinates for all of the entities
func (gs *Gamescreen) MapLevel() {
	var startx = 7 // Determines at which coordinate the rendering starts
	var starty = 6 // Determines at which coordinate the rendering starts

	col.Border.bCoords = make(map[Coordinates]int)

	file, err := os.Open(fmt.Sprintf("data/lvl/level%d.txt", CurrentLevel))
	if err != nil {
		log.Fatalf("Level with number %d was not found", CurrentLevel)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xline := strings.Split(scanner.Text(), "")
		for i, v := range xline {
			switch v {
			case "#":
				col.Border.bCoords[Coordinates{
					X: startx + i,
					Y: starty,
				}] = 1
			case "P":
				col.Player.Coordinates = Coordinates{
					X: startx + i,
					Y: starty,
				}
				col.Player.SetPosition(col.Player.X, col.Player.Y)
			case "G":
				g := NewGoal()
				g.Coordinates = Coordinates{
					X: startx + i,
					Y: starty,
				}

				g.SetPosition(g.X, g.Y)
				col.Goals = append(col.Goals, g)
			case "C":
				c := NewCrate()
				c.Coordinates = Coordinates{
					X: startx + i,
					Y: starty,
				}

				c.SetPosition(c.X, c.Y)
				col.Crates = append(col.Crates, c)
			}
		}
		starty++
	}
	file.Close()
}

// RemoveGameEntities delete all entities on the game screen.
func (gs *Gamescreen) RemoveGameEntities() {
	gs.RemoveEntity(col.Border)
	gs.RemoveEntity(col.Player)

	for _, v := range col.Goals {
		gs.RemoveEntity(v)
	}

	for _, v := range col.Crates {
		gs.RemoveEntity(v)
	}
}

// Move calculate and check collisions for the current move
func (gs *Gamescreen) Move(dir string) {
	for ci, cv := range col.Crates {
		if col.Player.CheckCrateCollision(cv, dir) {
			col.Player.SetPosition(col.Player.X, col.Player.Y)
			cv.MoveCrate(dir, false)
			crateslicecopy := cv.CopyCrateSlice()
			if cv.CheckBorderCollision() {
				col.Player.MovePlayer(dir, true)
				cv.MoveCrate(dir, true)
			}
			if cv.CheckCrateCollision(cv.RemoveCrate(crateslicecopy, ci)) {
				cv.MoveCrate(dir, true)
				col.Player.MovePlayer(dir, true)
			}
			for _, gv := range col.Goals {
				if cv.CheckGoalCollision(gv) {
					if cv.CheckActivatedGoalCollision(gv) {
						col.Player.MovePlayer(dir, true)
						cv.MoveCrate(dir, true)
					} else {
						gv.isActivated = true
						gs.RemoveEntity(cv)
						col.Crates = cv.RemoveCrate(col.Crates, ci)
						_ = gs.CheckLevelCompletion()
					}
				}
			}
		}
	}
	if col.Player.CheckBorderCollision(dir) {
		col.Player.MovePlayer(dir, true)
	}
	for _, gv := range col.Goals {
		if col.Player.CheckActivatedGoalCollision(gv, dir) {
			col.Player.MovePlayer(dir, true)
		}
	}
	col.Player.MovePlayer(dir, false)
}

// CheckLevelCompletion will remove the crate entity and print a text when the level is beaten.
func (gs *Gamescreen) CheckLevelCompletion() bool {
	if len(col.Crates) == 0 {
		gs.LevelCompleted.SetText("You have completed this level")
		if sw != nil {
			gs.Time.SetText("Time: " + sw.CheckpointTime())
		}
		gs.Instructions[0].SetColor(tl.ColorGreen, tl.ColorBlack)
		gs.Instructions[1].SetColor(tl.ColorRed, tl.ColorBlack)

		return true
	}
	return false
}

// ChangeLevel will change the level given the player input.
func (gs *Gamescreen) ChangeLevel(selection string) {
	switch selection {
	case "next":
		CurrentLevel++
	case "previous":
		CurrentLevel--
	}

	if CurrentLevel < TotalLevels {
		gs.RemoveGameEntities()
		gs.UpdateText()
		gs.AddGameEntities()
	} else {
		gcs := NewGameCompletionScreen()
		game.Screen().SetLevel(gcs)
	}
}

// SaveGame save your current level to a file and show a confirmation.
func (gs *Gamescreen) SaveGame() {
	gs.LevelCompleted.SetText("Your game has been saved")
	ioutil.WriteFile("data/ui/loadgame.txt", []byte(strconv.Itoa(CurrentLevel)), 0644)
}

// SaveTime save your collective time after beating the game.
func (gs *GameCompletionScreen) SaveTime(finalTime string) {
	var newRow []byte
	datetime := time.Now()
	newRow = []byte(fmt.Sprintf("\n|" + fmt.Sprintf("%s", datetime.Format("01-02-2006 15:04:05")) + "|" + fmt.Sprintf("%s", finalTime) + "|" + fmt.Sprintf("%d", TotalLevels) + "|"))
	f, err := os.OpenFile("data/HIGHSCORES.md", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}

	_, err2 := f.Write(newRow)
	if err2 != nil {
		log.Fatalf("Error writing to file: %s", err2)
	}

	f.Close()
}

// RestartLevel is a function that will reset the level by setting all of the entites to their default positions.
func (gs *Gamescreen) RestartLevel() {
	gs.RemoveGameEntities()
	gs.AddGameEntities()
}

// UpdateText will update the text on the game screen. This function is called when the level is changed
func (gs *Gamescreen) UpdateText() {
	gs.CurrentLevelText.SetText(fmt.Sprintf("---| Current Level %d of %d |---", CurrentLevel, TotalLevels))
	gs.LevelCompleted.SetText("")
	gs.Time.SetText("")
	gs.Instructions[0].SetColor(tl.ColorWhite, tl.ColorBlack)
	gs.Instructions[1].SetColor(tl.ColorWhite, tl.ColorBlack)
}
