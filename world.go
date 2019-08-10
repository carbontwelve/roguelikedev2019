package main

import (
	"container/heap"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"raylibtinkering/position"
	"raylibtinkering/state"
	"raylibtinkering/ui"
	"time"
)

type World struct {
	state.State
	NextTurnMove position.Position
	MessageLog   *MessageLog
	Terrain      *Terrain
	FovMap       *FovMap
	Entities     *Entities
	Events       *eventQueue
	Ev           event
	EventIndex   int
	Depth        int
	Turn         int
	seed         int64
	Version      string
	FOVRecompute bool
	FOVAlgo      FOVAlgo
	inputDelay   float32
	MouseX       int
	MouseY       int
	MouseHover   bool
	Camera       *ui.Camera
}

func (w *World) InitWorld() {
	genConfig := Config{map[string]interface{}{
		"roomMaxSize":        10,
		"roomMinSize":        6,
		"maxRooms":           30,
		"maxMonstersPerRoom": 3,
	}}

	// Generate the terrain and set up the player entity
	w.Entities.Set("player", NewEntity(w.Terrain.Generate(TutorialTerrainGenerator, w.Entities, genConfig), '@', "Player", "Player", true, &PlayerBrain{}, NewFighter(30, 2, 5), RoActor, EtPlayer))
	w.Terrain.SetExplored(w.Entities.Get("player").position)

	// Set blocked tiles from terrain
	w.Terrain.SetFOVBlocked(w.FovMap)

	// Set camera initial position
	w.Camera.FollowTarget(w.Entities.Get("player").position)

	if w.Depth == 1 {
		w.Events = &eventQueue{}
		heap.Init(w.Events)
	} else {
		// g.CleanEvents() @todo learn what this does... see https://github.com/anaseto/boohu/blob/master/game.go#L654
	}

	for name := range w.Entities.Entities {
		if name == "player" {
			w.PushEvent(&simpleEvent{ERank: 0, EAction: PlayerTurn})
			w.PushEvent(&simpleEvent{ERank: 500, EAction: HealPlayer}) // heal the player every 50 turns
		} else {
			// @todo look at https://github.com/anaseto/boohu/blob/master/game.go#L641 and figure out ERank
			w.PushEvent(&monsterEvent{ERank: 10, EAction: MonsterTurn, NMons: name})
		}
	}
}

// @todo move this to a on state change function as we may not want to reset on World construction...
// @todo refactor DungeonWidth/Height to something more practical
func (w *World) Pushed(owner *state.Engine) error {
	owner.Screen.Reset()

	msgLog := NewMessageLog("MessageLog", position.DungeonWidth, 6, 0, position.DungeonHeight-6)

	owner.Screen.Set(ui.NewComponent("Viewport", position.DungeonWidth-24, position.DungeonHeight-5, 0, 0, true), 10)
	owner.Screen.Set(msgLog, 20)
	owner.Screen.Set(ui.NewComponent("Statistics", 25, position.DungeonHeight-5, position.DungeonWidth-25, 0, true), 25)
	owner.Screen.Set(ui.NewComponent("Map", position.DungeonWidth, position.DungeonHeight, 0, 0, false), 9999)
	owner.Screen.Set(ui.NewComponent("Mouse", position.DungeonWidth, position.DungeonHeight, 0, 0, true), 9999)
	owner.Screen.Get("MessageLog").SetBorderStyle(ui.SingleWallBorder)
	owner.Screen.Get("Statistics").SetBorderStyle(ui.SingleWallBorder)

	camera := ui.NewCamera(owner.Screen.Get("Map"), owner.Screen.Get("Viewport"))

	w.Owner = owner
	w.Camera = camera
	w.MessageLog = msgLog

	w.InitWorld()
	return nil
}

func NewWorld() *World {
	world := &World{
		NextTurnMove: position.Position{0, 0},
		State:        state.State{Quit: false},
		Terrain:      NewTerrain(position.DungeonWidth, position.DungeonHeight),
		FovMap:       NewFovMap(position.DungeonWidth, position.DungeonHeight),
		Entities:     &Entities{Entities: make(map[string]*Entity)},
		Depth:        1,
		seed:         time.Now().Unix(), //@todo fill this from user input....
		Version:      Version,
		FOVRecompute: true,
		FOVAlgo:      FOVCircular,
		inputDelay:   0.11,
		//MessageLog:   msgLog,
		//Camera:       camera,
	}
	// world.InitWorld()
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
	rl.ClearBackground(ui.GameColours["Bg"])

	uiMap := w.Owner.Screen.Get("Map")

	// Draw Terrain
	for x := 0; x < w.Terrain.w; x++ {
		for y := 0; y < w.Terrain.h; y++ {
			pos := position.Position{x, y}
			cell := w.Terrain.Cell(pos)

			if w.FovMap.IsVisible(pos) {
				if cell.T == WallCell {
					uiMap.SetChar(178, pos, ui.GameColours["WallFOV"], ui.ColourNC)
				} else {
					uiMap.SetChar('.', pos, ui.GameColours["FloorFOV"], ui.ColourNC)
				}
			} else if cell.Explored == true {
				if cell.T == WallCell {
					uiMap.SetChar(178, pos, ui.GameColours["Wall"], ui.ColourNC)
				} else {
					uiMap.SetChar('.', pos, ui.GameColours["Floor"], ui.ColourNC)
				}
			}
		}
	}

	// Draw Entities
	for _, entity := range w.Entities.SortedByRenderOrder() {
		if w.FovMap.IsVisible(entity.position) {
			uiMap.SetChar(entity.char, entity.position, ui.GameColours[entity.color], ui.ColourNC)
		}
	}

	w.Camera.FollowTarget(w.Entities.Get("player").position)

	// Tmp Draw Mouse cursor for debug
	var CursorColour rl.Color
	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		CursorColour = ui.GameColours["WallFOV"]
	} else {
		CursorColour = ui.GameColours["Player"]
	}
	w.Owner.Screen.Get("Mouse").SetChar(178, position.Position{X: int(w.MouseX), Y: int(w.MouseY)}, CursorColour, ui.ColourNC)
}

func (w *World) AddMessage(msg SimpleMessage) {
	w.MessageLog.AddMessage(Message{Turn: uint(w.Turn / 10), Message: msg.Message, Colour: msg.Colour})
}

func (w *World) Update(dt float32) {
	playerEntity := w.Entities.Get("player")

	if w.Events.Len() == 0 {
		w.Quit = true
		return
	}

	mousePos := rl.GetMousePosition()
	if mousePos.X > 0 && mousePos.Y > 0 && mousePos.X < float32(rl.GetScreenWidth()) && mousePos.Y < float32(rl.GetScreenHeight()) {
		w.MouseHover = true
		w.MouseX = int(mousePos.X / 10) // divided by cell size... this needs refactoring
		w.MouseY = int(mousePos.Y / 10)
	} else {
		w.MouseHover = false
	}

	newPos := position.Position{0, 0}

	if rl.IsKeyDown(rl.KeyUp) {
		newPos = playerEntity.position.N()
	} else if rl.IsKeyDown(rl.KeyDown) {
		newPos = playerEntity.position.S()
	} else if rl.IsKeyDown(rl.KeyLeft) {
		newPos = playerEntity.position.W()
	} else if rl.IsKeyDown(rl.KeyRight) {
		newPos = playerEntity.position.E()
	} else if rl.IsKeyPressed(rl.KeySpace) {
		w.Owner.ChangeState(NewWorld())
	}

	// Only allow the player to make their turn evey 0.12 seconds... I think it's seconds.
	// This has the effect of ensuring the players turn time is constant and doesn't
	// speed up as the queue gets emptied due to the queue taking less and less time
	// to run.
	if w.inputDelay <= 0 && !newPos.Zero() {
		w.inputDelay = 0.11
		w.NextTurnMove = newPos
	} else {
		w.inputDelay -= dt
	}

	// Run the stack each update, if the next item is PlayerTurn and the
	// playerEntity.NextMove is invalid (e.g Zero) then the event will
	// add itself to the top of the queue to be executed indefinitely
	// until the player gives their input.
	if playerEntity.Exists() {
		ev := w.PopIEvent().Event
		w.Turn = ev.Rank()
		w.Ev = ev
		ev.Action(w)
	}

	// Recompute Player FOV if needed
	if w.FOVRecompute == true {
		w.Terrain.SetExplored(playerEntity.position)
		w.FovMap.ResetVisibility()
		w.FOVAlgo(w.FovMap, playerEntity.position.X, playerEntity.position.Y, 10, true)

		for _, idx := range w.FovMap.visibleCache {
			w.Terrain.SetExplored(position.Idxtopos(idx))
		}

		w.FOVRecompute = false
	}

	// Write to Ui.Statistics
	uiStatistics := w.Owner.Screen.Get("Statistics")

	uiStatistics.SetRow(fmt.Sprintf("HP: %d/%d", w.Entities.Get("player").Fighter.HP, w.Entities.Get("player").Fighter.MaxHP), position.Position{1, 1}, ui.GameColours["Fg"], ui.ColourNC)
	uiStatistics.SetRow(fmt.Sprintf("Turn: %d", w.Turn/10), position.Position{1, 2}, ui.GameColours["Fg"], ui.ColourNC)
	uiStatistics.SetRow(fmt.Sprintf("Mouse (x,y): (%d,%d)", w.MouseX, w.MouseY), position.Position{1, 3}, ui.GameColours["Fg"], ui.ColourNC)
}

func (w World) Save(filename string) error {
	return nil // @todo fill this!!
}

func (w World) GetName() string {
	return "World"
}
