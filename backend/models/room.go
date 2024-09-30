package models

type Room struct {
	Model
	Name   string  `json:"name"`
	Lights []Light `json:"lights" gorm:"foreignKey:RoomRef"`
	On     bool    `json:"on" gorm:"-"`
}
