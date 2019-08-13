package ui

import (
	"fmt"
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
	SetTitle(title string)
	SetBorderStyle(bs BorderStyle)
	SetBorderColour(bc BorderColour)
	ReDraw()
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

type TitleStyle struct {
	leftChar, rightChar       uint
	paddingLeft, paddingRight uint
}

type BorderStyle struct {
	V, H, NE, SE, SW, NW uint
}

// Colours in order: Top, Right, Bottom, Left
type BorderColour [4]string

var DefaultBorderColour = BorderColour{"UiLines", "UiLines", "UiLines", "UiLines"}

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
	borderColor      BorderColour
	border           BorderStyle
	bordered         bool
	cells            map[position.Position]*Cell
	camera           *Camera
	visible          bool
	autoClear        bool
	title            string
	titled           bool
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
		c.ReDraw()
	}
}

func (c *Component) SetBorderColour(bc BorderColour) {
	c.borderColor = bc
	if c.bordered {
		c.ReDraw()
	}
}

func (c *Component) drawBorder() {
	if !c.bordered {
		return
	}

	// Top/Bottom
	for x := uint(0); x < c.Width-1; x++ {
		c.SetChar(c.border.H, position.Position{X: int(x), Y: 0}, GameColours[c.borderColor[0]], GameColours["Bg"])
		c.SetChar(c.border.H, position.Position{X: int(x), Y: int(c.Height - 1)}, GameColours[c.borderColor[2]], GameColours["Bg"])
	}

	// Left/Right
	for y := uint(0); y < c.Height-1; y++ {
		c.SetChar(c.border.V, position.Position{X: 0, Y: int(y)}, GameColours[c.borderColor[3]], GameColours["Bg"])
		c.SetChar(c.border.V, position.Position{X: int(c.Width - 1), Y: int(y)}, GameColours[c.borderColor[1]], GameColours["Bg"])
	}

	c.SetChar(c.border.NE, position.Position{X: int(c.Width - 1), Y: 0}, GameColours[c.borderColor[0]], GameColours["Bg"])
	c.SetChar(c.border.SE, position.Position{X: int(c.Width - 1), Y: int(c.Height - 1)}, GameColours[c.borderColor[1]], GameColours["Bg"])
	c.SetChar(c.border.SW, position.Position{X: 0, Y: int(c.Height - 1)}, GameColours[c.borderColor[2]], GameColours["Bg"])
	c.SetChar(c.border.NW, position.Position{X: 0, Y: 0}, GameColours[c.borderColor[3]], GameColours["Bg"])
}

// Redraw the permanent style of this component. e.g Border, Title, etc
func (c *Component) ReDraw() {
	c.drawBorder() // Border comes before title
	c.drawTitle()  // Title is drawn on top of border and replaces some tiles
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

func (c *Component) SetTitle(title string) {
	c.title = title
	c.titled = c.title != ""
	c.drawTitle()
}

func (c *Component) drawTitle() {
	if c.titled == false {
		return
	}

	t := fmt.Sprintf("] %s [", c.title)
	c.SetString(t, position.Position{2, 0}, GameColours["Fg"], GameColours["Bg"])
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
		Name:        name,
		Width:       w,
		Height:      h,
		OffsetX:     offX,
		OffsetY:     offY,
		cells:       make(map[position.Position]*Cell),
		visible:     visible,
		autoClear:   true,
		borderColor: DefaultBorderColour,
	}

	for cY := uint(0); cY < h; cY++ {
		for cX := uint(0); cX < w; cX++ {
			component.cells[position.Position{X: int(cX), Y: int(cY)}] = &Cell{x: int(cX) + offX, y: int(cY) + offY}
		}
	}

	return component
}
