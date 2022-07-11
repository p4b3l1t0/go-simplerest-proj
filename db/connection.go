package db 

import (
	"gorm.io/gorm"
	"gorm.io/driver/posgres"
)

var DSN = "host=localhost user=p4b3l password=mysecretpassword dbname=gorm port=5432"
var DB *gorm.DB

func DBconnect(){
	var error error
	DB, error := gorm.Open(postgres.Open(DNS), &gorm.config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Database PosgreSQL connected")
	}

}