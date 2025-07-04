package game

import (
	"github.com/Gvardmeister/TextGame/internal/domain/player"
	"github.com/Gvardmeister/TextGame/internal/domain/room"
)

var theGame *Game

func initGame() {
	theGame = NewGame()
	theGame.InitGame()
}

type Game struct {
	Player *player.Player
	Rooms  map[string]*room.Room
}

func NewGame() *Game {
	return &Game{
		Rooms: make(map[string]*room.Room),
	}
}

func (g *Game) InitGame() {
	// Создание кухни
	kitchen := room.NewRoom("кухня")
	kitchen.Items["чай"] = true

	// Создание комнаты
	roomOne := room.NewRoom("комната")
	roomOne.Items["ключи"] = true
	roomOne.Items["рюкзак"] = true
	roomOne.Items["конспекты"] = true

	// Создание коридора и улицы
	corridor := room.NewRoom("коридор")
	street := room.NewRoom("улица")

	// Связи между комнатами
	kitchen.ConnectionsRoom["коридор"] = corridor
	roomOne.ConnectionsRoom["коридор"] = corridor
	corridor.ConnectionsRoom["кухня"] = kitchen
	corridor.ConnectionsRoom["комната"] = roomOne
	corridor.ConnectionsRoom["улица"] = street
	street.ConnectionsRoom["коридор"] = corridor

	// Добавление комнат
	g.Rooms["кухня"] = kitchen
	g.Rooms["комната"] = roomOne
	g.Rooms["коридор"] = corridor
	g.Rooms["улица"] = street

	// Создание игрока
	g.Player = player.NewPlayer(kitchen)
}

func handleCommand(command string) string {
	return ""
}
