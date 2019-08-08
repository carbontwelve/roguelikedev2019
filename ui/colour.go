package ui

import rl "github.com/gen2brain/raylib-go/raylib"

// No colour/transparent
var ColourNC = rl.Color{R: 0, G: 0, B: 0, A: 0}

// 16 Colour ANSI Palette:
// This is loaded by the LoadTheme function and used by LinkColours to
// fill the colours used by the game. This has an effect of allowing
// themes to be switched at run time if needed.
var (
	ColourForeground      rl.Color
	ColourBackground      rl.Color
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

func MapThemeToColours(theme Theme) {
	colours := theme.AsRaylibColor()

	// Eight Normal Colours
	ColourAnsiBlack = colours[0]
	ColourAnsiRed = colours[1]
	ColourAnsiGreen = colours[2]
	ColourAnsiYellow = colours[3]
	ColourAnsiBlue = colours[4]
	ColourAnsiPurple = colours[5]
	ColourAnsiCyan = colours[6]
	ColourAnsiWhite = colours[7]

	// Eight Light Variants
	ColourAnsiGrey = colours[8]
	ColourAnsiLightRed = colours[9]
	ColourAnsiLightGreen = colours[10]
	ColourAnsiLightYellow = colours[11]
	ColourAnsiLightBlue = colours[12]
	ColourAnsiLightPurple = colours[13]
	ColourAnsiLightCyan = colours[14]
	ColourAnsiLightGrey = colours[15]

	// Special Colours
	ColourBackground = colours[256]
	ColourForeground = colours[257]
}

func LinkWorkingColourPalette() {
	ColourBg = ColourBackground
	ColourFg = ColourForeground
	ColourUiLines = ColourAnsiLightGrey
	ColourPlayer = ColourAnsiLightYellow
	ColourWall = ColourAnsiGrey
	ColourWallFOV = ColourAnsiLightGrey
	ColourFloor = ColourAnsiGrey
	ColourFloorFOV = ColourAnsiGrey
	ColourBlood = ColourAnsiLightRed
}
