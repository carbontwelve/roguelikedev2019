package main

import (
	"container/heap"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

type World struct {
	State
	Terrain      *Terrain
	FovMap       *FovMap
	Entities     *Entities
	Events       *eventQueue
	EventIndex   int
	Depth        int
	seed         int64
	Version      string
	FOVRecompute bool
	FOVAlgo      FOVAlgo
}

func (w *World) InitWorld() {
	genConfig := Config{map[string]interface{}{
		"roomMaxSize":        10,
		"roomMinSize":        6,
		"maxRooms":           30,
		"maxMonstersPerRoom": 3,
	}}

	// Generate the terrain and set up the player entity
	w.Entities.Set("player", NewEntity(w.Terrain.Generate(TutorialTerrainGenerator, w.Entities, genConfig), '@', "Player", PlayerColour, true, NullAi{}, NewFighter(30, 2, 5)))
	w.Terrain.SetExplored(w.Entities.Get("player").position)

	// Set blocked tiles from terrain
	w.Terrain.SetFOVBlocked(w.FovMap)

	if w.Depth == 1 {
		w.Events = &eventQueue{}
		heap.Init(w.Events)
	} else {
		// g.CleanEvents() @todo learn what this does... see https://github.com/anaseto/boohu/blob/master/game.go#L654
	}

	for name, _ := range w.Entities.Entities {
		if name == "player" {
			w.PushEvent(&simpleEvent{ERank: 0, EAction: PlayerTurn})
		} else {
			// @todo look at https://github.com/anaseto/boohu/blob/master/game.go#L641 and figure out ERank
			w.PushEvent(&monsterEvent{ERank: 10, EAction: MonsterTurn, NMons: name})
		}
	}
}

func NewWorld(e *Engine) *World {
	world := &World{
		State:        State{e: e},
		Terrain:      NewTerrain(DungeonWidth, DungeonHeight),
		FovMap:       NewFovMap(DungeonWidth, DungeonHeight),
		Entities:     &Entities{Entities: make(map[string]*Entity)},
		Depth:        1,
		seed:         time.Now().Unix(), //@todo fill this from user input....
		Version:      Version,
		FOVRecompute: true,
		FOVAlgo:      FOVCircular,
	}
	world.InitWorld()
	return world
}

func (w *World) PushEvent(ev event) {
	iev := iEvent{Event: ev, Index: w.EventIndex}
	w.EventIndex++
	heap.Push(w.Events, iev)
}

func (w *World) PushAgainEvent(ev event) {
	iev := iEvent{Event: ev, Index: 0}
	heap.Push(w.Events, iev)
}

func (w *World) PopIEvent() iEvent {
	iev := heap.Pop(w.Events).(iEvent)
	return iev
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
		if w.FovMap.IsVisible(entity.position) {
			w.e.font.Draw(entity.char, entity.position.Vector2(w.e.font.sprites.TWidth, w.e.font.sprites.THeight), entity.color)
		}
	}

	rl.DrawText(fmt.Sprintf("HP: %d/%d", w.Entities.Get("player").Fighter.HP, w.Entities.Get("player").Fighter.MaxHP), int32(rl.GetScreenWidth()-100), 20, 10, PlayerColour)
}

func (w *World) Update(dt float32) {
	playerEntity := w.Entities.Get("player")

	//attackTarget := func(at Position) bool {
	//	target := w.Entities.GetBlockingAtPosition(at)
	//	if target == nil {
	//		return false
	//	}
	//	playerEntity.Fighter.Attack(target)
	//	// w.turnState = AiTurn
	//	return true
	//}

	//if w.turnState == PlayerTurn {
	//	if rl.IsKeyDown(rl.KeyUp) && w.Terrain.Cell(playerEntity.position.N()).T == FreeCell && !attackTarget(playerEntity.position.N()) {
	//		playerEntity.Move(0, -1)
	//		w.FOVRecompute = true
	//		w.turnState = AiTurn
	//	} else if rl.IsKeyDown(rl.KeyDown) && w.Terrain.Cell(playerEntity.position.S()).T == FreeCell && !attackTarget(playerEntity.position.S()) {
	//		playerEntity.Move(0, 1)
	//		w.FOVRecompute = true
	//		w.turnState = AiTurn
	//	} else if rl.IsKeyDown(rl.KeyLeft) && w.Terrain.Cell(playerEntity.position.W()).T == FreeCell && !attackTarget(playerEntity.position.W()) {
	//		playerEntity.Move(-1, 0)
	//		w.FOVRecompute = true
	//		w.turnState = AiTurn
	//	} else if rl.IsKeyDown(rl.KeyRight) && w.Terrain.Cell(playerEntity.position.E()).T == FreeCell && !attackTarget(playerEntity.position.E()) {
	//		playerEntity.Move(1, 0)
	//		w.FOVRecompute = true
	//		w.turnState = AiTurn
	//	} else if rl.IsKeyDown(rl.KeySpace) {
	//		w.e.ChangeState(NewWorld(w.e))
	//	}
	//} else if w.turnState == AiTurn {
	//	w.Entities.TickAi(*w.Terrain, *w.FovMap)
	//	w.turnState = PlayerTurn
	//}

	if w.FOVRecompute == true {
		w.Terrain.SetExplored(playerEntity.position)
		w.FovMap.ResetVisibility()
		w.FOVAlgo(w.FovMap, playerEntity.position.X, playerEntity.position.Y, 10, true)

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
