package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Tile struct {
	r rl.Rectangle
}

type TileSheet struct {
	TxTiles rl.Texture2D // Sprite sheet texture
	TWidth  int
	THeight int
	Cols    int
	Rows    int
	Tiles   []*Tile
}

func newSpriteSheet(tx rl.Texture2D, w, h int) *TileSheet {
	cols := int(math.Ceil(float64(tx.Width / int32(w))))
	rows := int(math.Ceil(float64(tx.Height / int32(h))))

	s := &TileSheet{
		TxTiles: tx,
		TWidth:  w,
		THeight: h,
		Cols:    cols,
		Rows:    rows,
		Tiles:   make([]*Tile, cols*rows),
	}

	// 320 x 90
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			s.Set(x, y, &Tile{
				r: rl.NewRectangle(float32(x*w), float32(x*h), float32(w), float32(h)),
			})
			fmt.Println("Setting Tile (", x, ",", y, ") as (", x*w, ",", y*h, ",", (x*w)+10, ",", (y*h)+10, ")")
		}
	}

	return s
}

func (s *TileSheet) Unload() {
	rl.UnloadTexture(s.TxTiles)
}

func (t TileSheet) At(x, y int) *Tile {
	return t.Tiles[t.IdxAt(x, y)]
}

func (t TileSheet) IdxAt(x, y int) int {
	return (t.Cols * y) + x
}

func (t *TileSheet) Set(x, y int, val *Tile) {
	t.Tiles[t.IdxAt(x, y)] = val
}
