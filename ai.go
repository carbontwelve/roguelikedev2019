package main

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

	target := entities.Get("player")

	if fov.IsVisible(self.position) {
		if self.position.Distance(target.position) >= 2 {
			self.MoveTowards(target.position, *entities, terrain)
		} else if target.Fighter.HP > 0 {
			self.Fighter.Attack(target)
		}
	}
}
