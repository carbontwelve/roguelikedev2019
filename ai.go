package main

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

type PlayerBrain struct {
	HindBrain
}

func (b PlayerBrain) HandleTurn(w *World, ev event) {
	at := w.NextTurnMove

	if at.Zero() == false && w.Terrain.Cell(at).T == FreeCell {
		target := w.Entities.GetBlockingAtPosition(at)
		if target != nil {
			// Player is moving and destination is blocked by Entity
			// @todo refactor for items?
			// @todo some entities may block but not be attackable?
			b.owner.Fighter.Attack(target)
		} else {
			// Player is moving and destination is unblocked by terrain lets move
			// to the destination position
			b.owner.MoveTo(at)
		}
		w.FOVRecompute = true
		w.NextTurnMove = Position{0, 0}
	}

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
			b.owner.Fighter.Attack(target)
		}
	}

	ev.Renew(w, 10)
}
