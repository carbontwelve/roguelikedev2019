package ui

import "github.com/gen2brain/raylib-go/raylib"

var MousePos *Mouse

type Mouse struct {
	Hover bool
	X, Y  int
}

func (m *Mouse) Update(mousePos rl.Vector2) {
	if mousePos.X > 0 && mousePos.Y > 0 && mousePos.X < float32(rl.GetScreenWidth()) && mousePos.Y < float32(rl.GetScreenHeight()) {
		m.Hover = true
		m.X = int(mousePos.X / 10) // divided by cell size... this needs refactoring
		m.Y = int(mousePos.Y / 10)
	} else {
		m.Hover = false
	}
}
