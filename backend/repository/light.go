package repository

import (
	"openHueApp/backend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LightRepository struct {
	*BaseRepository
}

func NewLightRepository(db *gorm.DB) *LightRepository {
	return &LightRepository{
		NewBaseRepository(db, models.Light{}),
	}
}

func (r *LightRepository) GetOrCreate(light models.Light) (models.Light, error) {
	id := light.GetId()
	result := r.db.Preload(clause.Associations).Where(&models.Light{Model: models.Model{ID: id}}).FirstOrCreate(&light)
	if result.Error != nil {
		return light, result.Error
	}

	return light, nil
}
