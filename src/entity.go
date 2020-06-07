package trisoban

var col *EntityCollection

//EntityCollection collection of all entities within the game.
type EntityCollection struct {
	Player *Player
	Crates []*Crate
	Goals  []*Goal
	Border *Border
}

// Coordinates is used to save the X and Y coordinates of multiple entities.
type Coordinates struct {
	X int
	Y int
}

// NewEntityCollection creates a collection of all the entities in the game
func NewEntityCollection() *EntityCollection {
	col := new(EntityCollection)
	col.Crates = []*Crate{}
	col.Goals = []*Goal{}
	col.Player = NewPlayer()
	col.Border = NewBorder()

	return col
}
