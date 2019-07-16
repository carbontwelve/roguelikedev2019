package main

import rl "github.com/gen2brain/raylib-go/raylib"

type TileSheet struct {
	TxTiles rl.Texture2D // Sprite sheet texture
	TWidth  int
	Theight int
}

func newSpriteSheet(tx rl.Texture2D, w, h int) *TileSheet {
	s := &TileSheet{
		TxTiles: tx,
		TWidth:  w,
		Theight: h,
	}

	return s
}

func (s *TileSheet) Unload() {
	rl.UnloadTexture(s.TxTiles)
}
