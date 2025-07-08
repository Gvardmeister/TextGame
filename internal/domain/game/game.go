package game

import (
	"strings"

	"github.com/Gvardmeister/TextGame/internal/domain/player"
	"github.com/Gvardmeister/TextGame/internal/domain/room"
	"github.com/Gvardmeister/TextGame/internal/domain/state"
)

var theGame *Game

func initGame() {
	theGame = NewGame()
	theGame.InitGame()
}

type Game struct {
	Player     *player.Player
	Rooms      map[string]*room.Room
	StreetRoom *room.Room
}

func NewGame() *Game {
	return &Game{
		Rooms: make(map[string]*room.Room),
	}
}

func (g *Game) InitGame() {
	state.DoorOpened = false

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
	g.Player = player.NewPlayer(kitchen, street)
	g.StreetRoom = street
}

func handleCommand(command string) string {
	if theGame == nil {
		initGame()
	}

	parts := strings.Split(command, " ")
	if len(parts) == 0 {
		return "пустая комната"
	}

	switch strings.ToLower(parts[0]) {
	case "осмотреться":
		return theGame.Player.LookAround()
	case "идти":
		if len(parts) < 2 {
			return "параметр для перемещения - отсутствует."
		}
		return theGame.Player.MoveTo(strings.ToLower(parts[1]))
	case "взять":
		if len(parts) < 2 {
			return "параметр для взятия объекта - отсутствует."
		}
		return theGame.Player.TakeItem(strings.ToLower(parts[1]))
	case "применить":
		if len(parts) < 3 {
			return "параметры для применения - отсутствуют."
		}
		return theGame.Player.UseItem(strings.ToLower(parts[1]), strings.ToLower(parts[2]))
	case "надеть":
		if len(parts) < 2 {
			return "параметр для надевания - отсутствует."
		}
		return theGame.Player.EquipItem(strings.ToLower(parts[1]))
	default:
		return "неизвестная команда"
	}
}
