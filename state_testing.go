package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TestingState struct {
	lastKeyCode, x, y int
	State
}

func NewTestingState(e *Engine) *TestingState {
	s := &TestingState{
		State:       State{e: e},
		x:           0,
		y:           0,
		lastKeyCode: 0,
	}

	return s
}

func (s TestingState) Draw(dt float32) {
	rl.ClearBackground(UIBackgroundColour)

	// test sprite sheet whole
	for y := 0; y < s.e.sprites.Rows; y++ {
		for x := 0; x < s.e.sprites.Cols; x++ {
			s.e.sprites.At(x, y).Draw(rl.NewVector2(float32(10+(10*x)), float32(50+(10*y))), PlayerColour)
		}
	}

	s.e.font.Draw(s.lastKeyCode, rl.NewVector2(10, 200), PlayerColour)
}

func (s *TestingState) Update(dt float32) {
	keyCode := int(rl.GetKeyPressed())
	if keyCode > -1 {
		if keyCode != s.lastKeyCode {
			fmt.Println(keyCode)
		}
		s.lastKeyCode = keyCode
	}
}

func (s TestingState) GetName() string {
	return "Testing"
}
