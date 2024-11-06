package database

import (
	"example/restapi/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db

	model.MigrateAuthor(db)
	model.MigratePost(db)
	model.MigrateComment(db)

	return db, nil
}
