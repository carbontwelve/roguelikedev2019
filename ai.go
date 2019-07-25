package main

import "fmt"

type Ai interface {
	Tick(self *Entity, entities *Entities, terrain Terrain, fov FovMap)
}

type NullAi struct{}

func (n NullAi) Tick(self *Entity, entities *Entities, terrain Terrain, fov FovMap) {
	// NOP
}

type SimpleAi struct{}

func (a SimpleAi) Tick(self *Entity, entities *Entities, terrain Terrain, fov FovMap) {
	// fmt.Println ("The " + self.Name + " ponders the meaning of its existence.")

	player := entities.Get("player")

	if fov.IsVisible(self.position) {
		if self.position.Distance(player.position) >= 2 {
			self.MoveTowards(player.position, *entities, terrain)
		} else if player.Fighter.HP > 0 {
			fmt.Println("The " + self.Name + " insults you")
		}
	}

}
