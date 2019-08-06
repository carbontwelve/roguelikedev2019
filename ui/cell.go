package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"raylibtinkering/position"
)

//
// Cell:
//
// The Cell struct is a container for an area found in a Component draw buffer. It acts
// as storage for what is to be displayed on the screen at a given moment.
//

type Cell struct {
	char   int
	fg, bg rl.Color
	x, y   uint // Offset x,y based upon Viewport x + cCell Position in the cells map
}

func (c Cell) GetDrawPosition() position.Position {
	return position.Position{X: int(c.x), Y: int(c.y)}
}
