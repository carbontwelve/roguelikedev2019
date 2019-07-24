package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Entity struct {
	position Position
	char     int
	color    rl.Color
}

func NewEntity(pos Position, char int, color rl.Color) *Entity {
	return &Entity{
		position: pos,
		char:     char,
		color:    color,
	}
}

func (e *Entity) Move(dx, dy int, gameMap *GameMap) {
	if !gameMap.IsBlocked(e.position.X+dx, e.position.Y+dy) {
		e.position.X += dx
		e.position.Y += dy
		gameMap.FOVRecompute = true // Need to recalculate FOV when player has moved
	}
}

func (e Entity) Draw(engine *Engine) {
	position := rl.NewVector2(float32(e.position.X*engine.font.sprites.THeight), float32(e.position.Y*engine.font.sprites.THeight))
	engine.font.Draw(e.char, position, e.color)
}
