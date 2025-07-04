package player_interface

type Person interface {
	HasItem(item string) bool
	AddItem(item string)
	EquipItem(item string)
	IsEquipped(item string) bool
	GetCurrentRoomName() string
	SetCurrentRoom(name string)
}
