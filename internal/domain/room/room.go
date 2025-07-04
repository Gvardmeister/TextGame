package room

import player_interface "github.com/Gvardmeister/TextGame/internal/domain/interfaces"

type Room struct {
	Name            string
	Items           map[string]bool
	ConnectionsRoom map[string]*Room
	Action          map[string]func(p player_interface.Person, args []string) string // действия (команды) в комнате
}

func NewRoom(name string) *Room {
	return &Room{
		Name:            name,
		Items:           make(map[string]bool),
		ConnectionsRoom: make(map[string]*Room),
		Action:          make(map[string]func(p player_interface.Person, args []string) string),
	}
}
