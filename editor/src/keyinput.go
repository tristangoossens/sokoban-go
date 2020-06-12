package trisobaneditor

import (
	tl "github.com/JoelOtter/termloop"
)

// Tick listen for keyinput on the editor screen
func (s *Scene) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowUp:
			s.c.Move("up")
		case tl.KeyArrowDown:
			s.c.Move("down")
		case tl.KeyArrowLeft:
			s.c.Move("left")
		case tl.KeyArrowRight:
			s.c.Move("right")
		case tl.KeyF1:
			s.c.DrawPlayer()
			s.savetext.SetColor(tl.ColorWhite, tl.ColorBlack)
		case tl.KeyF2:
			s.c.DrawCrate()
			s.savetext.SetColor(tl.ColorWhite, tl.ColorBlack)
		case tl.KeyF3:
			s.c.DrawGoal()
			s.savetext.SetColor(tl.ColorWhite, tl.ColorBlack)
		case tl.KeySpace:
			s.c.DrawBorder()
			s.savetext.SetColor(tl.ColorWhite, tl.ColorBlack)
		case tl.KeyInsert:
			s.df.SaveLevel()
		case tl.KeyDelete:
			s.c.DeleteEntity(s.c.Coordinates)
			s.savetext.SetColor(tl.ColorWhite, tl.ColorBlack)
		case tl.KeyEnter:
			sc = NewScene()
			g.Screen().SetLevel(sc)
		}
	}
}
