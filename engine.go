package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"raylibtinkering/ui"
)

type Engine struct {
	states *stack
	screen *ui.Screen
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
	_, s := ui.NewScreen(uint(rl.GetScreenWidth()), uint(rl.GetScreenHeight()), ui.NewTileset("arial10x10.png", ui.LayoutTcod, 10, 10))

	engine := &Engine{
		states: NewStack(),
		screen: s,
	}

	engine.PushState(NewLobbyState(engine))

	return engine
}

func (g *Engine) Unload() {
	g.screen.Unload()
}
