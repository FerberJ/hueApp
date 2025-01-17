package dao

import (
	"openHueApp/backend/models"
	"openHueApp/backend/repository"

	"github.com/google/uuid"
)

func (a *App) CreateGroup(group models.Group) error {
	r := repository.NewGroupRepository(a.SqlLite.Db)
	err := r.Create(&group)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) CreateGroupWithName(name string) error {
	group := models.Group{
		Name: name,
		Model: models.Model{
			ID: uuid.NewString(),
		},
	}
	r := repository.NewGroupRepository(a.SqlLite.Db)
	err := r.Create(&group)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) GetGroup(id string) (models.Group, error) {
	var group = &models.Group{}
	r := repository.NewGroupRepository(a.SqlLite.Db)
	res, err := r.Get(id)
	if err != nil {
		return *group, err
	}

	group = res.(*models.Group)

	return *group, nil
}

func (a *App) GetGroups() ([]models.Group, error) {
	var groups = []models.Group{}
	r := repository.NewGroupRepository(a.SqlLite.Db)
	res, err := r.GetAll()
	if err != nil {
		return groups, err
	}

	groups = res.([]models.Group)

	for _, group := range groups {
		for _, light := range group.Lights {
			group.On = false
			if light.On {
				group.On = true
			}
		}
	}

	return groups, nil
}

func (a *App) UpdateGroup(group models.Group) error {
	r := repository.NewGroupRepository(a.SqlLite.Db)
	err := r.Update(group)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) DeleteGroup(id string) error {
	r := repository.NewGroupRepository(a.SqlLite.Db)
	err := r.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) GetGroupByName(name string) (models.Group, error) {
	r := repository.NewGroupRepository(a.SqlLite.Db)
	return r.GetByName(name)
}

func (a *App) ToggleGroupLike(group models.Group) error {
	changedEntity := group.ToggleLiked()

	r := repository.NewGroupRepository(a.SqlLite.Db)
	return r.ToggleLiked(changedEntity, "groups")
}

func (a *App) GetLikedGroups() ([]models.Group, error) {
	var groups = []models.Group{}
	r := repository.NewGroupRepository(a.SqlLite.Db)
	res, err := r.GetAllLiked()
	if err != nil {
		return groups, err
	}

	groups = res.([]models.Group)

	for _, group := range groups {
		for _, light := range group.Lights {
			group.On = false
			if light.On {
				group.On = true
			}
		}
	}

	return groups, nil
}
