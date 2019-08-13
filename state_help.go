package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"raylibtinkering/state"
	"raylibtinkering/ui"
)

type HelpState struct {
	state.State
}

func NewHelpState() *HelpState {
	s := &HelpState{
		State: state.State{Quit: false},
	}

	return s
}

func (s *HelpState) Pushed(owner *state.Engine) error {
	owner.Screen.Reset() // @todo move this to a on state change function as we may not want to reset on World construction...

	offsetX := 3
	// @todo / 10 should be / tileHeight for the tile grid this is working on...
	offsetY := (rl.GetScreenHeight() / 10) - 8

	owner.Screen.Set(NewButton("BackBtn", "Back", 16, 2, 2, offsetX, offsetY, BtnTextCenter, ui.DefaultBorderColour, func() {
		s.Owner.ChangeState(NewLobbyState())
	}), 99)

	s.Owner = owner
	return nil
}

func (s *HelpState) Popped(owner *state.Engine) error {
	return nil
}

func (s *HelpState) Tick(dt float32) {
	// NOP
}

func (s HelpState) GetName() string {
	return "Help"
}
