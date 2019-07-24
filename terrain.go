package main

import (
	"fmt"
	"math"
	"strings"
)

// Define terrain enum, used to determine if a map cell
// is a wall or free
type cellType int

//
// Terrain Generator functions return the players starting Position.
//
type TerrainGeneratorFunc func(*Terrain) Position

func TestTerrainGenerator(t *Terrain) Position {
	t.Fill(FreeCell)
	t.SetCell(Position{30, 22}, WallCell)
	t.SetCell(Position{31, 22}, WallCell)
	t.SetCell(Position{32, 22}, WallCell)
	return Position{30, 23}
}

const (
	WallCell cellType = iota
	FreeCell
)

//
// A Map is made up of cells, each cell can be either a wall or free.
// If the cell has entered the players FOV then its Explored flag
// will be set to true.
//
type Cell struct {
	T        cellType
	Explored bool
}

type Terrain struct {
	w, h  int
	Cells []Cell
}

//
// Runs the provided terrain generator function against
// this struct.
//
func (t *Terrain) Generate(f TerrainGeneratorFunc) Position {
	return f(t)
}

func (t *Terrain) Cell(pos Position) Cell {
	fmt.Println(pos.idx())
	return t.Cells[pos.idx()]
}

func (t *Terrain) SetCell(pos Position, c cellType) {
	t.Cells[pos.idx()].T = c
}

func (t *Terrain) SetExplored(pos Position) {
	t.Cells[pos.idx()].Explored = true
}

func (t Terrain) ToString() string {
	var sb strings.Builder

	for y := 0; y < t.h; y++ {
		for x := 0; x < t.w; x++ {
			if t.Cell(Position{x, y}).T == WallCell {
				sb.WriteString("#")
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

//
// Fill Terrain with tiles
// @todo extend this to use a selection?
//
func (t *Terrain) Fill(c cellType) {
	for x := 0; x < t.w; x++ {
		for y := 0; y < t.h; y++ {
			t.SetCell(Position{x, y}, c)
		}
	}
}

//
// Returns the center cell of the terrain
// @todo bounds check
//
func (t Terrain) Center() Position {
	return Position{int(math.Floor(float64(t.w/2))) - 1, int(math.Floor(float64(t.h/2))) - 1}
}

//
// Returns whether the coordinate is inside the map bounds.
//
func (t Terrain) Inside(pos Position) bool {
	return pos.valid(t.w, t.h)
}

func GenTutorialTerrain(w, h int) *Terrain {
	d := &Terrain{w: w, h: h}
	d.Cells = make([]Cell, h*w)
	d.Fill(WallCell)

	d.SetCell(Position{10, 10}, FreeCell)

	return d
}

func NewTerrain(w, h int) *Terrain {
	d := &Terrain{w: w, h: h}
	d.Cells = make([]Cell, h*w)
	d.Fill(FreeCell)
	return d
}

//
// Dig a room into the terrain
// @todo set what terrain type to "dig"
//
func (t *Terrain) AddRoom(room Rect) {
	for x := room.x1 + 1; x < room.x2; x++ {
		for y := room.y1 + 1; y < room.y2; y++ {
			pos := Position{x, y}
			if pos.valid(t.w, t.h) {
				t.SetCell(pos, FreeCell)
			}
		}
	}
}

//
// Dig a horizontal tunnel through the terrain
// @todo set what terrain type to "dig"
//
func (t *Terrain) AddHTunnel(x1, x2, y int) {
	for x := int(math.Min(float64(x1), float64(x2))); x < int(math.Max(float64(x1), float64(x2)))+1; x++ {
		pos := Position{x, y}
		if pos.valid(t.w, t.h) {
			t.SetCell(pos, FreeCell)
		}
	}
}

//
// Dig a vertical tunnel through the terrain
// @todo set what terrain type to "dig"
//
func (t *Terrain) AddVTunnel(y1, y2, x int) {
	for y := int(math.Min(float64(y1), float64(y2))); y < int(math.Max(float64(y1), float64(y2)))+1; y++ {
		pos := Position{x, y}
		if pos.valid(t.w, t.h) {
			t.SetCell(pos, FreeCell)
		}
	}
}
