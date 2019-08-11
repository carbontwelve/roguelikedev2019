package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"raylibtinkering/position"
	"unicode/utf8"
)

type ComponentI interface {
	GetCells() map[position.Position]*Cell
	IsVisible() bool
	AutoClears() bool
	GetName() string
	GetWidth() uint
	GetHeight() uint
	GetInnerHeight() uint
	GetInnerWidth() uint
	SetBorderStyle(bs BorderStyle)
	DrawBorder()
	SetChar(r uint, p position.Position, fg, bg rl.Color)
	Clear()
	ClearRow(y uint)
	ClearCol(x uint, r rune)
	HandleUserInput()
	SetInputHandler(iH ComponentInputEventHandler)
	HasInputHandler() bool
	SetRow(str string, p position.Position, fg, bg rl.Color)
	SetString(str string, p position.Position, fg, bg rl.Color)
	SetCamera(cam *Camera)
	SetAutoClear(b bool)
}

type BorderStyle struct {
	V, H, NE, SE, SW, NW uint
}

var SingleWallBorder = BorderStyle{
	TCOD_CHAR_VLINE, TCOD_CHAR_HLINE, TCOD_CHAR_NE, TCOD_CHAR_SE, TCOD_CHAR_SW, TCOD_CHAR_NW,
}

var ZeroWallBorder = BorderStyle{
	0, 0, 0, 0, 0, 0,
}

type ComponentInputEventHandler func(component *Component)

type Component struct {
	Name             string
	Width, Height    uint
	OffsetX, OffsetY int
	border           BorderStyle
	bordered         bool
	cells            map[position.Position]*Cell
	camera           *Camera
	visible          bool
	autoClear        bool
	inputHandler     ComponentInputEventHandler
}

func (c *Component) SetInputHandler(iH ComponentInputEventHandler) {
	c.inputHandler = iH
}

func (c Component) HasInputHandler() bool {
	return c.inputHandler != nil
}

func (c *Component) HandleUserInput() {
	c.inputHandler(c)
}

func (c Component) GetName() string {
	return c.Name
}

func (c Component) GetWidth() uint {
	return c.Width
}

func (c Component) GetHeight() uint {
	return c.Height
}

func (c Component) GetInnerHeight() uint {
	if c.bordered == false {
		return c.Height
	}

	return c.Height - 2
}

func (c Component) GetInnerWidth() uint {
	if c.bordered == false {
		return c.Width
	}
	return c.Width - 2
}

func (c Component) IsVisible() bool {
	return c.visible
}

func (c Component) AutoClears() bool {
	return c.autoClear
}

func (c *Component) SetBorderStyle(bs BorderStyle) {
	c.border = bs
	c.bordered = bs != ZeroWallBorder

	if c.bordered {
		c.DrawBorder()
	}
}

func (c *Component) DrawBorder() {
	if !c.bordered {
		return
	}
	for x := uint(0); x < c.Width-1; x++ {
		c.SetChar(c.border.H, position.Position{X: int(x), Y: 0}, GameColours["UiLines"], GameColours["bg"])
		c.SetChar(c.border.H, position.Position{X: int(x), Y: int(c.Height - 1)}, GameColours["UiLines"], GameColours["bg"])
	}

	for y := uint(0); y < c.Height-1; y++ {
		c.SetChar(c.border.V, position.Position{X: 0, Y: int(y)}, GameColours["UiLines"], GameColours["bg"])
		c.SetChar(c.border.V, position.Position{X: int(c.Width - 1), Y: int(y)}, GameColours["UiLines"], GameColours["bg"])
	}

	c.SetChar(c.border.NE, position.Position{X: int(c.Width - 1), Y: 0}, GameColours["UiLines"], GameColours["bg"])
	c.SetChar(c.border.SE, position.Position{X: int(c.Width - 1), Y: int(c.Height - 1)}, GameColours["UiLines"], GameColours["bg"])
	c.SetChar(c.border.SW, position.Position{X: 0, Y: int(c.Height - 1)}, GameColours["UiLines"], GameColours["bg"])
	c.SetChar(c.border.NW, position.Position{X: 0, Y: 0}, GameColours["UiLines"], GameColours["bg"])
}

func (c *Component) SetChar(r uint, p position.Position, fg, bg rl.Color) {
	cell := c.cells[p]
	if cell == nil {
		return
	}
	cell.char = r
	cell.bg = bg
	cell.fg = fg
}

func (c *Component) Clear() {
	var (
		xMin, xMax, yMin, yMax uint
	)

	if c.bordered {
		xMin = 1
		xMax = c.Width - 2
		yMin = 1
		yMax = c.Height - 2
	} else {
		xMin = 0
		xMax = c.Width
		yMin = 0
		yMax = c.Height
	}

	for y := yMin; y < yMax; y++ {
		for x := xMin; x < xMax; x++ {
			c.cells[position.Position{X: int(x), Y: int(y)}].Reset()
		}
	}
}

func (c *Component) ClearRow(y uint) {
	var (
		xMin, xMax uint
	)

	if c.bordered {
		xMin = 1
		xMax = c.Width - 2
	} else {
		xMin = 0
		xMax = c.Width
	}

	for x := xMin; x < xMax; x++ {
		c.cells[position.Position{X: int(x), Y: int(y)}].Reset()
	}
}

func (c *Component) ClearCol(x uint, r rune) {
	// ...
}

func (c *Component) SetRow(str string, p position.Position, fg, bg rl.Color) {
	c.ClearRow(uint(p.Y))
	c.SetString(str, p, fg, bg)
}

func (c *Component) SetString(str string, p position.Position, fg, bg rl.Color) {
	maxX := uint(p.X + utf8.RuneCountInString(str))
	if maxX > c.Width {
		maxX = c.Width // truncate...
	}

	a := []rune(str)
	for i, r := range a {
		x := p.X + i
		c := c.cells[position.Position{X: x, Y: p.Y}]
		c.char = uint(r)
		c.bg = bg
		c.fg = fg
	}
}

func (c *Component) SetCamera(cam *Camera) {
	c.camera = cam
}

func (c *Component) SetAutoClear(b bool) {
	c.autoClear = b
}

func (c Component) GetCells() map[position.Position]*Cell {
	return c.cells
}

//
// Constructor
//
func NewComponent(name string, w, h uint, offX, offY int, visible bool) *Component {
	component := &Component{
		Name:      name,
		Width:     w,
		Height:    h,
		OffsetX:   offX,
		OffsetY:   offY,
		cells:     make(map[position.Position]*Cell),
		visible:   visible,
		autoClear: true,
	}

	for cY := uint(0); cY < h; cY++ {
		for cX := uint(0); cX < w; cX++ {
			component.cells[position.Position{X: int(cX), Y: int(cY)}] = &Cell{x: int(cX) + offX, y: int(cY) + offY}
		}
	}

	return component
}
