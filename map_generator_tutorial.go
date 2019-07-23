package main

import "math/rand"

func TutorialMapGenerator(mapW, mapH, roomMaxSize, roomMinSize, maxRooms, maxMonstersPerRoom int) *GameMap {
	gameMap := NewGameMap(mapW, mapH)

	placeEntities := func(room Rect) {
		monsterCount := rand.Intn(maxMonstersPerRoom)

		for i := 1; i < monsterCount; i++ {

			// Choose a random location in the room
			x := room.x1 + 1 + rand.Intn(room.x2-1)
			y := room.y1 + 1 + rand.Intn(room.y2-1)

			// Check no entity exists at that location

		}

		//	for i in range(number_of_monsters):
		//	# Choose a random location in the room
		//	x = randint(room.x1 + 1, room.x2 - 1)
		//	y = randint(room.y1 + 1, room.y2 - 1)
		//
		//	if not any([entity for entity in entities if entity.x == x and entity.y == y]):
		//if randint(0, 100) < 80:
		//monster = Entity(x, y, 'o', libtcod.desaturated_green)
		//else:
		//monster = Entity(x, y, 'T', libtcod.darker_green)
		//
		//entities.append(monster)

	}

	var isValid bool
	var rooms []Rect
	numRooms := 0

	for i := 1; i < maxRooms; i++ {
		// random width and height
		w := roomMinSize + rand.Intn(roomMaxSize)
		h := roomMinSize + rand.Intn(roomMaxSize)

		// random position without going out of the boundaries of the map
		x := rand.Intn(mapW - w - 1)
		y := rand.Intn(mapH - h - 1)

		newRoom := NewRect(x, y, w, h)
		isValid = true

		// Check the new room is valid
		for _, room := range rooms {
			if newRoom.Intersect(room) {
				isValid = false
				break
			}
		}

		if isValid == true {
			// New room doesnt intersect any other rooms and is therefore
			// valid. Continue.
			gameMap.AddRoom(newRoom)

			// center coordinates of new room, will be useful later
			newX, newY := newRoom.Center()

			if numRooms == 0 {
				// this is the first room, where the player starts at
				gameMap.PlayerStartX = newX
				gameMap.PlayerStartY = newY
			} else {
				// all rooms after the first:
				// connect it to the previous room with a tunnel
				// center coordinates of previous room
				prevX, prevY := rooms[numRooms-1].Center()
				if rand.Intn(1) == 1 {
					// first move horizontally, then vertically
					gameMap.AddHTunnel(prevX, newX, prevY)
					gameMap.AddVTunnel(prevY, newY, newX)
				} else {
					// first move vertically, then horizontally
					gameMap.AddVTunnel(prevY, newY, prevX)
					gameMap.AddHTunnel(prevX, newX, newY)
				}
			}

			// finally, append the new room to the list
			rooms = append(rooms, newRoom)
			numRooms++
		}
	}

	return gameMap
}
