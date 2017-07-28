package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func createUser(c echo.Context) error {
	db := initDB()
	defer db.Close()

	var user Users
	c.Bind(&user)

	if user.Firstname != "" && user.Lastname != "" {
		// INSERT INTO "users" (name) VALUES (user.Name);
		db.Create(&user)
		// Display JSON result
		return c.JSON(http.StatusCreated, user)
	} else {
		// Display error
		return c.JSON(http.StatusUnprocessableEntity, "Fields are empty")
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
}

func getUsers(c echo.Context) error {
	// Connection to the database
	db := initDB()
	// Close connection database
	defer db.Close()

	var users []Users
	// SELECT * FROM users
	db.Find(&users)

	// Display JSON result
	return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	// Connection to the database
	db := initDB()
	// Close connection database
	defer db.Close()

	id, _ := strconv.Atoi(c.Param("id"))
	var user Users
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.ID != 0 {
		// Display JSON result
		return c.JSON(http.StatusOK, user)
	} else {
		// Display JSON error
		return c.JSON(http.StatusNotFound, "User not found")
	}
}

func updateUser(c echo.Context) error {
	// Connection to the database
	db := initDB()
	// Close connection database
	defer db.Close()

	// Get id user
	id, _ := strconv.Atoi(c.Param("id"))
	var user Users
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.ID != 0 {

		var newUser Users
		c.Bind(&newUser)

		if newUser.Firstname != "" && newUser.Lastname != "" {
			// Update field
			user.Firstname = newUser.Firstname
			user.Lastname = newUser.Lastname

			// UPDATE users SET firstname='newUser.Firstname', lastname='newUser.Lastname' WHERE id = user.Id;
			db.Save(&user)
			// Display modified data in JSON message "success"
			return c.JSON(http.StatusOK, user)
		} else {
			// Display JSON error
			return c.JSON(http.StatusUnprocessableEntity, "Fields are empty")
		}

	} else {
		// Display JSON error
		return c.JSON(http.StatusNotFound, "User not found")
	}
}

func deleteUser(c echo.Context) error {
	// Connection to the database
	db := initDB()
	// Close connection database
	defer db.Close()

	// Get id user
	id, _ := strconv.Atoi(c.Param("id"))
	var user Users
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.ID != 0 {
		// DELETE FROM users WHERE id = user.Id
		db.Delete(&user)
		// Display JSON result
		return c.JSON(http.StatusOK, id)
	} else {
		// Display JSON error
		return c.JSON(http.StatusNotFound, "User not found")
	}
}
