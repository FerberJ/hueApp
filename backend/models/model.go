package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        string         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updateAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Liked     bool           `json:"liked"`
}

type ModelInterface interface {
	GetId() string
	ToggleLiked() Model
}

// Get ID
func (model Model) GetId() string {
	return model.ID
}

func (model Model) ToggleLiked() Model {
	model.Liked = !model.Liked
	return model
}
