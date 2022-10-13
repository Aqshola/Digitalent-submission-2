package database

import (
	"fmt"
	"log"
	"restapi/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
	db       *gorm.DB
	err      error
)

func StartDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open((postgres.Open(psqlInfo)), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connect database", err)
	}

	db.Debug().AutoMigrate(models.Orders{}, models.Item{})
}

func GetDB() *gorm.DB {
	return db
}
