package main

type GameState interface {
	Draw(dt float32)
	Update(dt float32)
	GetName() string
	Save(filename string) error
	Load(filename string) error
}

type State struct {
	e *Engine
}

//
// Save State to disk
//
func (s State) Save(filename string) error {
	return nil
}

//
// Load State from disk
//
func (s *State) Load(filename string) error {
	return nil
}
