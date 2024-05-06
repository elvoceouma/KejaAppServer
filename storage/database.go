package storage

import (
	"log"
	"os"

	model "github.com/elvoceouma/KejaServerSide/domain/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("There is an Error loading the .env file")
	}
	dcn := os.Getenv("DB_CONNECTION_STRING")
	db, dbError := gorm.Open(postgres.Open(dcn), &gorm.Config{})
	if dbError != nil {
		log.Panic("error connection to db")
	}

	DB = db
	return db
}
func performMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
	)
}

func InitializeDatabase() *gorm.DB {
	db := ConnectDatabase()
	performMigrations(db)
	return db
}
