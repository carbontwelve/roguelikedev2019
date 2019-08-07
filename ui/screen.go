package ui

import (
	"raylibtinkering/position"
)

type componentZOrder struct {
	c     *Component
	order int
}

//
// Screen:
//
// This acts as a container for component composition. When initiated you pass it the windows width and height
// then a Tileset. Screen will use the Tileset and the windows width/height to work out how many Rows and columns
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
	Rows, Cols    uint
	components    map[string]*componentZOrder
	positionCache map[position.Position]string // cache of component position so we can tell if a mouse pointer is hovering
	drawOrder     []string
	tileset       *Tileset
}

func (s *Screen) HandleEvents() {
	// @todo loop over components and handle any user input per component e.g for buttons
}

func (s Screen) Draw() {
	// @todo draw to render interface... this can in future be terminal or graphical
}

func (s Screen) Get(k string) *Component {
	return s.components[k].c
}

func (s *Screen) Set(c *Component, zIndex int) {
	s.components[c.Name] = &componentZOrder{c: c, order: zIndex}

	// @todo populate drawOrder
}

func (s *Screen) Reset() {
	s.components = make(map[string]*componentZOrder)
	s.drawOrder = make([]string, 0)
	s.positionCache = make(map[position.Position]string)
}

func (s *Screen) Unload() {
	s.tileset.Unload()
}

func NewScreen(w, h uint, t *Tileset) (error, *Screen) {
	screen := &Screen{width: w, height: h, components: make(map[string]*componentZOrder), drawOrder: make([]string, 0), tileset: t}

	// @todo check tile width and height are > 0
	screen.Cols = w / uint(t.sprites.TileWidth)
	screen.Rows = h / uint(t.sprites.TileHeight)

	return nil, screen
}
