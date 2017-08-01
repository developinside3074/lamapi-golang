package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Initial Database
	initDB()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS default
	// Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
	e.Use(middleware.CORS())

	// CORS restricted
	// Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	// wth GET, PUT, POST or DELETE method.
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
	// 	AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	// }))

	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	// Routes
	// -- Auth route
	e.POST("/login", login)

	e.POST("/users", createUser, middleware.JWTWithConfig(config))
	e.GET("/users", getUsers)
	e.GET("/user/:id", getUser)
	e.PUT("/user/:id", updateUser, middleware.JWTWithConfig(config))
	e.DELETE("/user/:id", deleteUser, middleware.JWTWithConfig(config))

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
