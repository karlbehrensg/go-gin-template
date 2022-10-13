package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DSN = "host=localhost user=postgres password=postgres dbname=templates port=5432"
var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "go_gin.",
			SingularTable: false,
		}})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connected")
	}
}
