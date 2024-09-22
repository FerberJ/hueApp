package dao

import (
	"openHueApp/backend/models"
	"openHueApp/backend/repository"
)

func (a *App) GetGroupOptions() ([]models.GroupOptions, error) {
	var groupOptions = []models.GroupOptions{}
	var groups = []models.Group{}
	r := repository.NewGroupRepository(a.SqlLite.Db)
	res, err := r.GetAll()
	if err != nil {
		return groupOptions, nil
	}
	groups = res.([]models.Group)

	for _, group := range groups {
		groupOptions = append(groupOptions, models.GroupOptions{
			Value: group.Name,
			Label: group.Name,
		})
	}

	return groupOptions, nil
}
