package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Engine struct {
	states  *stack
	sprites *SpriteSheet
	font    *Font
	ui      int
}

func (g *Engine) PushState(state GameState) {
	g.states.Push(state)
}

func (g *Engine) PopState() {
	g.states.Pop()
}

func (g *Engine) ChangeState(state GameState) {
	if g.states.Len() > 0 {
		g.states.Pop()
	}
	g.PushState(state)
}

func (g *Engine) PeekState() GameState {
	if g.states.Len() == 0 {
		return nil
	}
	return g.states.Peek().(GameState)
}

func newEngine() *Engine {
	engine := &Engine{
		states:  NewStack(),
		sprites: newSpriteSheet(rl.LoadTexture("arial10x10.png"), 10, 10),
		font:    newFont("arial10x10.png", 10, 10),
	}

	engine.PushState(NewLobbyState(engine))

	return engine
}

func (g *Engine) Unload() {
	g.sprites.Unload()
}
