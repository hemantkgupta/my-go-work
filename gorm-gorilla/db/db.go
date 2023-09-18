package db

import (
	"fmt"
	"github.com/hemantkgupta/my-go-work/gorm-gorilla/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "hemantg"
		password = "hemantg@123"
		dbname   = "hemantgtestdb"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		log.Println("Error in connecting to database: ", err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
