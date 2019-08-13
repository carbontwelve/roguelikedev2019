package ui

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var MousePos *Mouse

type Mouse struct {
	Hover    bool
	X, Y     int
	CellSize uint
}

func (m *Mouse) Update(mousePos rl.Vector2) {
	if mousePos.X > 0 && mousePos.Y > 0 && mousePos.X < float32(rl.GetScreenWidth()) && mousePos.Y < float32(rl.GetScreenHeight()) {
		m.Hover = true
		m.X = int(mousePos.X / float32(m.CellSize))
		m.Y = int(mousePos.Y / float32(m.CellSize))
	} else {
		m.Hover = false
	}
}

// Returns true if mouse is found inside cell area on screen
func (m Mouse) Inside(x1, y1, x2, y2 int) bool {
	return m.Hover == true && m.X >= x1 && m.X <= x2 && m.Y >= y1 && m.Y <= y2
}
