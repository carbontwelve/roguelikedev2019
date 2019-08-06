package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"math"
	"raylibtinkering/position"
	"sort"
	"strconv"
	"strings"
)

type renderOrder int

const (
	RoCorpse renderOrder = iota
	RoItem
	RoActor
)

type entityType int

const (
	EtPlayer entityType = iota
	EtMonster
	EtItem
)

type Fighter struct {
	MaxHP, HP, Defense, Power int
	owner                     *Entity
}

func (f *Fighter) SetOwner(e *Entity) {
	f.owner = e
}

type InteractionResult map[string]interface{}

type InteractionResults struct {
	Results []InteractionResult
}

func (r *InteractionResults) Push(fr InteractionResult) {
	r.Results = append(r.Results, fr)
}

func (r *InteractionResults) Merge(frs InteractionResults) {
	for _, i := range frs.Results {
		r.Push(i)
	}
}

func NewInteractionResults() InteractionResults {
	return InteractionResults{Results: make([]InteractionResult, 0)}
}

func (f *Fighter) TakeDamage(amount int) InteractionResults {
	ret := NewInteractionResults()

	newHP := f.HP - amount
	if newHP < 0 {
		f.HP = 0
	} else {
		f.HP = newHP
	}

	if f.HP == 0 {
		result := make(InteractionResult)
		result["death"] = f.owner
		result["message"] = SimpleMessage{Message: fmt.Sprintf("%s is dead.", f.owner.Name), Colour: rl.Red}
		ret.Push(result)
	}

	return ret
}

func (f *Fighter) Heal(amount int) {
	newHP := f.HP + amount
	if newHP > f.MaxHP {
		f.HP = f.MaxHP
	} else {
		f.HP = newHP
	}
}

func (f Fighter) Attack(target *Entity) InteractionResults {
	ret := NewInteractionResults()
	result := make(InteractionResult)
	damage := f.Power - target.Fighter.Defense

	var msg string

	if damage > 0 {
		if f.owner.Type == EtPlayer {
			msg = fmt.Sprintf("you attack %s for %d hit points.", target.Name, damage)
		} else if f.owner.Type == EtMonster {
			msg = fmt.Sprintf("%s attacks you for %d hit points.", f.owner.Name, damage)
		} else {
			msg = fmt.Sprintf("%s attacks %s for %d hit points.", f.owner.Name, target.Name, damage)
		}

		result["message"] = SimpleMessage{Message: msg, Colour: rl.Orange}
		ret.Push(result)
		ret.Merge(target.Fighter.TakeDamage(damage))
	} else {
		if f.owner.Type == EtPlayer {
			msg = fmt.Sprintf("you attack %s and miss, doing no damage.", target.Name)
		} else if f.owner.Type == EtMonster {
			msg = fmt.Sprintf("%s lurches at you but does no damage.", f.owner.Name)
		} else {
			msg = fmt.Sprintf("%s attacks %s but does no damage.", f.owner.Name, target.Name)
		}

		result["message"] = SimpleMessage{Message: msg, Colour: rl.Orange}
		ret.Push(result)
	}

	return ret
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

func (e Entities) FoundAtPosition(pos position.Position) bool {
	for _, entity := range e.Entities {
		if entity.position.Same(pos) {
			return true
		}
	}
	return false
}

func (e Entities) BlockingAtPosition(pos position.Position) bool {
	for _, entity := range e.Entities {
		if entity.position.Same(pos) {
			return entity.blocks
		}
	}
	return false
}

func (e Entities) GetBlockingAtPosition(pos position.Position) *Entity {
	for _, entity := range e.Entities {
		if entity.Exists() && entity.position.Same(pos) && entity.blocks == true {
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

type RenderOrderEntityList []*Entity

type RenderOrderEntity struct {
	Key   string
	Value renderOrder
}

//
// There has to be a better way of sorting these, but this works and is
// quick enough for now.
//
func (e Entities) SortedByRenderOrder() RenderOrderEntityList {
	ero := make(RenderOrderEntityList, len(e.Entities))
	tmp := make([]RenderOrderEntity, len(e.Entities))

	i := 0
	for k, v := range e.Entities {
		tmp[i] = RenderOrderEntity{k, v.RenderOrder}
		i++
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].Value < tmp[j].Value
	})

	i = 0
	for _, kv := range tmp {
		ero[i] = e.Entities[kv.Key]
		i++
	}

	return ero
}

//func MonsterAction (e *Entity, w *World, ev event)

type Entity struct {
	Brain              Brain
	Fighter            *Fighter
	Name               string
	RenderOrder        renderOrder
	Type               entityType
	position           position.Position
	char               int
	color              rl.Color
	blocks             bool
	TurnActionFunction func(e *Entity, w *World, ev event)
}

func NewEntity(pos position.Position, char int, name string, color rl.Color, blocking bool, b Brain, f *Fighter, renderOrder renderOrder, entityType entityType) *Entity {
	entity := &Entity{
		Name:        name,
		Type:        entityType,
		position:    pos,
		char:        char,
		color:       color,
		blocks:      blocking,
		Fighter:     f,
		Brain:       b,
		RenderOrder: renderOrder,
	}

	entity.Fighter.SetOwner(entity)
	entity.Brain.SetOwner(entity)
	return entity
}

func (e Entity) Exists() bool {
	return e.Fighter.HP > 0
}

func (e *Entity) HandleTurn(w *World, ev event) {
	e.Brain.HandleTurn(w, ev)
}

func (e *Entity) Move(dx, dy int) {
	e.position.X += dx
	e.position.Y += dy
}

func (e *Entity) MoveTo(pos position.Position) {
	e.position.X = pos.X
	e.position.Y = pos.Y
}

func (e Entity) NextMove(dx, dy int) position.Position {
	return position.Position{e.position.X + dx, e.position.Y + dy}
}

func (e *Entity) MoveTowards(pos position.Position, entities Entities, terrain Terrain) {

	dx := pos.X - e.position.X
	dy := pos.Y - e.position.Y
	distance := e.position.Distance(pos)

	dx = int(math.Round(float64(dx / distance)))
	dy = int(math.Round(float64(dy / distance)))

	to := position.Position{e.position.X + dx, e.position.Y + dy}

	if terrain.Cell(to).T == FreeCell && entities.BlockingAtPosition(to) == false {
		e.Move(dx, dy)
	}
}

func (e Entity) Destination(dx, dy int) position.Position {
	return position.Position{e.position.X + dx, e.position.Y + dy}
}

func (e Entity) Draw(engine *Engine) {
	position := rl.NewVector2(float32(e.position.X*engine.font.sprites.TileHeight), float32(e.position.Y*engine.font.sprites.TileHeight))
	engine.font.Draw(e.char, position, e.color)
}
