package ui

type BorderStyle struct {
	V, H, NE, SE, SW, NW int
}

type Component struct {
	Name             string
	Width, Height    uint
	OffsetX, OffsetY int
	border           BorderStyle
}

func (c *Component) SetBorderStyle(bs BorderStyle) {
	c.border = bs
}

//
// Constructor
//
func NewComponent(name string, w, h uint, offX, offY int) *Component {
	return &Component{
		Name:    name,
		Width:   w,
		Height:  h,
		OffsetX: offX,
		OffsetY: offY,
	}
}
