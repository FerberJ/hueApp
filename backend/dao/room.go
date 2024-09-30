package dao

import (
	"openHueApp/backend/models"
	"openHueApp/backend/repository"

	"github.com/google/uuid"
)

func (a *App) CreateRoom(room models.Room) error {
	r := repository.NewRoomRepository(a.SqlLite.Db)
	err := r.Create(&room)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) GetRoom(id string) (models.Room, error) {
	var room = &models.Room{}
	r := repository.NewRoomRepository(a.SqlLite.Db)
	res, err := r.Get(id)
	if err != nil {
		return *room, err
	}

	room = res.(*models.Room)

	return *room, nil
}

func (a *App) GetRooms() ([]models.Room, error) {
	var rooms = []models.Room{}
	r := repository.NewRoomRepository(a.SqlLite.Db)
	res, err := r.GetAll()
	if err != nil {
		return rooms, err
	}

	rooms = res.([]models.Room)

	return rooms, nil
}

func (a *App) CreateRoomWithName(name string) error {
	group := models.Room{
		Name: name,
		Model: models.Model{
			ID: uuid.NewString(),
		},
	}
	r := repository.NewRoomRepository(a.SqlLite.Db)
	err := r.Create(&group)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) UpdateRoom(room models.Room) error {
	r := repository.NewRoomRepository(a.SqlLite.Db)
	err := r.Update(room)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) DeleteRoom(id string) error {
	r := repository.NewRoomRepository(a.SqlLite.Db)
	err := r.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) GetRoomByName(name string) (models.Room, error) {
	r := repository.NewRoomRepository(a.SqlLite.Db)
	return r.GetByName(name)
}

func (a *App) ToggleRoomLike(room models.Room) error {
	changedEntity := room.ToggleLiked()

	r := repository.NewLightRepository(a.SqlLite.Db)
	return r.ToggleLiked(changedEntity, "rooms")
}
