package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"raylibtinkering/position"
	"raylibtinkering/state"
	"raylibtinkering/ui"
)

type MorgueState struct {
	state.State
}

func NewMorgueState() *MorgueState {
	s := &MorgueState{
		State: state.State{Quit: false},
	}

	return s
}

func (s *MorgueState) Pushed(owner *state.Engine) error {
	owner.Screen.Reset() // @todo move this to a on state change function as we may not want to reset on World construction...

	viewport := owner.Screen.Set(ui.NewComponent("Viewport", position.DungeonWidth-2, position.DungeonHeight-2, 1, 1, true), 10)
	viewport.SetAutoClear(false)
	viewport.SetTitle("Morgue")
	viewport.SetBorderStyle(ui.SingleWallBorder)

	smallButtonStyle := DefaultButtonStyle
	smallButtonStyle.normal.borderStyle = ui.ZeroWallBorder
	smallButtonStyle.hover.borderStyle = ui.ZeroWallBorder

	owner.Screen.Set(NewButton("BackBtn", "] <B>ack [", 0, 1, 1, (rl.GetScreenWidth()/10)-14, 0, BtnTextCenter, smallButtonStyle, func() {
		s.Owner.ChangeState(NewLobbyState())
	}), 99)

	s.Owner = owner
	return nil
}

func (s *MorgueState) Popped(owner *state.Engine) error {
	return nil
}

func (s *MorgueState) Tick(dt float32) {
	// NOP
}

func (s MorgueState) GetName() string {
	return "Morgue"
}
