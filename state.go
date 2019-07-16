package main

type GameState interface {
	Draw(dt float32)
	Update(dt float32)
	SetGame(g *Game)
	GetName() string
}

type State struct {
	g *Game
}

func (s *State) SetGame(g *Game) {
	s.g = g
}
