package handlers

import (
	"encoding/json"
	"github.com/hemantkgupta/my-go-work/gorm-gorilla/models"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Errro in reading body from request: ", err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// Append to the Books table
	if result := h.DB.Create(&book); result.Error != nil {
		log.Println("Error in creatig book: ", result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
