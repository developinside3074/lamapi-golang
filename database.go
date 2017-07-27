package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func initDB() *gorm.DB {
	// Openning file
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open("postgres", connectionString)

	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}

	// Creating the table
	if !db.HasTable(&Users{}) {
		db.CreateTable(&Users{})
		db.Set("gorm:table_options", "OIDS = FALSE").CreateTable(&Users{})
	}

	return db
}
