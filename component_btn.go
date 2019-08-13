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

type iButtonStyle struct {
	backColour  string
	textColour  string
	borderStyle ui.BorderStyle
	borderColor ui.BorderColour
}

type ButtonStyle struct {
	normal iButtonStyle
	hover  iButtonStyle
}

var DefaultButtonStyle = ButtonStyle{
	normal: iButtonStyle{
		backColour:  "Bg",
		textColour:  "Fg",
		borderStyle: ui.SingleWallBorder,
		borderColor: ui.DefaultBorderColour,
	},
	hover: iButtonStyle{
		backColour:  "Bg",
		textColour:  "AnsiPurple",
		borderStyle: ui.SingleWallBorder,
		borderColor: ui.BorderColour{"AnsiPurple", "AnsiPurple", "AnsiPurple", "AnsiPurple"},
	},
}

type Button struct {
	ui.Component
	style ButtonStyle
}

func NewButton(name, title string, width, padX, padY uint, offX, offY int, align btnTextAlign, style ButtonStyle, onClick func()) *Button {
	var txtPos position.Position

	titleLen := uint(utf8.RuneCountInString(title))

	if width < titleLen {
		width = titleLen
	}

	w := padX*2 + width
	h := 1 + padY*2

	mL := &Button{
		Component: *ui.NewComponent(name, w, h, offX, offY, true),
		style:     style,
	}

	if align == BtnTextLeft {
		txtPos = position.Position{int(padX), int(h / 2)}
	} else if align == BtnTextCenter {
		txtPos = position.Position{int(int(w/2) - utf8.RuneCountInString(title)/2), int(h / 2)}
	} else if align == BtnTextRight {
		txtPos = position.Position{int(int(w) - utf8.RuneCountInString(title) - int(padX)), int(h / 2)}
	}

	mL.SetAutoClear(false) // We will "redraw" on hover
	mL.SetBorderColour(style.normal.borderColor)
	mL.SetBorderStyle(style.normal.borderStyle)
	mL.SetString(title, txtPos, ui.GameColours[style.normal.textColour], ui.GameColours[style.normal.backColour])

	//fmt.Println(fmt.Sprintf("W,H (%d, %d), TextPos (%d, %d), Textlen %d", w,h, txtPos.X, txtPos.Y, titleLen))

	hovering := false

	iH := func(component *ui.Component) {
		if ui.MousePos.X >= offX && ui.MousePos.X <= offX+int(w) && ui.MousePos.Y >= offY && ui.MousePos.Y <= offY+int(h) {
			mL.Clear()
			mL.SetString(title, txtPos, ui.GameColours[style.hover.textColour], ui.GameColours[style.hover.backColour])
			mL.SetBorderColour(mL.style.hover.borderColor)
			mL.SetBorderStyle(mL.style.hover.borderStyle)
			hovering = true
		} else {
			mL.Clear()
			mL.SetString(title, txtPos, ui.GameColours[style.normal.textColour], ui.GameColours[style.normal.backColour])
			mL.SetBorderColour(mL.style.normal.borderColor)
			mL.SetBorderStyle(mL.style.normal.borderStyle)
			hovering = false
		}

		if hovering && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			onClick()
		}
	}

	mL.SetInputHandler(iH)

	return mL
}
