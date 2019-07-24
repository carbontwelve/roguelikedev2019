package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type LobbyState struct {
	State
}

func NewLobbyState(e *Engine) *LobbyState {
	s := &LobbyState{
		State: State{e: e},
	}

	return s
}

func (s LobbyState) Draw(dt float32) {
	rl.ClearBackground(UIBackgroundColour)
}

func (s *LobbyState) Update(dt float32) {
	if rl.IsKeyPressed(rl.KeySpace) {
		s.e.ChangeState(NewWorld(s.e))
	}
}

func (s LobbyState) GetName() string {
	return "Lobby"
}
