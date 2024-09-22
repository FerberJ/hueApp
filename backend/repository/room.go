package repository

import (
	"openHueApp/backend/models"

	"gorm.io/gorm"
)

type RoomRepository struct {
	*BaseRepository
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{
		NewBaseRepository(db, models.Room{}),
	}
}

func (r *RoomRepository) GetByName(name string) (models.Room, error) {
	var room models.Room
	result := r.db.Where("name = ?", name).First(&room)
	if result.Error != nil {
		return room, result.Error
	}

	return room, nil
}
