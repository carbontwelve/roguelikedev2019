package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"raylibtinkering/position"
	"raylibtinkering/ui"
)

type LobbyState struct {
	State
}

func NewLobbyState(e *Engine) *LobbyState {
	e.screen.Reset() // @todo move this to a on state change function as we may not want to reset on World construction...
	e.screen.Set(ui.NewComponent("Viewport", position.DungeonWidth, position.DungeonHeight, 0, 0, true), 10)

	e.screen.Get("Viewport").SetAutoClear(false)

	s := &LobbyState{
		State: State{e: e, Quit: false},
	}

	s.DrawColourSquares()

	return s
}

func (s LobbyState) Draw(dt float32) {
	rl.ClearBackground(ui.ColourBg)
}

func (s LobbyState) DrawColourSquares() {

	order := [16]string{
		"ColourAnsiWhite", "ColourAnsiLightGrey", "ColourAnsiGrey", "ColourAnsiBlack",
		"ColourAnsiBlue", "ColourAnsiLightBlue",
		"ColourAnsiGreen", "ColourAnsiLightGreen",
		"ColourAnsiCyan", "ColourAnsiLightCyan",
		"ColourAnsiRed", "ColourAnsiLightRed",
		"ColourAnsiPurple", "ColourAnsiLightPurple",
		"ColourAnsiYellow", "ColourAnsiLightYellow",
	}

	square := []uint{
		ui.TCOD_CHAR_DNW, ui.TCOD_CHAR_DHLINE, ui.TCOD_CHAR_DHLINE, ui.TCOD_CHAR_DHLINE, ui.TCOD_CHAR_DNE,
		ui.TCOD_CHAR_DVLINE, ui.TCOD_CHAR_BLOCK1, ui.TCOD_CHAR_BLOCK2, ui.TCOD_CHAR_BLOCK3, ui.TCOD_CHAR_DVLINE,
		ui.TCOD_CHAR_DVLINE, ui.TCOD_CHAR_BLOCK1, ui.TCOD_CHAR_BLOCK2, ui.TCOD_CHAR_BLOCK3, ui.TCOD_CHAR_DVLINE,
		ui.TCOD_CHAR_DVLINE, ui.TCOD_CHAR_BLOCK1, ui.TCOD_CHAR_BLOCK2, ui.TCOD_CHAR_BLOCK3, ui.TCOD_CHAR_DVLINE,
		ui.TCOD_CHAR_DSW, ui.TCOD_CHAR_DHLINE, ui.TCOD_CHAR_DHLINE, ui.TCOD_CHAR_DHLINE, ui.TCOD_CHAR_DSE,
	}

	viewport := s.e.screen.Get("Viewport")

	xOff := 12
	yOff := 12

	x := 0

	for _, c := range order {
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				bgColour := ui.ColourNC
				if ui.CurrentTheme.Colours[c] == ui.ColourAnsiBlack {
					bgColour = ui.ColourAnsiLightGrey
				}
				viewport.SetChar(square[y*5+x], position.Position{xOff + x, yOff + y}, ui.CurrentTheme.Colours[c], bgColour)
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
}

func (s *LobbyState) Update(dt float32) {
	if rl.IsKeyPressed(rl.KeySpace) {
		s.e.ChangeState(NewWorld(s.e))
	}

	if rl.IsKeyPressed(rl.KeyF) {
		s.e.screen.Get("Viewport").Clear()
		s.DrawColourSquares()
	}
}

func (s LobbyState) GetName() string {
	return "Lobby"
}
