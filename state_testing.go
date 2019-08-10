package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"raylibtinkering/state"
	"raylibtinkering/ui"
)

type TestingState struct {
	lastKeyCode, x, y int
	state.State
}

func NewTestingState() *TestingState {
	s := &TestingState{
		State:       state.State{},
		x:           0,
		y:           0,
		lastKeyCode: 0,
	}

	return s
}

func (s TestingState) Draw(dt float32) {
	rl.ClearBackground(ui.GameColours["Bg"])

	// test sprite sheet whole
	for y := uint(0); y < s.e.screen.Rows; y++ {
		for x := uint(0); x < s.e.screen.Cols; x++ {
			//s.e.sprites.At(x, y).Draw(rl.NewVector2(float32(10+(10*x)), float32(50+(10*y))), GameColours["Player"])
		}
	}

	//s.e.font.Draw(s.lastKeyCode, rl.NewVector2(10, 200), GameColours["Player"])
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
