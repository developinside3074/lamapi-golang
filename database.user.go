package main

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	Firstname string `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname  string `gorm:"not null" form:"lastname" json:"lastname"`
}

func initDbUser(db *gorm.DB) {
	// Creating the table
	if !db.HasTable(&Users{}) {
		db.CreateTable(&Users{})
	}
}
