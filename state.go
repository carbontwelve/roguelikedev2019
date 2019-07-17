package main

type GameState interface {
	Draw(dt float32)
	Update(dt float32)
	SetGame(g *Engine)
	GetName() string
}

type State struct {
	g *Engine
}

func (s *State) SetGame(g *Engine) {
	s.g = g
}
