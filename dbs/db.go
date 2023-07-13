package dbs

import (
	"air-line-reservation-backend/entities"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func GetDB() *gorm.DB {
	return dbInstance
}

func setupPostgresDataSource() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
	dbHost,
	dbPort,
	dbUser,
	dbName,
	dbPassword)

	db, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func IntializeDatabase() error {
	var db *gorm.DB
	var err error
	
	db, err = setupPostgresDataSource()

	if err != nil {
		return err
	}

	err = entities.AutoMigrate(db)
	if err != nil {
		return err
	}
	dbInstance = db
	return nil
}