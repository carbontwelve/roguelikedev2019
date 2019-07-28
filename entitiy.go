package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"math"
	"strconv"
	"strings"
)

type Fighter struct {
	MaxHP, HP, Defense, Power int
	owner                     *Entity
}

func (f *Fighter) SetOwner(e *Entity) {
	f.owner = e
}

func (f *Fighter) TakeDamage(amount int) {
	newHP := f.HP - amount
	if newHP < 0 {
		f.HP = 0
	} else {
		f.HP = newHP
	}
}

func (f *Fighter) Heal(amount int) {
	newHP := f.HP + amount
	if newHP > f.MaxHP {
		f.HP = f.MaxHP
	} else {
		f.HP = newHP
	}
}

func (f Fighter) Attack(target *Entity) {
	damage := f.Power - target.Fighter.Defense

	if damage > 0 {
		target.Fighter.TakeDamage(damage)
		fmt.Println(fmt.Sprintf("%s attacks %s for %d hit points.", f.owner.Name, target.Name, damage))
	} else {
		fmt.Println(fmt.Sprintf("%s attacks %s but does no damage.", f.owner.Name, target.Name))
	}
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

//func MonsterAction (e *Entity, w *World, ev event)

type Entity struct {
	Brain              Brain
	Fighter            *Fighter
	Name               string
	Exists             bool
	position           Position
	char               int
	color              rl.Color
	blocks             bool
	TurnActionFunction func(e *Entity, w *World, ev event)
}

func NewEntity(pos Position, char int, name string, color rl.Color, blocking bool, b Brain, f *Fighter) *Entity {
	entity := &Entity{
		Name:     name,
		Exists:   true,
		position: pos,
		char:     char,
		color:    color,
		blocks:   blocking,
		Fighter:  f,
		Brain:    b,
	}

	entity.Fighter.SetOwner(entity)
	entity.Brain.SetOwner(entity)
	return entity
}

func (e *Entity) HandleTurn(w *World, ev event) {
	e.Brain.HandleTurn(w, ev)
}

func (e *Entity) Move(dx, dy int) {
	e.position.X += dx
	e.position.Y += dy
}

func (e *Entity) MoveTo(pos Position) {
	e.position.X = pos.X
	e.position.Y = pos.Y
}

func (e Entity) NextMove(dx, dy int) Position {
	return Position{e.position.X + dx, e.position.Y + dy}
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
