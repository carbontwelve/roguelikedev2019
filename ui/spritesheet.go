package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Sprite struct {
	R rl.Rectangle
	t *SpriteSheet
}

func (t Sprite) Draw(position rl.Vector2, tint rl.Color) {
	rl.DrawTextureRec(t.t.TxTiles, t.R, position, tint)
}

type SpriteSheet struct {
	TxTiles    rl.Texture2D // Sprite sheet texture
	TileWidth  int
	TileHeight int
	Cols       int
	Rows       int
	Tiles      []*Sprite
}

func NewSpriteSheet(tx rl.Texture2D, w, h int) *SpriteSheet {
	cols := int(math.Ceil(float64(tx.Width / int32(w))))
	rows := int(math.Ceil(float64(tx.Height / int32(h))))

	tileSheet := &SpriteSheet{
		TxTiles:    tx,
		TileWidth:  w,
		TileHeight: h,
		Cols:       cols,
		Rows:       rows,
		Tiles:      make([]*Sprite, cols*rows),
	}

	// 320 x 90
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			tileSheet.Set(x, y, &Sprite{
				R: rl.NewRectangle(float32(x*w), float32(y*h), float32(w), float32(h)),
				t: tileSheet,
			})
			// fmt.Println("Setting Sprite (", x, ",", y, ") Idx (", tileSheet.IdxAt(x, y), ") as Rect(", x*w, ",", y*h, ",", (x*w)+10, ",", (y*h)+10, ")")
		}
	}

	return tileSheet
}

func (t *SpriteSheet) Unload() {
	rl.UnloadTexture(t.TxTiles)
}

func (t SpriteSheet) At(x, y int) *Sprite {
	return t.Tiles[t.IdxAt(x, y)]
}

func (t SpriteSheet) AtIdx(idx int) *Sprite {
	// @todo add bounds check
	return t.Tiles[idx]
}

func (t SpriteSheet) IdxAt(x, y int) int {
	return (t.Cols * y) + x
}

func (t *SpriteSheet) Set(x, y int, val *Sprite) {
	t.Tiles[t.IdxAt(x, y)] = val
}
