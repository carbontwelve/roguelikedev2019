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
	e.screen.Reset() // @todo move this to a on state change function as we may not want to reset on World construction...
	e.screen.Set(ui.NewComponent("Viewport", position.DungeonWidth, position.DungeonHeight, 0, 0, true), 10)

	e.screen.Get("Viewport").SetAutoClear(false)

	s := &LobbyState{
		State: state.State{e: e, Quit: false},
	}

	return s
}

func (s LobbyState) Draw(dt float32) {
	rl.ClearBackground(ui.GameColours["Bg"])
}

func (s *LobbyState) Update(dt float32) {
	if rl.IsKeyPressed(rl.KeySpace) {
		s.e.ChangeState(NewWorld(s.e))
	}
}

func (s LobbyState) GetName() string {
	return "Lobby"
}
