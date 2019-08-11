package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"raylibtinkering/position"
	"raylibtinkering/ui"
)

type Button struct {
	ui.Component
}

//func (b *Button) Draw () {
//
//}

func NewButton(name string, w, h uint, offX, offY int, onClick func()) *Button {
	mL := &Button{
		Component: *ui.NewComponent(name, w, h, offX, offY, true),
	}

	hovering := false

	iH := func(component *ui.Component) {
		if ui.MousePos.X >= offX && ui.MousePos.X <= offX+int(w) && ui.MousePos.Y >= offY && ui.MousePos.Y <= offY+int(h) {
			mL.Clear()
			mL.SetString(name, position.Position{2, 2}, ui.GameColours["AnsiPurple"], ui.ColourNC)
			hovering = true
		} else {
			if hovering == true {
				mL.Clear()
				mL.SetString(name, position.Position{2, 2}, ui.GameColours["Fg"], ui.ColourNC)
			}
		}

		if hovering && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			onClick()
		}
	}

	mL.SetAutoClear(false)
	mL.SetBorderStyle(ui.SingleWallBorder)
	mL.SetString(name, position.Position{2, 2}, ui.GameColours["Fg"], ui.ColourNC)
	mL.SetInputHandler(iH)

	return mL
}
