package ui

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// No colour/transparent
var ColourNC = rl.Color{R: 0, G: 0, B: 0, A: 0}

// Game colours, these are the names of colours the game works with and can be
// assigned by LinkColours to different theme colours.
var GameColours map[string]rl.Color

func MapThemeToColours(theme Theme) {
	fmt.Println(fmt.Sprintf("Mapping Theme [%s]", theme.Name))
	GameColours = make(map[string]rl.Color)
	colours := theme.AsRaylibColor()

	// 16 Colour Terminal Palette seems to work nicely with roguelikes.

	// Eight Normal Colours
	GameColours["AnsiBlack"] = colours[0]
	GameColours["AnsiRed"] = colours[1]
	GameColours["AnsiGreen"] = colours[2]
	GameColours["AnsiYellow"] = colours[3]
	GameColours["AnsiBlue"] = colours[4]
	GameColours["AnsiPurple"] = colours[5]
	GameColours["AnsiCyan"] = colours[6]
	GameColours["AnsiWhite"] = colours[7]

	// Eight Light Variants
	GameColours["AnsiGrey"] = colours[8]
	GameColours["AnsiLightRed"] = colours[9]
	GameColours["AnsiLightGreen"] = colours[10]
	GameColours["AnsiLightYellow"] = colours[11]
	GameColours["AnsiLightBlue"] = colours[12]
	GameColours["AnsiLightPurple"] = colours[13]
	GameColours["AnsiLightCyan"] = colours[14]
	GameColours["AnsiLightGrey"] = colours[15]

	// Special Colours
	GameColours["Bg"] = colours[256]
	GameColours["Fg"] = colours[257]
}

func LinkWorkingColourPalette() {
	GameColours["LogNormal"] = GameColours["Fg"]
	GameColours["LogGood"] = GameColours["AnsiGreen"]
	GameColours["LogBad"] = GameColours["AnsiRed"]
	GameColours["LogInfo"] = GameColours["AnsiPurple"]
	GameColours["UiLines"] = GameColours["AnsiLightGrey"]
	GameColours["Player"] = GameColours["AnsiLightYellow"]
	GameColours["Wall"] = GameColours["AnsiGrey"]
	GameColours["WallFOV"] = GameColours["AnsiLightGrey"]
	GameColours["Floor"] = GameColours["AnsiGrey"]
	GameColours["FloorFOV"] = GameColours["AnsiGrey"]
	GameColours["BloodRed"] = GameColours["AnsiLightRed"]
}
