package db

import (
	"log"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	dbModels "demoapp/db/models"
)

var DB *gorm.DB

func ConnectToDb(wg *sync.WaitGroup) {
	defer wg.Done()

	DATABASE_URL := viper.GetString("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(DATABASE_URL))

	if err != nil {
		log.Fatalln("Failed to connect to database")
	}

	db.AutoMigrate(
		&dbModels.User{},
		&dbModels.Task{},
	)

	DB = db
}

func GetDb() *gorm.DB {
	return DB
}
