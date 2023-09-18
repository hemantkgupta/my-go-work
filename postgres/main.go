package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "hemantg"
		password = ""
		dbname   = "hemantgtestdb"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Error in opening database: ", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("Error in pinging database: ", err)
	}
	fmt.Println("Successfully connected!")

	insertStmt := `insert into "students"("name", "roll_number") values('Jacob', 20)`
	_, err = db.Exec(insertStmt)
	if err != nil {
		log.Println("Error in inserting data into database: ", err)
	}

	insertDynStmt := `insert into "students"("name", "roll_number") values($1, $2)`
	_, err = db.Exec(insertDynStmt, "Jack", 21)
	if err != nil {
		log.Println("Error in inserting data into database: ", err)
	}
}
