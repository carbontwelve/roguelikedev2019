package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
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

	s.Owner = owner
	return nil
}

func (s *LobbyState) Popped(owner *state.Engine) error {
	return nil
}

func (s LobbyState) Draw(dt float32) {
	rl.ClearBackground(ui.GameColours["Bg"])
}

func (s *LobbyState) Update(dt float32) {
	if rl.IsKeyPressed(rl.KeySpace) {
		s.Owner.ChangeState(NewWorld())
	}
}

func (s LobbyState) GetName() string {
	return "Lobby"
}
