package main

import (
	"github.com/gorilla/mux"
	"github.com/hemantkgupta/my-go-work/gorm-gorilla/db"
	"github.com/hemantkgupta/my-go-work/gorm-gorilla/handlers"
	"log"
	"net/http"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	//router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	//router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	//router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	//router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":8080", router)
}
