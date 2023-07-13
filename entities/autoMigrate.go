package entities

import (
	"air-line-reservation-backend/entities/example"

	"gorm.io/gorm"
)

/**
	Auto Migration for the desired entities
	For example in this case is: Example
**/
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&example.Example{})
}