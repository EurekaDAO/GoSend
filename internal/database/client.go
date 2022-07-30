package database

import (
	"log"

	"github.com/eurekadao/gosend/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	Instance.AutoMigrate(
		&models.Email{},
	)
	log.Println("Database Migration Completed!")
}

func SdkDbStart(_instance *gorm.DB) {
	Instance = _instance
}
