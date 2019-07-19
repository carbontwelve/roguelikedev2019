package main

type GameState interface {
	Draw(dt float32)
	Update(dt float32)
	GetName() string
	Save(filename string) error
	Load(filename string) error
}

type State struct {
	e        *Engine
	Entities map[string]*Entity
	GameMap  *GameMap
}

func (s *State) SetEntities(e map[string]*Entity) {
	s.Entities = e
}

func (s *State) SetGameMap(m *GameMap) {
	s.GameMap = m
}

func (s *State) SetEntity(name string, e *Entity) {
	if s.Entities == nil {
		s.SetEntities(make(map[string]*Entity))
	}
	s.Entities[name] = e
}

func (s *State) GetEntity(name string) *Entity {
	return s.Entities[name]
}

func (s State) DrawEntities() {
	for _, entity := range s.Entities {
		entity.Draw(s.e)
	}
}

func (s State) DrawMap() {
	s.GameMap.Draw(s.e)
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
