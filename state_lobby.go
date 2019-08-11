package main

import (
	"raylibtinkering/position"
	"raylibtinkering/state"
	"raylibtinkering/ui"
)

type LobbyState struct {
	state.State
}

func NewLobbyState() *LobbyState {
	s := &LobbyState{
		State: state.State{Quit: false},
	}

	return s
}

func (s *LobbyState) Pushed(owner *state.Engine) error {
	owner.Screen.Reset() // @todo move this to a on state change function as we may not want to reset on World construction...

	owner.Screen.Set(ui.NewComponent("Viewport", position.DungeonWidth, position.DungeonHeight, 0, 0, true), 10)
	owner.Screen.Get("Viewport").SetAutoClear(false)

	owner.Screen.Set(NewButton("New Game", 18, 4, 1, 10, func() {
		s.Owner.ChangeState(NewWorld())
	}), 99)

	s.Owner = owner
	return nil
}

func (s *LobbyState) Popped(owner *state.Engine) error {
	return nil
}

func (s *LobbyState) Tick(dt float32) {
	// NOP
}

func (s LobbyState) GetName() string {
	return "Lobby"
}
