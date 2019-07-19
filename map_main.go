package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
)

type GameMap struct {
	MWidth, MHeight, PlayerStartX, PlayerStartY int
	Tiles                                       []*Tile
}

func NewGameMap(w, h int) *GameMap {
	gameMap := GameMap{MWidth: w, MHeight: h, Tiles: make([]*Tile, w*h)}

	// Initialise the tiles
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			gameMap.Set(x, y, NewTile(true, true))
		}
	}

	// Wide open space with four walls
	//for x := 1; x < w-1; x++ {
	//	for y := 1; y < h-1; y++ {
	//		gameMap.At(x, y).blocked = false
	//		gameMap.At(x, y).blockSight = false
	//	}
	//}

	//gameMap.AddRoom(NewRect(20, 15, 10, 15))
	//gameMap.AddRoom(NewRect(35, 15, 10, 15))
	//gameMap.AddVTunnel(25, 40, 23)

	return &gameMap
}

func (m *GameMap) AddRoom(room Rect) {
	for x := room.x1 + 1; x < room.x2; x++ {
		for y := room.y1 + 1; y < room.y2; y++ {
			tile := m.At(x, y)
			tile.blocked = false
			tile.blockSight = false
		}
	}
}

func (m *GameMap) AddHTunnel(x1, x2, y int) {
	for x := int(math.Min(float64(x1), float64(x2))); x < int(math.Max(float64(x1), float64(x2)))+1; x++ {
		tile := m.At(x, y)
		tile.blocked = false
		tile.blockSight = false
	}
}

func (m *GameMap) AddVTunnel(y1, y2, x int) {
	for y := int(math.Min(float64(y1), float64(y2))); y < int(math.Max(float64(y1), float64(y2)))+1; y++ {
		tile := m.At(x, y)
		tile.blocked = false
		tile.blockSight = false
	}
}

func (m GameMap) Center() (int, int) {
	// @todo bounds check
	return int(math.Floor(float64(m.MWidth/2))) - 1, int(math.Floor(float64(m.MHeight/2))) - 1
}

func (m GameMap) Draw(engine *Engine) {
	for x := 0; x < m.MWidth; x++ {
		for y := 0; y < m.MHeight; y++ {
			wall := m.At(x, y).blocked
			if wall == true {
				// Draw Wall
				engine.font.Draw('#', rl.Vector2{X: float32(x * engine.font.sprites.TWidth), Y: float32(y * engine.font.sprites.THeight)}, rl.Gray)
			} else {
				engine.font.Draw('.', rl.Vector2{X: float32(x * engine.font.sprites.TWidth), Y: float32(y * engine.font.sprites.THeight)}, rl.RayWhite)
			}
		}
	}
}

func (m GameMap) IsBlocked(x, y int) bool {
	return m.At(x, y).blocked
}

func (m GameMap) At(x, y int) *Tile {
	return m.Tiles[m.IdxAt(x, y)]
}

// @todo add bounds check
func (m GameMap) AtIdx(idx int) *Tile {
	return m.Tiles[idx]
}

func (m GameMap) IdxAt(x, y int) int {
	return (m.MWidth * y) + x
}

func (m *GameMap) Set(x, y int, val *Tile) {
	m.Tiles[m.IdxAt(x, y)] = val
}
