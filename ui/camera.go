package ui

import (
	"fmt"
	"raylibtinkering/position"
)

type Camera struct {
	offset      position.Position
	origin      ComponentI
	destination ComponentI
}

func (c *Camera) FollowTarget(p position.Position) {
	c.offset.X = p.X - int(c.destination.GetWidth()/2)
	c.offset.Y = p.Y - int(c.destination.GetHeight()/2)

	if c.offset.X < 0 {
		c.offset.X = 0
	}
	if c.offset.Y < 0 {
		c.offset.Y = 0
	}

	for Y := 0; Y < int(c.destination.GetHeight()); Y++ {
		for X := 0; X < int(c.destination.GetWidth()); X++ {
			offX := X + c.offset.X
			offY := Y + c.offset.Y

			if offY >= int(c.origin.GetHeight()) {
				offY = int(c.origin.GetHeight()) - 1
			}

			if offX >= int(c.origin.GetWidth()) { //80 wide = 0 - 79
				offX = int(c.origin.GetWidth()) - 1
			}

			if uint(offX) > c.origin.GetWidth() {
				panic(fmt.Sprintf("The Offset X (%d) is greater than the destination width (%d)", offX, c.origin.GetWidth()))
			}

			if uint(offY) > c.origin.GetHeight() {
				panic(fmt.Sprintf("The Offset Y (%d) is greater than the destination height (%d)", offY, c.origin.GetHeight()))
			}

			cell := c.origin.GetCells()[position.Position{X: offX, Y: offY}]

			if cell == nil {
				panic(fmt.Sprintf("The Cell at (%d,%d) is nil", offX, offY))
			}

			c.destination.SetChar(cell.char, position.Position{X: X, Y: Y}, cell.fg, cell.bg)
		}
	}
}

func (c Camera) Debug() {
	fmt.Println(fmt.Sprintf("Camera offset (%d,%d), Viewport (%d x %d), Max (%d, %d)", c.offset.X, c.offset.Y, c.destination.GetWidth(), c.destination.GetHeight(), c.origin.GetWidth(), c.origin.GetHeight()))
}

func NewCamera(origin, destination ComponentI) *Camera {
	camera := &Camera{
		origin:      origin,
		destination: destination,
		offset:      position.Position{X: 0, Y: 0},
	}

	origin.SetCamera(camera)
	return camera
}
