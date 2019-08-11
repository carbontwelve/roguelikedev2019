package state

type GameState interface {
	Pushed(owner *Engine) error // Executed when the state is pushed onto the StateMachine stack
	Popped(owner *Engine) error // Executed when the state is popped off of the StateMachine stack
	Tick(dt float32)
	GetName() string
	Save(filename string) error
	Load(filename string) error
	ShouldQuit() bool
}

type State struct {
	Owner *Engine
	Quit  bool
}

func (s *State) Pushed(owner *Engine) error {
	s.Owner = owner
	return nil
}

func (s *State) Popped(owner *Engine) error {
	return nil
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

func (s State) ShouldQuit() bool {
	return s.Quit
}
