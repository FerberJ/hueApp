package repository

import (
	"openHueApp/backend/models"

	"gorm.io/gorm"
)

type GroupRepository struct {
	*BaseRepository
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{
		NewBaseRepository(db, models.Group{}),
	}
}

func (r *GroupRepository) GetByName(name string) (models.Group, error) {
	var group models.Group
	result := r.db.Where("name = ?", name).First(&group)
	if result.Error != nil {
		return group, result.Error
	}

	return group, nil
}
