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

func (p *Player) LookAround() string {
	return ""
}

func (p *Player) MoveTo(roomName string) string {
	return ""
}

func (p *Player) TakeItem(item string) string {
	return ""
}

func (p *Player) UseItem(item, target string) string {
	return ""
}
