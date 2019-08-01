package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Brain interface {
	HandleTurn(w *World, ev event)
	SetOwner(e *Entity)
}

type HindBrain struct {
	owner *Entity
}

func (b *HindBrain) SetOwner(e *Entity) {
	b.owner = e
}

func (b HindBrain) HandleInteractionResults(w *World, results InteractionResults) {
	for _, result := range results.Results {
		if val, ok := result["message"]; ok {
			w.AddMessage(val.(SimpleMessage))

		}

		if val, ok := result["death"]; ok {
			e := val.(*Entity)
			e.Name = "remains of " + e.Name
			e.color = rl.Red
			e.char = '%'
			e.RenderOrder = RoCorpse
		}
	}
}

type PlayerBrain struct {
	HindBrain
}

func (b PlayerBrain) HandleTurn(w *World, ev event) {
	at := w.NextTurnMove

	if at.Zero() == false {
		if w.Terrain.Cell(at).T == FreeCell {
			target := w.Entities.GetBlockingAtPosition(at)
			if target != nil {
				// Player is moving and destination is blocked by Entity
				// @todo refactor for items?
				// @todo some entities may block but not be attackable?
				b.HandleInteractionResults(w, b.owner.Fighter.Attack(target))
			} else {
				// Player is moving and destination is unblocked by terrain lets move
				// to the destination position
				b.owner.MoveTo(at)
			}
			w.FOVRecompute = true
		}
		w.NextTurnMove = Position{0, 0}
	}

	// @see https://github.com/anaseto/boohu/blob/e193aa0453dce8b7ffcae62cfcd79877cb01635d/player.go#L447
	ev.Renew(w, 10)
}

type MonsterBrain struct {
	HindBrain
}

func (b MonsterBrain) HandleTurn(w *World, ev event) {
	target := w.Entities.Get("player")
	if w.FovMap.IsVisible(b.owner.position) {
		if b.owner.position.Distance(target.position) >= 2 {
			b.owner.MoveTowards(target.position, *w.Entities, *w.Terrain)
		} else if target.Fighter.HP > 0 {
			b.HandleInteractionResults(w, b.owner.Fighter.Attack(target))
		}
	}

	if b.owner.Exists() {
		ev.Renew(w, 10)
	}
}
