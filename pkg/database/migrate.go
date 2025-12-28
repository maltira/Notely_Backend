package database

import (
	"web-lab/internal/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Group{},
		&entity.Publication{},
		&entity.Category{},
		&entity.PublicationCategories{},
	)
	if err != nil {
		panic(err)
	}
}
