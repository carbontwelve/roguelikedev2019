package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainState struct {
	lastKeyCode, x, y int
	State
}

func NewMainState(g *Engine) *MainState {
	s := &MainState{
		State:       State{g},
		x:           0,
		y:           0,
		lastKeyCode: 0,
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

	s.g.font.Draw(s.lastKeyCode, rl.NewVector2(10, 200), rl.Green)
}

func (s *MainState) Update(dt float32) {
	keyCode := int(rl.GetKeyPressed())
	if keyCode > -1 {
		if keyCode != s.lastKeyCode {
			fmt.Println(keyCode)
		}
		s.lastKeyCode = keyCode
	}
}

func (s MainState) GetName() string {
	return "Main"
}
