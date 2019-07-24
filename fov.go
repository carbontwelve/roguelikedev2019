package main

/**
 * This source code originated from the ROG project and is
 * Licensed BSD 2-clause. I have made modifications in order to make
 * it fit within the scope of this project.
 *
 * Copyright (c) 2012 Joseph Hager. All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *    * Redistributions of source code must retain the above copyright
 * notice, this list of conditions and the following disclaimer.
 *    * Redistributions in binary form must reproduce the above
 * copyright notice, this list of conditions and the following disclaimer
 * in the documentation and/or other materials provided with the
 * distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 * @see https://github.com/ajhager/rog/blob/master/fov.go
 */

type FovMap struct {
	w, h             int
	blocked, visible []bool
	visibleCache     []int
}

func (f *FovMap) IsBlocked(pos Position) bool {
	return f.blocked[pos.idx()]
}

func (f *FovMap) SetBlocked(pos Position, c bool) {
	f.blocked[pos.idx()] = c
}

func (f *FovMap) IsVisible(pos Position) bool {
	return f.visible[pos.idx()]
}

func (f *FovMap) SetVisible(pos Position) {
	f.visible[pos.idx()] = true
	f.visibleCache = append(f.visibleCache, pos.idx())
}

func (f *FovMap) ResetVisibility() {
	for x := 0; x < f.w; x++ {
		for y := 0; y < f.h; y++ {
			f.visible[Position{x, y}.idx()] = false
		}
	}
	f.visibleCache = []int{}
}

//
// Returns whether the coordinate is inside the map bounds.
//
func (f FovMap) Inside(pos Position) bool {
	return pos.valid(f.w, f.h)
}

func NewFovMap(w, h int) *FovMap {
	return &FovMap{
		w:            w,
		h:            h,
		blocked:      make([]bool, w*h),
		visible:      make([]bool, w*h),
		visibleCache: make([]int, 0),
	}
}

//
// FOVAlgo takes a FOVMap x,y vantage, radius of the view, whether to include
// walls and then marks in the map which cells are viewable.
//
type FOVAlgo func(*GameMap, int, int, int, bool)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Circular Raycasting
func fovCircularCastRay(fov *FovMap, xo, yo, xd, yd, r2 int, walls bool) {
	curx := xo
	cury := yo
	in := false
	blocked := false
	if fov.Inside(Position{curx, cury}) {
		in = true
		fov.SetVisible(Position{curx, cury})
	}
	for _, p := range Line(xo, yo, xd, yd) {
		curx = p.X
		cury = p.Y
		if r2 > 0 {
			curRadius := (curx-xo)*(curx-xo) + (cury-yo)*(cury-yo)
			if curRadius > r2 {
				break
			}
		}
		if fov.Inside(Position{curx, cury}) {
			in = true
			if !blocked && fov.IsBlocked(Position{curx, cury}) {
				blocked = true
			} else if blocked {
				break
			}
			if walls || !blocked {
				fov.SetVisible(Position{curx, cury})
			}
		} else if in {
			break
		}
	}
}

func fovCircularPostProc(fov *FovMap, x0, y0, x1, y1, dx, dy int) {
	for cx := x0; cx <= x1; cx++ {
		for cy := y0; cy <= y1; cy++ {
			x2 := cx + dx
			y2 := cy + dy
			if fov.Inside(Position{cx, cy}) && fov.IsVisible(Position{cx, cy}) && !fov.IsBlocked(Position{cx, cy}) {
				if x2 >= x0 && x2 <= x1 {
					if fov.Inside(Position{x2, cy}) && fov.IsBlocked(Position{x2, cy}) {
						fov.SetVisible(Position{x2, cy})
					}
				}
				if y2 >= y0 && y2 <= y1 {
					if fov.Inside(Position{cx, y2}) && fov.IsBlocked(Position{cx, y2}) {
						fov.SetVisible(Position{cx, y2})
					}
				}
				if x2 >= x0 && x2 <= x1 && y2 >= y0 && y2 <= y1 {
					if fov.Inside(Position{x2, y2}) && fov.IsBlocked(Position{x2, y2}) {
						fov.SetVisible(Position{x2, y2})
					}
				}
			}
		}
	}
}

//
// FOVCircular Raycasts out from the vantage in a circle.
//
func FOVCircular(fov *FovMap, x, y, r int, walls bool) {
	xo := 0
	yo := 0
	xmin := 0
	ymin := 0
	xmax := fov.w
	ymax := fov.h
	r2 := r * r
	if r > 0 {
		xmin = max(0, x-r)
		ymin = max(0, y-r)
		xmax = min(fov.w, x+r+1)
		ymax = min(fov.h, y+r+1)
	}
	xo = xmin
	yo = ymin
	for xo < xmax {
		fovCircularCastRay(fov, x, y, xo, yo, r2, walls)
		xo++
	}
	xo = xmax - 1
	yo = ymin + 1
	for yo < ymax {
		fovCircularCastRay(fov, x, y, xo, yo, r2, walls)
		yo++
	}
	xo = xmax - 2
	yo = ymax - 1
	for xo >= 0 {
		fovCircularCastRay(fov, x, y, xo, yo, r2, walls)
		xo--
	}
	xo = xmin
	yo = ymax - 2
	for yo > 0 {
		fovCircularCastRay(fov, x, y, xo, yo, r2, walls)
		yo--
	}
	if walls {
		fovCircularPostProc(fov, xmin, ymin, x, y, -1, -1)
		fovCircularPostProc(fov, x, ymin, xmax-1, y, 1, -1)
		fovCircularPostProc(fov, xmin, y, x, ymax-1, -1, 1)
		fovCircularPostProc(fov, x, y, xmax-1, ymax-1, 1, 1)
	}
}
