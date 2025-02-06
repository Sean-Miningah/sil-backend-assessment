package repos

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPGStorage(datasource string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(datasource), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
