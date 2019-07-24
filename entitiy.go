package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

//
// A Basic Dictionary of Entities
//
type Entities struct {
	Entities map[string]*Entity
}

func (e *Entities) Set(k string, v *Entity) {
	if e.Entities == nil {
		e.Entities = map[string]*Entity{}
	}
	e.Entities[k] = v
}

func (e Entities) Get(k string) *Entity {
	return e.Entities[k]
}

func (e Entities) Delete(k string) {
	delete(e.Entities, k)
}

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

func (e *Entity) Move(dx, dy int) {
	e.position.X += dx
	e.position.Y += dy
}

func (e Entity) Draw(engine *Engine) {
	position := rl.NewVector2(float32(e.position.X*engine.font.sprites.THeight), float32(e.position.Y*engine.font.sprites.THeight))
	engine.font.Draw(e.char, position, e.color)
}
