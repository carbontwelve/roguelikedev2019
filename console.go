package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"unicode/utf8"
)

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

func (c cCell) GetDrawPosition() Position {
	return Position{X: int(c.x), Y: int(c.y)}
}

//
// A viewport is an area on the screen that can be drawn to.
//
type Viewport struct {
	width, height uint
	x, y          uint
	cells         map[Position]*cCell
	bordered      bool
	border        cCellBorder
}

func NewViewport(w, h, x, y uint) *Viewport {
	vp := &Viewport{
		width:    w,
		height:   h,
		cells:    make(map[Position]*cCell),
		x:        x,
		y:        y,
		bordered: false,
		border:   cCellBorder{TCOD_CHAR_VLINE, TCOD_CHAR_HLINE, TCOD_CHAR_NE, TCOD_CHAR_SE, TCOD_CHAR_SW, TCOD_CHAR_NW},
	}

	for cY := uint(0); cY < h; cY++ {
		for cX := uint(0); cX < w; cX++ {
			vp.cells[Position{X: int(cX), Y: int(cY)}] = &cCell{x: cX + x, y: cY + y}
		}
	}
	return vp
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
		v.SetChar(v.border.H, Position{X: int(x), Y: 0}, rl.Orange, rl.Black)
		v.SetChar(v.border.H, Position{X: int(x), Y: int(v.height - 1)}, rl.Orange, rl.Black)
	}

	for y := uint(0); y < v.height-1; y++ {
		v.SetChar(v.border.V, Position{X: 0, Y: int(y)}, rl.Orange, rl.Black)
		v.SetChar(v.border.V, Position{X: int(v.width - 1), Y: int(y)}, rl.Orange, rl.Black)
	}

	v.SetChar(v.border.NE, Position{X: int(v.width - 1), Y: 0}, rl.Orange, rl.Black)
	v.SetChar(v.border.SE, Position{X: int(v.width - 1), Y: int(v.height - 1)}, rl.Orange, rl.Black)
	v.SetChar(v.border.SW, Position{X: 0, Y: int(v.height - 1)}, rl.Orange, rl.Black)
	v.SetChar(v.border.NW, Position{X: 0, Y: 0}, rl.Orange, rl.Black)
}

func (v *Viewport) SetChar(r int, p Position, fg, bg rl.Color) {
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
		v.cells[Position{X: int(x), Y: int(y)}].Reset()
	}
}

func (v *Viewport) ClearCol(x uint, r rune) {
	// ...
}

func (v *Viewport) SetString(str string, p Position, fg, bg rl.Color) {
	maxX := uint(p.X + utf8.RuneCountInString(str))
	if maxX > v.width {
		maxX = v.width // truncate...
	}

	a := []rune(str)
	for i, r := range a {
		x := p.X + i
		c := v.cells[Position{X: x, Y: p.Y}]
		c.char = int(r)
		c.bg = bg
		c.fg = fg
	}
}
