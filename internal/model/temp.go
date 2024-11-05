package model

import "gorm.io/gorm"

type Temp struct {
	gorm.Model
	Name  string `json: "name"`
	Email string `gorm:"unique;not null" json: "email"`
}

func MigrateTemp(db *gorm.DB) {
	db.AutoMigrate(&Temp{})
}
