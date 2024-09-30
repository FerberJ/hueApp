package models

type Group struct {
	Model
	Name   string  `json:"name" gorm:"unique"`
	Lights []Light `json:"lights" gorm:"foreignKey:GroupRef"`
	On     bool    `json:"on" gorm:"-"`
}
