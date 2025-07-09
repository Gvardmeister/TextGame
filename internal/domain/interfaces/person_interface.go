package interfaces

type Person interface {
	// Для проверки состояния и управления
	HasItem(item string) bool
	IsEquipped(item string) bool
	GetCurrentRoomName() string
	SetCurrentRoom(name string)

	// Методы для команд
	AddItem(item string) string
	EquipItem(item string) string

	LookAround() string
	MoveTo(roomName string) string
	TakeItem(item string) string
	UseItem(item, target string) string
}
