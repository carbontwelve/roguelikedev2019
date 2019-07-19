package main

import "github.com/gen2brain/raylib-go/raylib"

type Entity struct {
	x, y  int
	char  int
	color rl.Color
}

func NewEntity(x, y int, char int, color rl.Color) *Entity {
	return &Entity{
		x:     x,
		y:     y,
		char:  char,
		color: color,
	}
}

func (e *Entity) Move(dx, dy int, gameMap GameMap) {
	if !gameMap.IsBlocked(e.x+dx, e.y+dy) {
		e.x += dx
		e.y += dy
	}
}

func (e Entity) Draw(engine *Engine) {
	position := rl.NewVector2(float32(e.x*engine.font.sprites.THeight), float32(e.y*engine.font.sprites.THeight))
	engine.font.Draw(e.char, position, e.color)
}
