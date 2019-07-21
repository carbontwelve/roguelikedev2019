package main

//
// A tile on a map. It may or may not be blocked, and may or may not block sight.
//
type Tile struct {
	blocked    bool // Does this block the players movement
	blockSight bool // Does this block the players FOV
	inFOV      bool // Is this tile within the players FOV?
	seen       bool // Has the player seen this tile?
}

func NewTile(blocked, blockSight bool) *Tile {
	return &Tile{blocked: blocked, blockSight: blockSight, inFOV: false, seen: false}
}

func (t *Tile) SetInFOV() {
	t.inFOV = true
	t.seen = true
}
