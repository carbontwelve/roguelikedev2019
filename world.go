package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

type World struct {
	State
	Terrain  *Terrain
	Entities *Entities
	seed     int64
	Version  string
}

func NewWorld(e *Engine) *World {
	world := &World{
		State:    State{e: e},
		Terrain:  NewTerrain(DungeonWidth, DungeonHeight),
		Entities: &Entities{},
		seed:     time.Now().Unix(), //@todo fill this from user input....
		Version:  Version,
	}

	// Generate the terrain and set up the player entity
	world.Entities.Set("player", NewEntity(world.Terrain.Generate(TestTerrainGenerator), '@', PlayerColour))

	return world
}

func (w World) Draw(dt float32) {
	rl.ClearBackground(UIBackgroundColour)

	player := w.Entities.Get("player")
	w.e.font.Draw(player.char, player.position.Vector2(10), player.color)
}

func (w World) Update(dt float32) {
	playerEntity := w.Entities.Get("player")

	if rl.IsKeyDown(rl.KeyUp) && w.Terrain.Cell(playerEntity.position.N()).T == FreeCell {
		playerEntity.Move(0, -1)
	} else if rl.IsKeyDown(rl.KeyDown) && w.Terrain.Cell(playerEntity.position.S()).T == FreeCell {
		playerEntity.Move(0, 1)
	} else if rl.IsKeyDown(rl.KeyLeft) && w.Terrain.Cell(playerEntity.position.W()).T == FreeCell {
		playerEntity.Move(-1, 0)
	} else if rl.IsKeyDown(rl.KeyRight) && w.Terrain.Cell(playerEntity.position.E()).T == FreeCell {
		playerEntity.Move(1, 0)
	} else if rl.IsKeyDown(rl.KeySpace) {
		w.e.ChangeState(NewWorld(w.e))
	}

	// if s.GameMap.FOVRecompute == true {
	// s.GameMap.CalculateFov(playerEntity.x, playerEntity.y, 10, true, FOVCircular)
	// }
}

func (w World) Save(filename string) error {
	return nil // @todo fill this!!
}

func (w World) GetName() string {
	return "World"
}
