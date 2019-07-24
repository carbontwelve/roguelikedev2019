package main

type World struct {
	Terrain        *Terrain
	Entities       []Entity
	seed           int64
	CustomSeed     bool
	playerStartPos Position
}

//
// Set a custom seed for this map
//
func (w *World) SetSeed(seed int64) {
	w.CustomSeed = true
	w.seed = seed
}
