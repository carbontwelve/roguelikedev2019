package ui

import (
	"raylibtinkering/position"
	"sort"
)

type componentZOrder struct {
	c     ComponentI
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
	width, height     uint // In pixels e.g 800x600
	Rows, Cols        uint
	components        map[string]*componentZOrder
	positionCache     map[position.Position]string // cache of component position so we can tell if a mouse pointer is hovering
	handleEventsCache []string                     // List of components that have event handlers set
	drawOrder         []*componentZOrder
	dirty             bool
	tileset           *Tileset
}

func (s *Screen) HandleEvents() {
	if len(s.handleEventsCache) == 0 {
		return
	}

	// @todo loop over components and handle any user input per component e.g for buttons
	for _, name := range s.handleEventsCache {
		s.components[name].c.HandleUserInput()
	}

}

func (s *Screen) Draw() {
	if s.dirty {
		s.drawOrder = make([]*componentZOrder, 0)

		for _, v := range s.components {
			if v.c.IsVisible() {
				s.drawOrder = append(s.drawOrder, v)
			}
		}

		sort.Slice(s.drawOrder, func(i, j int) bool {
			return s.drawOrder[i].order < s.drawOrder[j].order
		})
		s.dirty = false
	}

	for _, kv := range s.drawOrder {
		for _, cell := range kv.c.GetCells() {
			s.tileset.Draw(cell.char, cell.GetDrawPosition(), cell.fg, cell.bg)
		}

		if kv.c.AutoClears() == true {
			kv.c.Clear()
		}
	}
}

func (s Screen) Get(k string) ComponentI {
	return s.components[k].c
}

func (s *Screen) Set(c ComponentI, zIndex int) {
	s.components[c.GetName()] = &componentZOrder{c: c, order: zIndex}
	s.dirty = true

	if c.HasInputHandler() {
		s.handleEventsCache = append(s.handleEventsCache, c.GetName())
	}
}

func (s *Screen) Reset() {
	s.components = make(map[string]*componentZOrder)
	s.drawOrder = make([]*componentZOrder, 0)
	s.handleEventsCache = make([]string, 0)
	s.positionCache = make(map[position.Position]string)
	s.dirty = true
}

func (s *Screen) Unload() {
	s.tileset.Unload()
}

func NewScreen(w, h uint, t *Tileset) (error, *Screen) {
	screen := &Screen{dirty: true, width: w, height: h, components: make(map[string]*componentZOrder), drawOrder: make([]*componentZOrder, 0), tileset: t}
	screen.Cols = w / uint(t.sprites.TileWidth)
	screen.Rows = h / uint(t.sprites.TileHeight)
	return nil, screen
}
