package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"raylibtinkering/position"
	"raylibtinkering/state"
	"raylibtinkering/ui"
	"unicode/utf8"
)

type DebugState struct {
	state.State
}

func NewDebugState() *DebugState {

	s := &DebugState{
		State: state.State{Quit: false},
	}

	s.DrawColourSquares()

	return s
}

func (d *DebugState) Pushed(owner *state.Engine) error {
	d.Owner.Screen.Reset() // @todo move this to a on state change function as we may not want to reset on World construction...

	d.Owner.Screen.Set(ui.NewComponent("Viewport", position.DungeonWidth, position.DungeonHeight, 0, 0, true), 10)
	d.Owner.Screen.Get("Viewport").SetAutoClear(false)

	d.Owner = owner
	return nil
}

func (d *DebugState) Popped(owner *state.Engine) error {
	return nil
}

func (d DebugState) Draw(dt float32) {
	rl.ClearBackground(ui.GameColours["Bg"])
}

func (d DebugState) DrawColourSquares() {

	colours := ui.LoadedThemeRepository.GetCurrentTheme().AsRaylibColor()

	order := [16]uint{
		7, 15,
		8, 0,
		4, 12,
		2, 10,
		6, 14,
		1, 9,
		5, 13,
		3, 11,
	}

	square := []uint{
		ui.TCOD_CHAR_DNW, ui.TCOD_CHAR_DHLINE, ui.TCOD_CHAR_DHLINE, ui.TCOD_CHAR_DHLINE, ui.TCOD_CHAR_DNE,
		ui.TCOD_CHAR_DVLINE, ui.TCOD_CHAR_BLOCK1, ui.TCOD_CHAR_BLOCK2, ui.TCOD_CHAR_BLOCK3, ui.TCOD_CHAR_DVLINE,
		ui.TCOD_CHAR_DVLINE, ui.TCOD_CHAR_BLOCK1, ui.TCOD_CHAR_BLOCK2, ui.TCOD_CHAR_BLOCK3, ui.TCOD_CHAR_DVLINE,
		ui.TCOD_CHAR_DVLINE, ui.TCOD_CHAR_BLOCK1, ui.TCOD_CHAR_BLOCK2, ui.TCOD_CHAR_BLOCK3, ui.TCOD_CHAR_DVLINE,
		ui.TCOD_CHAR_DSW, ui.TCOD_CHAR_DHLINE, ui.TCOD_CHAR_DHLINE, ui.TCOD_CHAR_DHLINE, ui.TCOD_CHAR_DSE,
	}

	viewport := d.Owner.Screen.Get("Viewport")

	xOff := 12
	yOff := 12

	x := 0

	for _, c := range order {
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				viewport.SetChar(square[y*5+x], position.Position{xOff + x, yOff + y}, colours[c], ui.ColourNC)
			}
		}

		xOff += 7
		x++

		if x >= 8 {
			x = 0
			yOff += 16
			xOff = 12
		}
	}

	currentTheme := ui.LoadedThemeRepository.GetCurrentTheme().Name
	viewport.SetString(currentTheme, position.Position{int(viewport.GetInnerWidth()/2) - utf8.RuneCountInString(currentTheme)/2, int(viewport.GetInnerHeight() / 2)}, ui.GameColours["Fg"], ui.ColourNC)
}

func (d *DebugState) Update(dt float32) {
	if rl.IsKeyPressed(rl.KeySpace) {
		d.Owner.ChangeState(NewWorld())
	}

	if rl.IsKeyPressed(rl.KeyO) || rl.IsKeyPressed(rl.KeyP) {
		d.Owner.Screen.Get("Viewport").Clear()
		d.DrawColourSquares()
	}
}

func (d DebugState) GetName() string {
	return "Debug"
}
