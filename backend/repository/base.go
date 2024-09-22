package repository

import (
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseRepository struct {
	db   *gorm.DB
	Type any
}

func NewBaseRepository(db *gorm.DB, entity any) *BaseRepository {
	return &BaseRepository{
		db:   db,
		Type: entity,
	}
}

type Repository interface {
	Create(entity any) error
	Get(id string) (any, error)
	Update(entity any) error
	Delete(id string) error
	GetAll() (any, error)
}

// getType returns the reflect.Type of the BaseRepository.
// It is used to retrieve the type information of the repository.
func (r *BaseRepository) getType() reflect.Type {
	return reflect.TypeOf(r.Type)
}

// getSliceInterface returns a new slice of the entity type.
// It uses reflection to create a new slice based on the type of the entity.
func (r *BaseRepository) getSliceInterface() any {
	// Creates a new zero-valued slice of the type stored in r, and returns it as an interface{}.
	return reflect.New(reflect.SliceOf(r.getType())).Elem().Interface()
}

// getInterface returns a new instance of the interface type associated with the BaseRepository.
// It uses reflection to create a new instance of the type specified by GetType method.
func (r *BaseRepository) getInterface() any {
	// Creates a new zero-valued instance of the type stored in r, and returns it as an interface{}.
	return reflect.New(r.getType()).Interface()
}

/*
Creates new Entity on the DB.

WARNING: Expect Pointer or else it will panic

	err := repository.Create(&entity)
*/
func (r *BaseRepository) Create(entity any) error {
	result := r.db.Create(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/*
Get entity by providing the id

	res, err := repository.Get(3)
	light := res.(*models.Light)
*/
func (r *BaseRepository) Get(id string) (any, error) {
	newEntity := r.getInterface()

	result := r.db.Preload(clause.Associations).First(newEntity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return newEntity, nil
}

/*
Updates Entity on the DB.

	err := repository.Update(&light)
*/
func (r *BaseRepository) Update(entity any) error {
	return r.db.Save(entity).Error
}

/*
Delete entity by providing the id

	err := respository.Delete(1)
*/
func (r *BaseRepository) Delete(id string) error {
	newEntity := r.getInterface()
	result := r.db.Delete(newEntity, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

/*
Gets list of entity from DB

Usage (example with models.Course):

	res, err := repository.GetAll()
	lights := res.([]models.Light)
*/
func (r *BaseRepository) GetAll() (any, error) {
	entities := r.getSliceInterface()
	result := r.db.Preload(clause.Associations).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}

func (r *BaseRepository) ToggleLiked(entity any, table string) error {
	return r.db.Table(table).Save(entity).Error
}
