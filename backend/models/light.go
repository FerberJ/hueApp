package models

type Light struct {
	Model
	CustomName string  `json:"customName"`
	Name       string  `json:"name"`
	On         bool    `json:"on" gorm:"-"`
	GroupRef   string  `json:"groupRef"`
	Group      Group   `json:"group" gorm:"foreignKey:GroupRef"`
	RoomRef    string  `json:"roomRef"`
	Room       Room    `json:"room" gorm:"foreignKey:RoomRef"`
	Brightness float32 `json:"brightness" gorm:"-"`
}
