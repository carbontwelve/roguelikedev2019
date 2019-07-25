package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
	"strconv"
	"strings"
)

type Fighter struct {
	MaxHP, HP, Defense, Power int
}

func NewFighter(MaxHP, Defense, Power int) *Fighter {
	return &Fighter{MaxHP: MaxHP, HP: MaxHP, Defense: Defense, Power: Power}
}

//
// A Basic Dictionary of Entities
//
type Entities struct {
	Entities map[string]*Entity
	counter  int
}

func (e *Entities) Set(k string, v *Entity) {
	if e.Entities == nil {
		e.Entities = map[string]*Entity{}
		e.counter = 1
	} else {
		e.counter++
	}
	e.Entities[k] = v
}

func (e Entities) FoundAtPosition(pos Position) bool {
	for _, entity := range e.Entities {
		if entity.position.Same(pos) {
			return true
		}
	}
	return false
}

func (e Entities) BlockingAtPosition(pos Position) bool {
	for _, entity := range e.Entities {
		if entity.position.Same(pos) {
			return entity.blocks
		}
	}
	return false
}

func (e Entities) GetBlockingAtPosition(pos Position) *Entity {
	for _, entity := range e.Entities {
		if entity.position.Same(pos) && entity.blocks == true {
			return entity
		}
	}
	return nil
}

func (e *Entities) TickAi(t Terrain, f FovMap) {
	for _, entity := range e.Entities {
		entity.Ai.Tick(entity, e, t, f)
	}
}

func (e *Entities) Append(v *Entity) {
	k := strings.Builder{}
	k.WriteString("entity-")
	k.WriteString(strconv.FormatInt(int64(e.counter), 10))
	e.Entities[k.String()] = v
	e.counter++
}

func (e Entities) Get(k string) *Entity {
	return e.Entities[k]
}

func (e Entities) Delete(k string) {
	delete(e.Entities, k)
}

type Entity struct {
	Ai       Ai
	Fighter  *Fighter
	Name     string
	position Position
	char     int
	color    rl.Color
	blocks   bool
}

func NewEntity(pos Position, char int, name string, color rl.Color, blocking bool, a Ai, f *Fighter) *Entity {
	return &Entity{
		Name:     name,
		position: pos,
		char:     char,
		color:    color,
		blocks:   blocking,
		Fighter:  f,
		Ai:       a,
	}
}

func (e *Entity) Move(dx, dy int) {
	e.position.X += dx
	e.position.Y += dy
}

func (e *Entity) MoveTowards(pos Position, entities Entities, terrain Terrain) {

	dx := pos.X - e.position.X
	dy := pos.Y - e.position.Y
	distance := e.position.Distance(pos)

	dx = int(math.Round(float64(dx / distance)))
	dy = int(math.Round(float64(dy / distance)))

	to := Position{e.position.X + dx, e.position.Y + dy}

	if terrain.Cell(to).T == FreeCell && entities.BlockingAtPosition(to) == false {
		e.Move(dx, dy)
	}
}

func (e Entity) Destination(dx, dy int) Position {
	return Position{e.position.X + dx, e.position.Y + dy}
}

func (e Entity) Draw(engine *Engine) {
	position := rl.NewVector2(float32(e.position.X*engine.font.sprites.THeight), float32(e.position.Y*engine.font.sprites.THeight))
	engine.font.Draw(e.char, position, e.color)
}
