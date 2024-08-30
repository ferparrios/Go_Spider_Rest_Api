package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ferparrios/Go_Spider_Rest_Api/db"
	"github.com/ferparrios/Go_Spider_Rest_Api/models"
)

func GetQuotesHandler(w http.ResponseWriter, r *http.Request){
	var quotes []models.Quote
	db.DB.Find(&quotes)
	json.NewEncoder(w).Encode(&quotes)
}

func CreateQuotesHandler(w http.ResponseWriter, r *http.Request){
	var quote models.Quote
	json.NewDecoder((r.Body)).Decode(&quote)
	createdQuote := db.DB.Create(&quote)
	err := createdQuote.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&quote)
}