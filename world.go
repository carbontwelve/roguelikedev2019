package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

type World struct {
	State
	Terrain      *Terrain
	FovMap       *FovMap
	Entities     *Entities
	seed         int64
	Version      string
	FOVRecompute bool
}

func NewWorld(e *Engine) *World {
	world := &World{
		State:        State{e: e},
		Terrain:      NewTerrain(DungeonWidth, DungeonHeight),
		FovMap:       NewFovMap(DungeonWidth, DungeonHeight),
		Entities:     &Entities{},
		seed:         time.Now().Unix(), //@todo fill this from user input....
		Version:      Version,
		FOVRecompute: true,
	}

	// Generate the terrain and set up the player entity
	world.Entities.Set("player", NewEntity(world.Terrain.Generate(TestTerrainGenerator), '@', PlayerColour))
	world.Terrain.SetExplored(world.Entities.Get("player").position)

	return world
}

func (w World) Draw(dt float32) {
	rl.ClearBackground(UIBackgroundColour)

	// Draw Terrain
	for x := 0; x < w.Terrain.w; x++ {
		for y := 0; y < w.Terrain.h; y++ {
			pos := Position{x, y}
			cell := w.Terrain.Cell(pos)

			if w.FovMap.IsVisible(pos) {
				if cell.T == WallCell {
					w.e.font.Draw(178, pos.Vector2(w.e.font.sprites.TWidth, w.e.font.sprites.THeight), LightWallColour)
				} else {
					w.e.font.Draw('.', pos.Vector2(w.e.font.sprites.TWidth, w.e.font.sprites.THeight), LightGroundColour)
				}
			} else if cell.Explored == true {
				if cell.T == WallCell {
					w.e.font.Draw(178, pos.Vector2(w.e.font.sprites.TWidth, w.e.font.sprites.THeight), DarkWallColour)
				} else {
					w.e.font.Draw('.', pos.Vector2(w.e.font.sprites.TWidth, w.e.font.sprites.THeight), DarkGroundColour)
				}
			}
		}
	}

	// Draw Entities
	for _, entity := range w.Entities.Entities {
		w.e.font.Draw(entity.char, entity.position.Vector2(w.e.font.sprites.TWidth, w.e.font.sprites.THeight), entity.color)
	}
}

func (w *World) Update(dt float32) {
	playerEntity := w.Entities.Get("player")

	if rl.IsKeyDown(rl.KeyUp) && w.Terrain.Cell(playerEntity.position.N()).T == FreeCell {
		playerEntity.Move(0, -1)
		w.FOVRecompute = true
	} else if rl.IsKeyDown(rl.KeyDown) && w.Terrain.Cell(playerEntity.position.S()).T == FreeCell {
		playerEntity.Move(0, 1)
		w.FOVRecompute = true
	} else if rl.IsKeyDown(rl.KeyLeft) && w.Terrain.Cell(playerEntity.position.W()).T == FreeCell {
		playerEntity.Move(-1, 0)
		w.FOVRecompute = true
	} else if rl.IsKeyDown(rl.KeyRight) && w.Terrain.Cell(playerEntity.position.E()).T == FreeCell {
		playerEntity.Move(1, 0)
		w.FOVRecompute = true
	} else if rl.IsKeyDown(rl.KeySpace) {
		w.e.ChangeState(NewWorld(w.e))
	}

	if w.FOVRecompute == true {
		w.Terrain.SetExplored(playerEntity.position)
		w.FovMap.ResetVisibility()
		FOVCircular(w.FovMap, playerEntity.position.X, playerEntity.position.Y, 10, true)

		for _, idx := range w.FovMap.visibleCache {
			w.Terrain.SetExplored(idxtopos(idx))
		}

		w.FOVRecompute = false
	}
}

func (w World) Save(filename string) error {
	return nil // @todo fill this!!
}

func (w World) GetName() string {
	return "World"
}
