package player

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Gvardmeister/TextGame/internal/domain/room"
)

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
	room := p.CurrentRoom

	items := make([]string, 0, len(room.Items))
	for item := range room.Items {
		items = append(items, item)
	}
	sort.Strings(items)

	description := ""

	switch room.Name {
	case "кухня":
		description += "ты находишься на кухне, "
		if len(items) > 0 {
			description += fmt.Sprintf("на столе: %s, надо собрать рюкзак и идти в универ", strings.Join(items, ", "))
		} else {
			description = "надо идти в универ"
		}
	case "комната":
		if len(items) == 0 {
			description = "пустая комната"
		} else {
			stol := []string{}
			stul := []string{}

			for _, item := range items {
				if item == "рюкзак" {
					stul = append(stul, item)
				} else {
					stol = append(stol, item)
				}
			}

			parts := []string{}
			if len(stol) > 0 {
				parts = append(parts, fmt.Sprintf("на столе: %s", strings.Join(stol, ", ")))
			}
			if len(stul) > 0 {
				parts = append(parts, fmt.Sprintf("на стуле: %s", strings.Join(stul, ", ")))
			}
			description += strings.Join(parts, ", ")
		}
	default:
		if len(items) > 0 {
			description += fmt.Sprintf("на полу: %s", strings.Join(items, ", "))
		} else {
			description = "ничего интересного"
		}
	}

	exits := make([]string, 0, len(room.ConnectionsRoom))
	for name := range room.ConnectionsRoom {
		exits = append(exits, name)
	}
	sort.Strings(exits)
	exitsDescrip := fmt.Sprintf("можно пройти - %s", strings.Join(exits, ", "))

	if description != "" {
		return fmt.Sprintf("%s. %s", description, exitsDescrip)
	}

	return exitsDescrip
}

func (p *Player) MoveTo(roomName string) string {
	room := p.CurrentRoom

	targetRoom, ok := room.ConnectionsRoom[roomName]
	if !ok {
		return fmt.Sprintf("нет пути в %s", roomName)
	}

	if roomName == "улица" && !p.HasItem("ключи") {
		return "дверь закрыта"
	}

	p.CurrentRoom = targetRoom

	switch roomName {
	case "комната":
		return "ты в своей комнате. можно пройти - коридор"
	case "кухня":
		return "кухня, ничего интересного. можно пройти - коридор"
	case "улица":
		return "на улице весна. можно пройти - домой"
	default:
		order := []string{"кухня", "комната", "улица"}
		exits := []string{}
		for _, name := range order {
			if _, ok := targetRoom.ConnectionsRoom[name]; ok {
				exits = append(exits, name)
			}
		}

		return fmt.Sprintf("ничего интересного. можно пройти - %s", strings.Join(exits, ", "))
	}
}

func (p *Player) TakeItem(item string) string {
	room := p.CurrentRoom

	if !room.Items[item] {
		return "нет такого"
	}

	if !p.Equipped["рюкзак"] && item != "рюкзак" {
		return "некуда класть"
	}

	p.Inventory[item] = true
	delete(room.Items, item)

	return fmt.Sprintf("предмет добавлен в инвентарь: %s", item)
}

func (p *Player) UseItem(item, target string) string {
	item = strings.ToLower(item)
	target = strings.ToLower(target)

	if !p.HasItem(item) {
		return fmt.Sprintf("нет предмета в инвентаре - %s", item)
	}

	if item == "ключи" && target == "дверь" {
		if _, ok := p.CurrentRoom.ConnectionsRoom["улица"]; ok {
			return "дверь открыта"
		}

		return "не к чему применить"
	}

	return "не к чему применить"
}

func (p *Player) HasItem(item string) bool {
	return p.Inventory[item]
}

func (p *Player) EquipItem(item string) string {
	item = strings.ToLower(item)

	if p.Equipped[item] {
		return "уже надето"
	}

	if p.HasItem(item) {
		p.Equipped[item] = true

		return fmt.Sprintf("вы надели: %s", item)
	}

	if p.CurrentRoom.Items[item] {
		delete(p.CurrentRoom.Items, item)
		p.Equipped[item] = true

		return fmt.Sprintf("вы надели: %s", item)
	}

	return "нет такого"
}
