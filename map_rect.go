package main

type Rect struct {
	x1, y1, x2, y2 int
}

func NewRect(x, y, w, h int) Rect {
	return Rect{
		x1: x,
		y1: y,
		x2: x + w,
		y2: y + h,
	}
}

//
// Returns the center point of this rectangle
//
func (r Rect) Center() (int, int) {
	return int((r.x1 + r.x2) / 2), int((r.y1 + r.y2) / 2)
}

//
// returns true if this rectangle intersects with another one
//
func (r Rect) Intersect(other Rect) bool {
	return r.x1 <= other.x2 && r.x2 >= other.x1 && r.y1 <= other.y2 && r.y2 >= other.y1
}
