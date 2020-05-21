package trisoban

import tl "github.com/JoelOtter/termloop"

var game *tl.Game
var gs *Gamescreen
var col *EntityCollection

// CurrentLevel this integer will determine the level you are currently on.
var CurrentLevel int

// TotalLevels this integer represents the total amount of levels.
var TotalLevels int

// Titlescreen is the level for the titlescreen, this contains the logo and the instruction text.
type Titlescreen struct {
	tl.Level
	Mainmenu   *tl.Entity
	EntityText []*tl.Text
}

// Gamescreen is the level for the gamescreen, this contains the currentleveltext and the beatlevel text.
type Gamescreen struct {
	tl.Level
	UI               *tl.Entity
	CurrentLevelText *tl.Text
	Instructions     []*tl.Text
	SaveConfirmation *tl.Text
	LevelCompleted   *tl.Text
}

// Entities

//EntityCollection collection of all entities within the game
type EntityCollection struct {
	Player *Player
	Crates []*Crate
	Goals  []*Goal
	Border *Border
}

// Player inherits the entity making it a drawable, it also inherits the border and coordinates struct.
type Player struct {
	*tl.Entity
	Border
	Coordinates
}

// Border inherits the entity making it a drawable. It also consists of the bordercoordinates in a map.
type Border struct {
	*tl.Entity
	bCoords map[Coordinates]int
}

// Crate inherits the entity making it a drawable, it also inherits the coordinates struct, and has a bool for the rachedgoal check.
type Crate struct {
	*tl.Entity
	Coordinates
	reachedGoal bool
}

// Goal inherits the entity making it a drawable, it also inherits the coordinates struct.
type Goal struct {
	*tl.Entity
	Coordinates
	isActivated bool
}

// Coordinates is used to save the X and Y coordinates of multiple entities.
type Coordinates struct {
	X int
	Y int
}
