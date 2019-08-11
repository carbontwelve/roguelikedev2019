package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"raylibtinkering/position"
	"raylibtinkering/ui"
	"unicode/utf8"
)

type btnTextAlign uint

const (
	BtnTextLeft btnTextAlign = iota
	BtnTextCenter
	BtnTextRight
)

type Button struct {
	ui.Component
}

func NewButton(name, title string, width, padX, padY uint, offX, offY int, align btnTextAlign, border ui.BorderColour, onClick func()) *Button {
	var txtPos position.Position

	titleLen := uint(utf8.RuneCountInString(title))

	if width < titleLen {
		width = titleLen
	}

	w := padX*2 + width
	h := 1 + padY*2

	mL := &Button{
		Component: *ui.NewComponent(name, w, h, offX, offY, true),
	}

	if align == BtnTextLeft {
		txtPos = position.Position{int(padX), int(h / 2)}
	} else if align == BtnTextCenter {
		txtPos = position.Position{int(int(w/2) - utf8.RuneCountInString(title)/2), int(h / 2)}
	} else if align == BtnTextRight {
		txtPos = position.Position{int(int(w) - utf8.RuneCountInString(title) - int(padX)), int(h / 2)}
	}

	mL.SetAutoClear(false) // We will "redraw" on hover
	mL.SetBorderStyle(ui.SingleWallBorder)
	mL.SetString(title, txtPos, ui.GameColours["Fg"], ui.ColourNC)

	//fmt.Println(fmt.Sprintf("W,H (%d, %d), TextPos (%d, %d), Textlen %d", w,h, txtPos.X, txtPos.Y, titleLen))

	hovering := false

	iH := func(component *ui.Component) {
		if ui.MousePos.X >= offX && ui.MousePos.X <= offX+int(w) && ui.MousePos.Y >= offY && ui.MousePos.Y <= offY+int(h) {
			mL.Clear()
			mL.SetString(title, txtPos, ui.GameColours["AnsiPurple"], ui.ColourNC)
			mL.SetBorderColour(ui.BorderColour{"AnsiPurple", "AnsiPurple", "AnsiPurple", "AnsiPurple"})
			hovering = true
		} else {
			//if hovering == true {
			mL.Clear()
			mL.SetString(title, txtPos, ui.GameColours["Fg"], ui.ColourNC)
			mL.SetBorderColour(border)
			hovering = false
			//}
		}

		if hovering && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			onClick()
		}
	}

	mL.SetInputHandler(iH)

	return mL
}
