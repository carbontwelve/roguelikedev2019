package ui

import rl "github.com/gen2brain/raylib-go/raylib"

// No colour/transparent
var ColourNC = rl.Color{R: 0, G: 0, B: 0, A: 0}

// @todo add 8 colour pallet?
// @see https://mayccoll.github.io/Gogh/
// Full Colour Pallet
var (
	ColourBlk36White    = rl.Color{255, 255, 255, 255}
	ColourBlk36DarkBlue = rl.Color{18, 23, 61, 255}

	// From https://github.com/devinceble/Elementary-OS-Terminal-Colors/blob/master/content/new.md#flat
	ColourTermFlat8White  = rl.Color{236, 240, 241, 255}
	ColourTermFlat8Black  = rl.Color{45, 62, 80, 255}
	ColourTermFlat8Blue   = rl.Color{43, 129, 180, 255}
	ColourTermFlat8Green  = rl.Color{39, 174, 97, 255}
	ColourTermFlat8Cyan   = rl.Color{22, 160, 134, 255}
	ColourTermFlat8Red    = rl.Color{193, 57, 41, 255}
	ColourTermFlat8Purple = rl.Color{143, 67, 175, 255}
	ColourTermFlat8Yellow = rl.Color{243, 156, 17, 255}
	ColourTermFlat8Grey   = rl.Color{52, 73, 94, 255}

	ColourTermFlat8LightBlue   = rl.Color{51, 152, 220, 255}
	ColourTermFlat8LightGreen  = rl.Color{44, 204, 114, 255}
	ColourTermFlat8LightCyan   = rl.Color{40, 161, 152, 255}
	ColourTermFlat8LightRed    = rl.Color{232, 76, 61, 255}
	ColourTermFlat8LightPurple = rl.Color{154, 89, 179, 255}
	ColourTermFlat8LightYellow = rl.Color{241, 196, 17, 255}
	ColourTermFlat8LightGrey   = rl.Color{190, 195, 199, 255}
)

// 16 Colour ANSI Palette:
// This is loaded by the LoadTheme function and used by LinkColours to
// fill the colours used by the game. This has an effect of allowing
// themes to be switched at run time if needed.
var (
	ColourAnsiWhite       rl.Color
	ColourAnsiBlack       rl.Color
	ColourAnsiBlue        rl.Color
	ColourAnsiGreen       rl.Color
	ColourAnsiCyan        rl.Color
	ColourAnsiRed         rl.Color
	ColourAnsiPurple      rl.Color
	ColourAnsiYellow      rl.Color
	ColourAnsiGrey        rl.Color
	ColourAnsiLightBlue   rl.Color
	ColourAnsiLightGreen  rl.Color
	ColourAnsiLightCyan   rl.Color
	ColourAnsiLightRed    rl.Color
	ColourAnsiLightPurple rl.Color
	ColourAnsiLightYellow rl.Color
	ColourAnsiLightGrey   rl.Color
)

// Game colours, these are the names of colours the game works with and can be
// assigned by LinkColours to different theme colours.
var (
	ColourBg       rl.Color
	ColourFg       rl.Color
	ColourUiLines  rl.Color
	ColourPlayer   rl.Color
	ColourWall     rl.Color
	ColourWallFOV  rl.Color
	ColourFloor    rl.Color
	ColourFloorFOV rl.Color

	ColourBlood rl.Color
)

func LoadTheme() {
	// ...
}

func LinkColours() {
	ColourBg = ColourTermFlat8Black
	ColourFg = ColourTermFlat8White
	ColourUiLines = ColourTermFlat8LightGrey
	ColourPlayer = ColourTermFlat8LightYellow
	ColourWall = ColourTermFlat8Grey
	ColourWallFOV = ColourTermFlat8LightGrey
	ColourFloor = ColourTermFlat8Grey
	ColourFloorFOV = ColourTermFlat8Grey
	ColourBlood = ColourTermFlat8Red
}
