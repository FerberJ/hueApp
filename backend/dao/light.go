package dao

import (
	"openHueApp/backend/models"
	"openHueApp/backend/repository"
)

func (a *App) CreateLight(light models.Light) error {
	r := repository.NewLightRepository(a.SqlLite.Db)
	err := r.Create(&light)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) GetLight(id string) (models.Light, error) {
	var light = &models.Light{}
	r := repository.NewLightRepository(a.SqlLite.Db)
	res, err := r.Get(id)
	if err != nil {
		return *light, err
	}

	light = res.(*models.Light)

	return *light, nil
}

func (a *App) GetLights() ([]models.Light, error) {
	var lights = []models.Light{}
	r := repository.NewLightRepository(a.SqlLite.Db)
	res, err := r.GetAll()
	if err != nil {
		return lights, err
	}

	lights = res.([]models.Light)

	return lights, nil
}

func (a *App) UpdateLight(light models.Light, roomName, groupName string) error {
	r := repository.NewLightRepository(a.SqlLite.Db)

	if groupName != "" {
		group, err := a.GetGroupByName(groupName)
		if err != nil {
			return err
		}
		light.Group = group
	}

	if roomName != "" {
		room, err := a.GetRoomByName(roomName)
		if err != nil {
			return err
		}
		light.Room = room
	}

	return r.Update(&light)
}

func (a *App) DeleteLight(id string) error {
	r := repository.NewLightRepository(a.SqlLite.Db)
	return r.Delete(id)
}

func (a *App) GetOrCreateLight(light models.Light) (models.Light, error) {
	r := repository.NewLightRepository(a.SqlLite.Db)
	return r.GetOrCreate(light)
}

func (a *App) ToggleLightLike(light models.Light) error {
	changedEntity := light.ToggleLiked()

	r := repository.NewLightRepository(a.SqlLite.Db)
	return r.ToggleLiked(changedEntity, "lights")
}
