package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DNS = "host=localhost user=postgres password=postgres dbname=postgres port=5432"
var DB *gorm.DB

func DBConecction() {

	var error error
	DB, error = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("DB Connection...")
	}
}
