package main

import (
	"net/http"

	"github.com/ferparrios/Go_Spider_Rest_Api/db"
	"github.com/ferparrios/Go_Spider_Rest_Api/models"
	"github.com/ferparrios/Go_Spider_Rest_Api/routes"
	"github.com/gorilla/mux"
)

func main(){
	db.DBConnection()

	db.DB.AutoMigrate(models.Quote{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/api/quotes", routes.GetQuotesHandler).Methods("GET")
	r.HandleFunc("/api/quotes", routes.CreateQuotesHandler).Methods("POST")


	http.ListenAndServe(":3000", r)
}