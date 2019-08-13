package main

import (
	"github.com/gen2brain/raylib-go/raylib"
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

	offsetX := 3
	// @todo / 10 should be / tileHeight for the tile grid this is working on...
	offsetY := (rl.GetScreenHeight() / 10) - 8

	newGameBtn := owner.Screen.Set(NewButton("NewGameBtn", "Play", 16, 2, 2, offsetX, offsetY, BtnTextCenter, ui.DefaultBorderColour, func() {
		s.Owner.ChangeState(NewWorld())
	}), 99)

	offsetX += int(3 + newGameBtn.GetWidth())

	continueGameBtn := owner.Screen.Set(NewButton("MorgueBtn", "Morgue", 16, 2, 2, offsetX, offsetY, BtnTextCenter, ui.DefaultBorderColour, func() {
		s.Owner.ChangeState(NewMorgueState())
	}), 99)

	offsetX += int(3 + continueGameBtn.GetWidth())

	settingsBtn := owner.Screen.Set(NewButton("SettingsBtn", "Settings", 16, 2, 2, offsetX, offsetY, BtnTextCenter, ui.DefaultBorderColour, func() {
		s.Owner.ChangeState(NewSettingsState())
	}), 99)

	offsetX += int(3 + settingsBtn.GetWidth())

	owner.Screen.Set(NewButton("HelpBtn", "?", 0, 2, 2, offsetX, offsetY, BtnTextCenter, ui.BorderColour{"AnsiRed", "AnsiGreen", "AnsiYellow", "AnsiBlue"}, func() {
		s.Owner.ChangeState(NewHelpState())
	}), 99)

	// helpBtn.SetBorderColour(ui.BorderColour{"AnsiRed", "AnsiGreen", "AnsiYellow", "AnsiBlue"})

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
