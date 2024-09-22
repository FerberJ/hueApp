package connection

import (
	"openHueApp/backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteDb struct {
	Path string
	Db   *gorm.DB
}

func NewSqliteDb(path string) *SqliteDb {
	db, _ := gorm.Open(sqlite.Open(path))
	s := SqliteDb{
		Path: path,
		Db:   db,
	}

	s.Db.AutoMigrate(
		&models.Light{},
		&models.Group{},
		&models.Room{},
	)

	return &s
}

func (s *SqliteDb) Close() error {
	db, err := s.Db.DB()
	if err != nil {
		return err
	}

	db.Close()
	return nil
}
