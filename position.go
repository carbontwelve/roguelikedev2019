package main

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
