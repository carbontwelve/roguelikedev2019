package main

import "github.com/gen2brain/raylib-go/raylib"

//
// @todo make this configurable?
//
const DungeonWidth = 80
const DungeonHeight = 45

//
// Data structure for storing cardinal coordinates
//
type Position struct {
	X, Y int
}

//
// Create a new Position for array index
//
func idxtopos(i int) Position {
	return Position{i % DungeonWidth, i / DungeonWidth}
}

//
// Get the positions array index
//
func (pos Position) idx() int {
	return pos.Y*DungeonWidth + pos.X
}

//
// Check if the Position is valid for the current dungeon size
//
func (pos Position) valid(w, h int) bool {
	return pos.Y >= 0 && pos.Y < h && pos.X >= 0 && pos.X < w
}

func (pos Position) N() Position {
	return Position{pos.X, pos.Y - 1}
}

func (pos Position) E() Position {
	return Position{pos.X + 1, pos.Y}
}

func (pos Position) S() Position {
	return Position{pos.X, pos.Y + 1}
}

func (pos Position) W() Position {
	return Position{pos.X - 1, pos.Y}
}

func (pos Position) Vector2(exp int) rl.Vector2 {
	return rl.Vector2{X: float32(pos.X * exp), Y: float32(pos.Y * exp)}
}
