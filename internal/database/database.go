package database

import (
	"example/restapi/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	models.MigrateAuthor(db)
	models.MigratePost(db)
	models.MigrateComment(db)

	return db, nil
}
