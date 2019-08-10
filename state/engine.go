package state

import (
	"github.com/gen2brain/raylib-go/raylib"
	"raylibtinkering/ui"
)

type Engine struct {
	States *Stack
	Screen *ui.Screen
}

func (e *Engine) PushState(state GameState) {
	state.Pushed(e)
	e.States.Push(state)
}

func (e *Engine) PopState() {
	e.PeekState().Popped(e)
	e.States.Pop()
}

func (e *Engine) ChangeState(state GameState) {
	if e.States.Len() > 0 {
		e.States.Pop()
	}
	e.PushState(state)
}

func (e *Engine) PeekState() GameState {
	if e.States.Len() == 0 {
		return nil
	}
	return e.States.Peek().(GameState)
}

func NewEngine(initialState GameState) *Engine {
	_, s := ui.NewScreen(uint(rl.GetScreenWidth()), uint(rl.GetScreenHeight()), ui.NewTileset("arial10x10.png", ui.LayoutTcod, 10, 10))

	engine := &Engine{
		States: NewStack(),
		Screen: s,
	}

	engine.PushState(initialState)

	return engine
}

func (e *Engine) Unload() {
	e.Screen.Unload()
}
