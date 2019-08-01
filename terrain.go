package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
	"math/rand"
	"strings"
)

// Define terrain enum, used to determine if a map cell
// is a wall or free
type cellType int

//
// Terrain Generator functions return the players starting Position.
//
type TerrainGeneratorFunc func(*Terrain, *Entities, Config) Position

func TestTerrainGenerator(t *Terrain, e *Entities, c Config) Position {
	t.Fill(FreeCell)
	t.SetCell(Position{30, 22}, WallCell)
	t.SetCell(Position{31, 22}, WallCell)
	t.SetCell(Position{32, 22}, WallCell)
	return Position{30, 23}
}

func TutorialTerrainGenerator(t *Terrain, e *Entities, c Config) Position {
	t.Fill(WallCell)

	var isValid bool
	var rooms []Rect
	var playerStartPosition Position

	numRooms := 0
	maxRooms := c.Get("maxRooms").(int)
	roomMinSize := c.Get("roomMinSize").(int)
	roomMaxSize := c.Get("roomMaxSize").(int)
	maxMonstersPerRoom := c.Get("maxMonstersPerRoom").(int)

	IntnBetween := func(min, max int) int {
		return rand.Intn(max-min) + min
	}

	placeEntities := func(room Rect) {
		monsterCount := rand.Intn(maxMonstersPerRoom)
		if monsterCount == 0 {
			return
		}
		for i := 0; i < monsterCount; i++ {
			//Choose a random location in the room

			location := Position{
				X: IntnBetween(room.x1+1, room.x2-1),
				Y: IntnBetween(room.y1+1, room.y2-1),
			}

			// Check no entity exists at that location
			if !e.FoundAtPosition(location) {
				var monster *Entity
				if rand.Intn(100) < 80 {
					monster = NewEntity(location, 'o', "Orc", rl.Green, true, &MonsterBrain{}, NewFighter(10, 0, 3), Actor)
				} else {
					monster = NewEntity(location, 'T', "Troll", rl.DarkGreen, true, &MonsterBrain{}, NewFighter(16, 1, 4), Actor)
				}
				e.Append(monster)
			}
		}
	}

	for i := 1; i < maxRooms; i++ {
		// random width and height
		w := roomMinSize + rand.Intn(roomMaxSize)
		h := roomMinSize + rand.Intn(roomMaxSize)

		// random Position without going out of the boundaries of the map
		x := rand.Intn(t.w - w - 1)
		y := rand.Intn(t.h - h - 1)

		newRoom := NewRect(x, y, w, h)
		isValid = true

		// Check the new room is valid
		for _, room := range rooms {
			if newRoom.Intersect(room) {
				isValid = false
				break
			}
		}

		if isValid == true {
			// New room doesnt intersect any other rooms and is therefore
			// valid. Continue.
			t.AddRoom(newRoom)

			// center coordinates of new room, will be useful later
			newX, newY := newRoom.Center()

			if numRooms == 0 {
				// this is the first room, where the player starts at
				playerStartPosition = Position{newX, newY}
			} else {
				// all rooms after the first:
				// connect it to the previous room with a tunnel
				// center coordinates of previous room
				prevX, prevY := rooms[numRooms-1].Center()
				if rand.Intn(1) == 1 {
					// first move horizontally, then vertically
					t.AddHTunnel(prevX, newX, prevY)
					t.AddVTunnel(prevY, newY, newX)
				} else {
					// first move vertically, then horizontally
					t.AddVTunnel(prevY, newY, prevX)
					t.AddHTunnel(prevX, newX, newY)
				}
			}

			// finally, append the new room to the list
			rooms = append(rooms, newRoom)
			numRooms++

			placeEntities(newRoom)
		}
	}
	return playerStartPosition
}

const (
	WallCell cellType = iota
	FreeCell
)

//
// A Map is made up of cells, each cell can be either a wall or free.
// If the cell has entered the players FOV then its Explored flag
// will be set to true.
//
type Cell struct {
	T        cellType
	Explored bool
}

type Terrain struct {
	w, h  int
	Cells []Cell
}

func (t Terrain) SetFOVBlocked(fov *FovMap) {
	for x := 0; x < t.w; x++ {
		for y := 0; y < t.h; y++ {
			pos := Position{x, y}
			if t.Cell(pos).T == WallCell {
				fov.SetBlocked(pos, true)
			}
		}
	}
}

//
// Runs the provided terrain generator function against
// this struct.
//
func (t *Terrain) Generate(f TerrainGeneratorFunc, e *Entities, c Config) Position {
	return f(t, e, c)
}

func (t *Terrain) Cell(pos Position) Cell {
	return t.Cells[pos.idx()]
}

func (t *Terrain) SetCell(pos Position, c cellType) {
	t.Cells[pos.idx()].T = c
}

func (t *Terrain) SetExplored(pos Position) {
	t.Cells[pos.idx()].Explored = true
}

func (t Terrain) ToString() string {
	var sb strings.Builder

	for y := 0; y < t.h; y++ {
		for x := 0; x < t.w; x++ {
			if t.Cell(Position{x, y}).T == WallCell {
				sb.WriteString("#")
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

//
// Fill Terrain with tiles
// @todo extend this to use a selection?
//
func (t *Terrain) Fill(c cellType) {
	for x := 0; x < t.w; x++ {
		for y := 0; y < t.h; y++ {
			t.SetCell(Position{x, y}, c)
		}
	}
}

//
// Returns the center cell of the terrain
// @todo bounds check
//
func (t Terrain) Center() Position {
	return Position{int(math.Floor(float64(t.w/2))) - 1, int(math.Floor(float64(t.h/2))) - 1}
}

//
// Returns whether the coordinate is inside the map bounds.
//
func (t Terrain) Inside(pos Position) bool {
	return pos.valid(t.w, t.h)
}

func NewTerrain(w, h int) *Terrain {
	d := &Terrain{w: w, h: h}
	d.Cells = make([]Cell, h*w)
	d.Fill(FreeCell)
	return d
}

//
// Dig a room into the terrain
// @todo set what terrain type to "dig"
//
func (t *Terrain) AddRoom(room Rect) {
	for x := room.x1 + 1; x < room.x2; x++ {
		for y := room.y1 + 1; y < room.y2; y++ {
			pos := Position{x, y}
			if pos.valid(t.w, t.h) {
				t.SetCell(pos, FreeCell)
			}
		}
	}
}

//
// Dig a horizontal tunnel through the terrain
// @todo set what terrain type to "dig"
//
func (t *Terrain) AddHTunnel(x1, x2, y int) {
	for x := int(math.Min(float64(x1), float64(x2))); x < int(math.Max(float64(x1), float64(x2)))+1; x++ {
		pos := Position{x, y}
		if pos.valid(t.w, t.h) {
			t.SetCell(pos, FreeCell)
		}
	}
}

//
// Dig a vertical tunnel through the terrain
// @todo set what terrain type to "dig"
//
func (t *Terrain) AddVTunnel(y1, y2, x int) {
	for y := int(math.Min(float64(y1), float64(y2))); y < int(math.Max(float64(y1), float64(y2)))+1; y++ {
		pos := Position{x, y}
		if pos.valid(t.w, t.h) {
			t.SetCell(pos, FreeCell)
		}
	}
}
