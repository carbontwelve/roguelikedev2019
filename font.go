package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

//
// The tcod_codec_ comes from https://github.com/libtcod/libtcod/blob/master/src/libtcod/sys_sdl_c.cpp#L165
// It is the codec for TCOD_FONT_LAYOUT_TCOD and converts from EASCII code-point -> raw tile position.
// BSD 3-Clause License
// Copyright © 2008-2019, Jice and the libtcod contributors. All rights reserved.
//

type tcod_codec_ [256]int

func getTcodCodec() *tcod_codec_ {
	return &tcod_codec_{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 76, 77, 0, 0, 0, 0, 0, /* 0 to 15 */
		71, 70, 72, 0, 0, 0, 0, 0, 64, 65, 67, 66, 0, 73, 68, 69, /* 16 to 31 */
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, /* 32 to 47 */
		16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, /* 48 to 63 */
		32, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, /* 64 to 79 */
		111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 33, 34, 35, 36, 37, /* 80 to 95 */
		38, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, /* 96 to 111 */
		143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 39, 40, 41, 42, 0, /* 112 to 127 */
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, /* 128 to 143 */
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, /* 144 to 159 */
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, /* 160 to 175 */
		43, 44, 45, 46, 49, 0, 0, 0, 0, 81, 78, 87, 88, 0, 0, 55, /* 176 to 191 */
		53, 50, 52, 51, 47, 48, 0, 0, 85, 86, 82, 84, 83, 79, 80, 0, /* 192 to 207 */
		0, 0, 0, 0, 0, 0, 0, 0, 0, 56, 54, 0, 0, 0, 0, 0, /* 208 to 223 */
		74, 75, 57, 58, 59, 60, 61, 62, 63, 0, 0, 0, 0, 0, 0, 0, /* 224 to 239 */
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, /* 240 to 255 */
	}
}

//
// Note: I am only implementing the TCOD Codec here, someone else can fork and
//       implement other formats if they need them :)
//
type Font struct {
	sprites  *TileSheet
	asciiMap map[int]int
}

func newFont(filename string, w, h int) *Font {
	font := &Font{
		sprites:  newSpriteSheet(rl.LoadTexture(filename), w, h),
		asciiMap: make(map[int]int),
	}

	font.decode()
	return font
}

func (f *Font) decode() {
	codec := getTcodCodec()

	for i := 0; i < 256; i++ {
		f.mapAsciiToFont(i, codec[i], 0)
	}
}

func (f *Font) mapAsciiToFont(asciiCode, fontCharX, fontCharY int) {
	tileId := fontCharX + fontCharY*f.sprites.Cols
	f.asciiMap[asciiCode] = tileId
}

//
// Output the mapping for checking
//
func (f Font) Debug() {
	for i := 0; i < 256; i++ {
		fmt.Println("ASCII [", i, "] to idx [", f.asciiMap[i], "]")
	}
}
