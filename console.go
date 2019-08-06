package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"raylibtinkering/position"
	"unicode/utf8"
)

type orderedComponent struct {
	v     *Viewport
	order int
}

//
// A Screen defines an overall drawable area and acts as a container for
// Viewports. When Viewports are added to the Screen they can be given
// a render order. This act to allow a form of "layering"
//

//
// Screen:
//
// This acts as a container for component composition. When initiated you pass it the windows width and height
// then a Tileset. Screen will use the Tileset and the windows width/height to work out how many rows and columns
// can be fit into the available windowed area.
//
// Once initiated you can add components with the Set function. Components are to be added with their draw
// order set as zIndex. This is so you can effectively layer components on top of one another.
//
// With each game loop iteration you should call HandleEvents followed by Draw. HandleEvents will allow components
// to respond to user input.
//
type Screen struct {
	width, height uint // In pixels e.g 800x600
	components    map[string]*orderedComponent
	positionCache map[position.Position]string // cache of component position so we can tell if a mouse pointer is hovering
	drawOrder     []string
	tileset       int // placeholder
}

func (s *Screen) HandleEvents() {
	// @todo loop over components and handle any user input per component e.g for buttons
}

func (s Screen) Draw() {
	// @todo draw to render interface... this can in future be terminal or graphical
}

func (s *Screen) Get(k string) *Viewport {
	return s.components[k].v
}

func (s *Screen) Set(k string, v *Viewport, zIndex int) {
	s.components[k] = &orderedComponent{v: v, order: zIndex}

	// @todo populate drawOrder
}

func NewScreen(w, h uint) *Screen {
	return &Screen{width: w, height: h, components: make(map[string]*orderedComponent), drawOrder: make([]string, 0)}
}

type cCellBorder struct {
	V, H, NE, SE, SW, NW int
}

type cCell struct {
	char   int
	fg, bg rl.Color
	x, y   uint // Offset x,y based upon Viewport x + cCell Position in the cells map
}

func (c *cCell) Reset() {
	c.char = *new(int)
	c.fg = *new(rl.Color)
	c.bg = *new(rl.Color)
}

func (c cCell) GetDrawPosition() position.Position {
	return position.Position{X: int(c.x), Y: int(c.y)}
}

//
// A viewport is an area on the screen that can be drawn to.
// TileGrid...
//
type Viewport struct {
	width, height uint
	x, y          uint
	cells         map[position.Position]*cCell
	modified      map[position.Position]bool
	bordered      bool
	border        cCellBorder
}

func NewViewport(w, h, x, y uint) *Viewport {
	vp := &Viewport{
		width:    w,
		height:   h,
		cells:    make(map[position.Position]*cCell),
		x:        x,
		y:        y,
		bordered: false,
		border:   cCellBorder{TCOD_CHAR_VLINE, TCOD_CHAR_HLINE, TCOD_CHAR_NE, TCOD_CHAR_SE, TCOD_CHAR_SW, TCOD_CHAR_NW},
	}

	for cY := uint(0); cY < h; cY++ {
		for cX := uint(0); cX < w; cX++ {
			vp.cells[position.Position{X: int(cX), Y: int(cY)}] = &cCell{x: cX + x, y: cY + y}
		}
	}
	return vp
}

func (v Viewport) GetCells() map[position.Position]*cCell {
	return v.cells
}

func (v Viewport) GetUpdatedCells() map[position.Position]*cCell {
	// @todo this will return just those cells that are found within the modified hash
	// it will then reset the modified hash. Used to get only cells updated since last
	// call. May be useful for speeding up things if needed...
	return make(map[position.Position]*cCell)
}

func (v *Viewport) SetBorder(b cCellBorder) {
	v.border = b
}

func (v *Viewport) SetBordered(b bool) {
	v.bordered = b
	if b {
		v.DrawBorder()
	}
}

func (v *Viewport) DrawBorder() {
	if !v.bordered {
		return
	}
	for x := uint(0); x < v.width-1; x++ {
		v.SetChar(v.border.H, position.Position{X: int(x), Y: 0}, ColourUiLines, ColourBg)
		v.SetChar(v.border.H, position.Position{X: int(x), Y: int(v.height - 1)}, ColourUiLines, ColourBg)
	}

	for y := uint(0); y < v.height-1; y++ {
		v.SetChar(v.border.V, position.Position{X: 0, Y: int(y)}, ColourUiLines, ColourBg)
		v.SetChar(v.border.V, position.Position{X: int(v.width - 1), Y: int(y)}, ColourUiLines, ColourBg)
	}

	v.SetChar(v.border.NE, position.Position{X: int(v.width - 1), Y: 0}, ColourUiLines, ColourBg)
	v.SetChar(v.border.SE, position.Position{X: int(v.width - 1), Y: int(v.height - 1)}, ColourUiLines, ColourBg)
	v.SetChar(v.border.SW, position.Position{X: 0, Y: int(v.height - 1)}, ColourUiLines, ColourBg)
	v.SetChar(v.border.NW, position.Position{X: 0, Y: 0}, ColourUiLines, ColourBg)
}

func (v *Viewport) SetChar(r int, p position.Position, fg, bg rl.Color) {
	c := v.cells[p]
	c.char = r
	c.bg = bg
	c.fg = fg
}

func (v *Viewport) ClearRow(y uint) {
	var (
		xMin, xMax uint
	)

	if v.bordered {
		xMin = 1
		xMax = v.width - 2
	} else {
		xMin = 0
		xMax = v.width
	}

	for x := xMin; x < xMax; x++ {
		v.cells[position.Position{X: int(x), Y: int(y)}].Reset()
	}
}

func (v *Viewport) ClearCol(x uint, r rune) {
	// ...
}

func (v *Viewport) SetRow(str string, p position.Position, fg, bg rl.Color) {
	v.ClearRow(uint(p.Y))
	v.SetString(str, p, fg, bg)
}

func (v *Viewport) SetString(str string, p position.Position, fg, bg rl.Color) {
	maxX := uint(p.X + utf8.RuneCountInString(str))
	if maxX > v.width {
		maxX = v.width // truncate...
	}

	a := []rune(str)
	for i, r := range a {
		x := p.X + i
		c := v.cells[position.Position{X: x, Y: p.Y}]
		c.char = int(r)
		c.bg = bg
		c.fg = fg
	}
}
