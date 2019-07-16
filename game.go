package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	states    *stack
	TxSprites rl.Texture2D
}

func (g *Game) PushState(state GameState) {
	g.states.Push(state)
}

func (g *Game) PopState() {
	g.states.Pop()
}

func (g *Game) ChangeState(state GameState) {
	if g.states.Len() > 0 {
		g.states.Pop()
	}
	g.PushState(state)
}

func (g *Game) PeekState() GameState {
	if g.states.Len() == 0 {
		return nil
	}
	return g.states.Peek().(GameState)
}

func newGame() *Game {
	game := &Game{
		states:    NewStack(),
		TxSprites: rl.LoadTexture("sprites.png"),
	}

	game.PushState(NewLobbyState(game))

	return game
}

func (g *Game) Unload() {
	rl.UnloadTexture(g.TxSprites)
}
