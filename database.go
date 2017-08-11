package main

import (
	"lamapi/models"

	"github.com/go-pg/pg"
)

func initDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "test",
	})
	return db
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{&models.Task{}} {
		err := db.CreateTable(model, nil)
		if err != nil {
			return err
		}
	}
	return nil
}
