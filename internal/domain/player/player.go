package player

import "github.com/Gvardmeister/TextGame/internal/domain/room"

type Player struct {
	CurrentRoom *room.Room
	Inventory   map[string]bool
	Equipped    map[string]bool
}

func NewPlayer(startRoom *room.Room) *Player {
	return &Player{
		CurrentRoom: startRoom,
		Inventory:   make(map[string]bool),
		Equipped:    make(map[string]bool),
	}
}
