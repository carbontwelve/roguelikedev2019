package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type LobbyState struct {
	State
}

func NewLobbyState(g *Engine) *LobbyState {
	s := &LobbyState{
		State: State{g},
	}

	return s
}

func (s LobbyState) Draw(dt float32) {
	// ...
}

func (s *LobbyState) Update(dt float32) {
	if rl.IsKeyPressed(rl.KeySpace) {
		s.g.ChangeState(NewMainState(s.g))
	}
}

func (s LobbyState) GetName() string {
	return "Lobby"
}
