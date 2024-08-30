package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=127.0.0.1 user=ferparrios password=mysecretpassword dbname=quotes port=5431 sslmode=disable"
var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil{
		log.Fatal(error)
	} else{
		log.Println("Data Base Connected")
	}
}