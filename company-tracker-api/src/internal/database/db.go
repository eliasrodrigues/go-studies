package database

import (
	"log"

	"company-tracker-api/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("storage/app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	err = db.AutoMigrate(&models.Company{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	return db
}
