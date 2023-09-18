package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

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

	_, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Error in connecting to database: ", err)
	}
	fmt.Println("Connected Successfully to the Database.")
}
