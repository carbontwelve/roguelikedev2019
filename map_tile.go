package main

//
// A tile on a map. It may or may not be blocked, and may or may not block sight.
//
type Tile struct {
	blocked    bool
	blockSight bool
}

func NewTile(blocked, blockSight bool) *Tile {
	return &Tile{blocked: blocked, blockSight: blockSight}
}
