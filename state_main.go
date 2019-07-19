package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"time"
)

//type event interface {
//	Rank() int
//	Action(*game)
//	Renew(*game, int)
//}

type MainState struct {
	State
}

func NewMainState(e *Engine) *MainState {
	rand.Seed(time.Now().Unix())

	s := &MainState{
		State: State{e: e},
	}

	gameMap := TutorialMapGenerator(80, 45, 10, 6, 30)
	s.SetEntity("player", NewEntity(gameMap.PlayerStartX, gameMap.PlayerStartY, '@', PlayerColour))
	s.SetGameMap(gameMap)
	return s
}

func (s MainState) Draw(dt float32) {
	rl.ClearBackground(UIBackgroundColour)
	s.DrawMap()
	s.DrawEntities()
}

func (s *MainState) Update(dt float32) {
	playerEntity := s.GetEntity("player")
	gm := *s.GameMap

	if rl.IsKeyPressed(rl.KeyUp) {
		playerEntity.Move(0, -1, gm)
	} else if rl.IsKeyPressed(rl.KeyDown) {
		playerEntity.Move(0, 1, gm)
	} else if rl.IsKeyPressed(rl.KeyLeft) {
		playerEntity.Move(-1, 0, gm)
	} else if rl.IsKeyPressed(rl.KeyRight) {
		playerEntity.Move(1, 0, gm)
	} else if rl.IsKeyPressed(rl.KeySpace) {
		s.e.ChangeState(NewMainState(s.e))
	}
}

func (s MainState) GetName() string {
	return "Main"
}
