package ui

import rl "github.com/gen2brain/raylib-go/raylib"

//type ColourTheme struct {
//	Name    string
//	Colours map[string]rl.Color
//}

//var CurrentTheme ColourTheme

// From https://www.gnome-look.org/p/1112201/
//var ThemeAmberTerm = ColourTheme{
//	Name: "AmberTerm",
//	Colours: map[string]rl.Color{
//		"ColourAnsiBlack":     rl.Color{15, 14, 13, 255},    // 01
//		"ColourAnsiRed":       rl.Color{140, 41, 32, 255},   // 02
//		"ColourAnsiGreen":     rl.Color{168, 83, 52, 255},   // 03
//		"ColourAnsiYellow":    rl.Color{140, 55, 21, 255},   // 04
//		"ColourAnsiBlue":      rl.Color{168, 67, 0, 255},    // 05
//		"ColourAnsiPurple":    rl.Color{150, 92, 21, 255},   // 06
//		"ColourAnsiCyan":      rl.Color{168, 126, 0, 255},   // 07
//		"ColourAnsiLightGrey": rl.Color{161, 150, 133, 255}, // 08
//
//		"ColourAnsiGrey":        rl.Color{26, 24, 21, 255},    // 09
//		"ColourAnsiLightRed":    rl.Color{227, 66, 52, 255},   // 10
//		"ColourAnsiLightGreen":  rl.Color{255, 127, 80, 255},  // 11
//		"ColourAnsiLightYellow": rl.Color{226, 88, 34, 255},   // 12
//		"ColourAnsiLightBlue":   rl.Color{255, 102, 0, 255},   // 13
//		"ColourAnsiLightPurple": rl.Color{237, 145, 33, 255},  // 14
//		"ColourAnsiLightCyan":   rl.Color{255, 191, 0, 255},   // 15
//		"ColourAnsiWhite":       rl.Color{247, 231, 206, 255}, // 16
//	},
//}

// From https://github.com/devinceble/Elementary-OS-Terminal-Colors/blob/master/content/new.md#flat
//var ThemeFlat8 = ColourTheme{
//	Name: "Flat",
//	Colours: map[string]rl.Color{
//		"ColourAnsiBlack":     rl.Color{45, 62, 80, 255},    // 01
//		"ColourAnsiRed":       rl.Color{193, 57, 41, 255},   // 02
//		"ColourAnsiGreen":     rl.Color{39, 174, 97, 255},   // 03
//		"ColourAnsiYellow":    rl.Color{243, 156, 17, 255},  // 04
//		"ColourAnsiBlue":      rl.Color{43, 129, 180, 255},  // 05
//		"ColourAnsiPurple":    rl.Color{143, 67, 175, 255},  // 06
//		"ColourAnsiCyan":      rl.Color{22, 160, 134, 255},  // 07
//		"ColourAnsiLightGrey": rl.Color{190, 195, 199, 255}, // 08
//
//		"ColourAnsiGrey":        rl.Color{52, 73, 94, 255},    // 09
//		"ColourAnsiLightRed":    rl.Color{232, 76, 61, 255},   // 10
//		"ColourAnsiLightGreen":  rl.Color{44, 204, 114, 255},  // 11
//		"ColourAnsiLightYellow": rl.Color{241, 196, 17, 255},  // 12
//		"ColourAnsiLightBlue":   rl.Color{51, 152, 220, 255},  // 13
//		"ColourAnsiLightPurple": rl.Color{154, 89, 179, 255},  // 14
//		"ColourAnsiLightCyan":   rl.Color{40, 161, 152, 255},  // 15
//		"ColourAnsiWhite":       rl.Color{236, 240, 241, 255}, // 16
//	},
//}

// From https://raw.githubusercontent.com/Mayccoll/Gogh/master/images/themes/fishtank.jpg
//var ThemeFishTank8 = ColourTheme{
//	Name: "Fish",
//	Colours: map[string]rl.Color{
//		"ColourAnsiBlack":     rl.Color{4, 7, 60, 255},      // 01
//		"ColourAnsiRed":       rl.Color{198, 0, 75, 255},    // 02
//		"ColourAnsiGreen":     rl.Color{172, 241, 88, 255},  // 03
//		"ColourAnsiYellow":    rl.Color{255, 205, 94, 255},  // 04
//		"ColourAnsiBlue":      rl.Color{81, 95, 184, 255},   // 05
//		"ColourAnsiPurple":    rl.Color{151, 110, 130, 255}, // 06
//		"ColourAnsiCyan":      rl.Color{149, 135, 98, 255},  // 07
//		"ColourAnsiLightGrey": rl.Color{236, 240, 252, 255}, // 08
//
//		"ColourAnsiGrey":        rl.Color{108, 91, 48, 255},   // 09
//		"ColourAnsiLightRed":    rl.Color{218, 75, 137, 255},  // 10
//		"ColourAnsiLightGreen":  rl.Color{219, 255, 168, 255}, // 11
//		"ColourAnsiLightYellow": rl.Color{252, 230, 170, 255}, // 12
//		"ColourAnsiLightBlue":   rl.Color{178, 190, 250, 255}, // 13
//		"ColourAnsiLightPurple": rl.Color{253, 165, 205, 255}, // 14
//		"ColourAnsiLightCyan":   rl.Color{165, 190, 135, 255}, // 15
//		"ColourAnsiWhite":       rl.Color{246, 255, 236, 255}, // 16
//	},
//}

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
