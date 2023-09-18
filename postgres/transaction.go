package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "hemantg"
		password = ""
		dbname   = "hemantgtestdb"
	)

	// Create a new connection to our database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Error in opening database: ", err)
	}

	// Create a new context, and begin a transaction
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Error in beginning transaction: ", err)
	}
	// `tx` is an instance of `*sql.Tx` through which we can execute our queries
	// Here, the query is executed on the transaction instance, and not applied to the database yet
	_, err = tx.ExecContext(ctx, "INSERT INTO pets (name, species) VALUES ('Fido', 'dog'), ('Albert', 'cat')")
	if err != nil {
		// In case we find any error in the query execution, rollback the transaction
		log.Println("Doing Rollback on the transaction.")
		tx.Rollback()
		return
	}

	// The next query is handled similarly
	_, err = tx.ExecContext(ctx, "INSERT INTO food (name, quantity) VALUES ('Dog Biscuit', 3), ('Cat Food', 5)")
	if err != nil {
		log.Println("Doing Rollback on the transaction.")
		tx.Rollback()
		return
	}

	// Finally, if no errors are recieved from the queries, commit the transaction
	// this applies the above changes to our database
	err = tx.Commit()
	if err != nil {
		log.Println("Error in committing the transaction: ", err)
	}
}
