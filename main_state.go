package main

import rl "github.com/gen2brain/raylib-go/raylib"

type MainState struct {
	State
}

const (
	spriteSize = 48
)

func NewMainState(g *Game) *MainState {
	s := &MainState{
		State: State{g},
	}

	return s
}

func (s MainState) Draw(dt float32) {
	rl.ClearBackground(rl.Yellow)
}

func (s *MainState) Update(dt float32) {
	if rl.IsKeyPressed(rl.KeySpace) {
		s.g.ChangeState(NewLobbyState(s.g))
	}
}

func (s MainState) GetName() string {
	return "Main"
}
