package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Book struct {
	id     int
	name   string
	author string
}

func main() {
	db, err := sql.Open("sqlite3", "books.db")
	if err != nil {
		log.Println("Error in opening the books database: ", err)
	}

	query1 := `CREATE TABLE IF NOT EXISTS books
	         (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)`
	statement, _ := db.Prepare(query1)
	_, err = statement.Exec()
	if err != nil {
		log.Println("Error in executing query: ", err)
	}

	query2 := `INSERT INTO books (name, author, isbn) VALUES (?, ?, ?)`
	statement, _ = db.Prepare(query2)
	_, err = statement.Exec("A Tale of Two Cities", "Charles Dickens", 140430547)
	if err != nil {
		log.Println("Error in executing query: ", err)
	}
	log.Println("Inserted the book into database!")

	query3 := `SELECT id, name, author FROM books`
	rows, _ := db.Query(query3)
	var tempBook Book
	for rows.Next() {
		err = rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		if err != nil {
			log.Println("Error in reading rows from scan: ", err)
		}
		log.Printf("ID:%d, Book:%s, Author:%s\n", tempBook.id,
			tempBook.name, tempBook.author)
	}
}
