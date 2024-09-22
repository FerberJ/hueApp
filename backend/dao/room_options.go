package dao

import (
	"openHueApp/backend/models"
	"openHueApp/backend/repository"
)

func (a *App) GetRoomOptions() ([]models.RoomOptions, error) {
	var roomOptions = []models.RoomOptions{}
	var rooms = []models.Room{}
	r := repository.NewRoomRepository(a.SqlLite.Db)
	res, err := r.GetAll()
	if err != nil {
		return roomOptions, nil
	}
	rooms = res.([]models.Room)

	for _, room := range rooms {
		roomOptions = append(roomOptions, models.RoomOptions{
			Value: room.Name,
			Label: room.Name,
		})
	}

	return roomOptions, nil
}
