package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sts := `
			DROP TABLE IF EXISTS cars;
			CREATE TABLE cars(id INTEGER PRIMARY KEY, name TEXT, price INT);
			INSERT INTO cars(name, price) VALUES('Audi',52642);
			INSERT INTO cars(name, price) VALUES('Mercedes',57127);
			INSERT INTO cars(name, price) VALUES('Skoda',9000);
			INSERT INTO cars(name, price) VALUES('Volvo',29000);
			INSERT INTO cars(name, price) VALUES('Bentley',350000);
			INSERT INTO cars(name, price) VALUES('Citroen',21000);
			INSERT INTO cars(name, price) VALUES('Hummer',41400);
			INSERT INTO cars(name, price) VALUES('Volkswagen',21600);
			`

	_, err = db.Exec(sts)
	if err != nil {
		log.Println("Error in executing create table query: ", err)
	}
	fmt.Println("The cars table is created.")

	rows, err := db.Query("SELECT * FROM cars")

	if err != nil {
		log.Println("Error in executing select query: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var price int

		err = rows.Scan(&id, &name, &price)

		if err != nil {
			log.Println("Error in scanning the table: ", err)
		}
		fmt.Printf("%d %s %d\n", id, name, price)
	}

	stm, err := db.Prepare("SELECT * FROM cars WHERE id = ?")

	var id int
	var name string
	var price int

	cid := 3

	err = stm.QueryRow(cid).Scan(&id, &name, &price)
	// Another way of writing without Prepare
	// 	row := db.QueryRow("SELECT * FROM cars WHERE id = ?", cid)
	//  err = row.Scan(&id, &name, &price)
	if err != nil {
		log.Println("Error in scan row query: ", err)
	}

	fmt.Printf("%d %s %d\n", id, name, price)

	res, err := db.Exec("DELETE FROM cars WHERE id IN (1, 2, 3)")

	if err != nil {
		log.Println("Error in delete rows query: ", err)
	}

	n, err := res.RowsAffected()

	if err != nil {
		log.Println("Error in getting rows affected: ", err)
	}

	fmt.Printf("The statement has affected %d rows\n", n)
}
