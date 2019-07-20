package main

/**
 * This source code originated from the ROG project and is
 * Licensed BSD 2-clause
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
func fovCircularCastRay(fov *GameMap, xo, yo, xd, yd, r2 int, walls bool) {
	curx := xo
	cury := yo
	in := false
	blocked := false
	if fov.Inside(curx, cury) {
		in = true
		fov.At(curx, cury).inFOV = true
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
		if fov.Inside(curx, cury) {
			in = true
			if !blocked && fov.At(curx, cury).blocked {
				blocked = true
			} else if blocked {
				break
			}
			if walls || !blocked {
				fov.At(curx, cury).inFOV = true
			}
		} else if in {
			break
		}
	}
}

func fovCircularPostProc(fov *GameMap, x0, y0, x1, y1, dx, dy int) {
	for cx := x0; cx <= x1; cx++ {
		for cy := y0; cy <= y1; cy++ {
			x2 := cx + dx
			y2 := cy + dy
			if fov.Inside(cx, cy) && fov.At(cx, cy).inFOV && !fov.At(cx, cy).blocked {
				if x2 >= x0 && x2 <= x1 {
					if fov.Inside(x2, cy) && fov.At(x2, cy).blocked {
						fov.At(x2, cy).inFOV = true
					}
				}
				if y2 >= y0 && y2 <= y1 {
					if fov.Inside(cx, y2) && fov.At(cx, y2).blocked {
						fov.At(cx, y2).inFOV = true
					}
				}
				if x2 >= x0 && x2 <= x1 && y2 >= y0 && y2 <= y1 {
					if fov.Inside(x2, y2) && fov.At(x2, y2).blocked {
						fov.At(x2, y2).inFOV = true
					}
				}
			}
		}
	}
}

//
// FOVCircular Raycasts out from the vantage in a circle.
//
func FOVCircular(fov *GameMap, x, y, r int, walls bool) {
	xo := 0
	yo := 0
	xmin := 0
	ymin := 0
	xmax := fov.MWidth
	ymax := fov.MHeight
	r2 := r * r
	if r > 0 {
		xmin = max(0, x-r)
		ymin = max(0, y-r)
		xmax = min(fov.MWidth, x+r+1)
		ymax = min(fov.MHeight, y+r+1)
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
