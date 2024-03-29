package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

/**
 * Some of the event code in this file originated from the game boohu
 * and is licensed ISC.
 *
 * Copyright (c) 2017 Yon <anaseto@bardinflor.perso.aquilenet.fr>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 *
 * @see https://github.com/anaseto/boohu/blob/master/events.go
 */

type event interface {
	Rank() int
	Action(*World)
	Renew(*World, int)
}

type iEvent struct {
	Event event
	Index int
}

type eventQueue []iEvent

func (evq eventQueue) Len() int {
	return len(evq)
}

func (evq eventQueue) Less(i, j int) bool {
	return evq[i].Event.Rank() < evq[j].Event.Rank() ||
		evq[i].Event.Rank() == evq[j].Event.Rank() && evq[i].Index < evq[j].Index
}

func (evq eventQueue) Swap(i, j int) {
	evq[i], evq[j] = evq[j], evq[i]
}

func (evq *eventQueue) Push(x interface{}) {
	no := x.(iEvent)
	*evq = append(*evq, no)
}

func (evq *eventQueue) Pop() interface{} {
	old := *evq
	n := len(old)
	no := old[n-1]
	*evq = old[0 : n-1]
	return no
}

type simpleAction int

const (
	PlayerTurn simpleAction = iota
	HealPlayer
)

type simpleEvent struct {
	ERank   int
	EAction simpleAction
}

func (sev *simpleEvent) Rank() int {
	return sev.ERank
}

func (sev *simpleEvent) Renew(w *World, delay int) {
	sev.ERank += delay
	if delay == 0 {
		w.PushAgainEvent(sev)
	} else {
		w.PushEvent(sev)
	}
}

func (sev *simpleEvent) Action(w *World) {
	switch sev.EAction {
	case PlayerTurn:
		rl.DrawText("Turn: Player", int32(rl.GetScreenWidth()-100), 30, 10, PlayerColour)

		if w.NextTurnMove.Zero() {
			// If no player input wait
			sev.Renew(w, 0)
			return
		}

		playerEntity := w.Entities.Get("player")

		if playerEntity.Exists() {
			playerEntity.Brain.HandleTurn(w, sev)
		}

	case HealPlayer:
		playerEntity := w.Entities.Get("player")
		playerEntity.Fighter.Heal(1)
		sev.Renew(w, 500)

		w.AddMessage(SimpleMessage{Message: "You feel your health improving", Colour: rl.Green})
	}
}

type monsterAction int

const (
	MonsterTurn monsterAction = iota
	MonsterDeath
)

type monsterEvent struct {
	ERank   int
	NMons   string
	EAction monsterAction
}

func (mev *monsterEvent) Rank() int {
	return mev.ERank
}

func (mev *monsterEvent) Action(w *World) {
	switch mev.EAction {
	case MonsterTurn:
		e := w.Entities.Get(mev.NMons)
		if e.Exists() {
			rl.DrawText(fmt.Sprintf("Turn: %s", mev.NMons), int32(rl.GetScreenWidth()-100), 30, 10, PlayerColour)
			e.Brain.HandleTurn(w, mev)
		} else {
			// We have died
			w.PushEvent(&monsterEvent{ERank: 0, EAction: MonsterDeath, NMons: mev.NMons})
		}
	case MonsterDeath:
		e := w.Entities.Get(mev.NMons)
		e.Name = "Dead " + e.Name
		e.color = rl.Red
		e.char = '%'
		e.RenderOrder = RoCorpse
	}
}
func (mev *monsterEvent) Renew(w *World, delay int) {
	mev.ERank += delay
	w.PushEvent(mev)
}
