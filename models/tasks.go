package models

import (
	"github.com/go-pg/pg"
)

// Task is a struct containing Task data
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetTasks get all tasks data
func GetTasks(db *pg.DB) ([]Task, error) {
	var tasks []Task
	err := db.Model(&tasks).Select()
	if err != nil {
		panic(err)
	}
	return tasks, err
}

// PutTask create new task data
func PutTask(db *pg.DB, name string) (*Task, error) {
	task := &Task{
		Name: name,
	}
	err := db.Insert(task)
	if err != nil {
		panic(err)
	}
	return task, err
}

// DeleteTask remove a task
func DeleteTask(db *pg.DB, id int) (interface{}, error) {
	var task Task
	res, err := db.Model(&task).Where("title = ?title").Delete()
	if err != nil {
		panic(err)
	}
	return res, err
}
