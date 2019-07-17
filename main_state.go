package main

import rl "github.com/gen2brain/raylib-go/raylib"

type MainState struct {
	x, y int
	State
}

const (
	spriteSize = 48
)

func NewMainState(g *Engine) *MainState {
	s := &MainState{
		State: State{g},
		x:     0,
		y:     0,
	}

	return s
}

func (s MainState) Draw(dt float32) {
	rl.ClearBackground(rl.Yellow)

	// test sprite sheet whole
	for y := 0; y < s.g.sprites.Rows; y++ {
		for x := 0; x < s.g.sprites.Cols; x++ {
			s.g.sprites.At(x, y).Draw(rl.NewVector2(float32(10+(10*x)), float32(50+(10*y))), rl.Green)
		}
	}

	s.g.sprites.At(s.x, s.y).Draw(rl.NewVector2(10, 200), rl.Green)

}

func (s *MainState) Update(dt float32) {
	if rl.IsKeyPressed(rl.KeySpace) {
		s.g.ChangeState(NewLobbyState(s.g))
	} else if rl.IsKeyPressed(rl.KeyUp) {
		s.x++
		if s.x > s.g.sprites.Cols {
			s.x = 0
			s.y++
		}
		if s.y > s.g.sprites.Rows {
			s.x = 0
			s.y = 0
		}
	}
}

func (s MainState) GetName() string {
	return "Main"
}
