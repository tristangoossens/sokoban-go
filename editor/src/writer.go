package trisobaneditor

import (
	"os"
	"strconv"

	tl "github.com/JoelOtter/termloop"
)

//FieldMap entity for the game field
type FieldMap struct {
	*tl.Entity
	fm map[Coordinates]string

	width, height  int
	startx, starty int
}

//MapDrawingField map the current drawing field.
func MapDrawingField() *FieldMap {
	fm := new(FieldMap)

	fm.startx = 8
	fm.starty = 7

	fm.width = 101
	fm.height = 17

	fm.fm = make(map[Coordinates]string)

	for h := 0; h <= fm.height; h++ {
		for w := 0; w < fm.width; w++ {
			fm.fm[Coordinates{
				x: fm.startx + w,
				y: h + fm.starty,
			}] = ""
		}
	}
	return fm
}

// SaveLevel save your level to a textfile
func (fm *FieldMap) SaveLevel() {

	r1, _ := sc.requirements[0].Color()
	r2, _ := sc.requirements[1].Color()

	if r1 == tl.ColorGreen && r2 == tl.ColorGreen {
		grid := [][]string{}
		var o []byte

		for x := 0; x <= fm.height; x++ {
			row := make([]string, fm.width)
			for y := 0; y < fm.width; y++ {
				row[y] = ""
			}
			grid = append(grid, row)
		}

		for i, v := range fm.fm {
			grid[(i.y - fm.starty)][(i.x - fm.startx)] = v
		}

		f, err := os.Create("../data/lvl/level" + strconv.Itoa(GetTotalLevels()+1) + ".txt")
		if err != nil {
			panic(err)
		}

		for _, v := range grid {
			writeRow := []byte{}
			for _, vv := range v {
				switch vv {
				case "Border":
					o = []byte("#")
				case "Player":
					o = []byte("P")
				case "Crate":
					o = []byte("C")
				case "Goal":
					o = []byte("G")
				case "":
					o = []byte(" ")
				}
				for _, vvv := range o {
					writeRow = append(writeRow, vvv)
				}
			}
			writeRow = append(writeRow, 0x0A)

			_, err := f.Write(writeRow)
			if err != nil {
				panic(err)
			}

			sc.savetext.SetColor(tl.ColorGreen, tl.ColorBlack)
		}
	}
}

// CheckDuplicateEntity check duplicate entities
func (fm *FieldMap) CheckDuplicateEntity(c Coordinates) bool {
	v, _ := fm.fm[c]

	if v != "" {
		return true
	}
	return false
}

// Draw is a function that will draw the border
func (fm *FieldMap) Draw(screen *tl.Screen) {
	for v := range fm.fm {
		screen.RenderCell(v.x, v.y, &tl.Cell{
			Bg: tl.ColorBlack,
		})
	}
}
