package handlers

import (
	"lamapi/models"
	"net/http"
	"strconv"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
)

// GetTasks endpoint
func GetTasks(db *pg.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := models.GetTasks(db)
		if err == nil {
			return c.JSON(http.StatusCreated, echo.Map{
				"data": res,
			})
		}
		return err
	}
}

// PutTask endpoint
func PutTask(db *pg.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var task models.Task
		// Map imcoming JSON body to the new Task
		c.Bind(&task)
		// Add a task using our new model
		res, err := models.PutTask(db, task.Name)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, echo.Map{
				"created": res,
			})
		}
		return err
	}
}

// DeleteTask endpoint
func DeleteTask(db *pg.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// Use our new model to delete a task
		res, err := models.DeleteTask(db, id)
		// Return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, echo.Map{
				"deleted": res,
			})
		}
		return err
	}
}
