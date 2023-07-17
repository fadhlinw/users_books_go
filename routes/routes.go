package routes

import (
	"tugas/controllers"
	"tugas/data"
	"tugas/middleware"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	middleware.LogMiddleware(e)

	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/create-user", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.POST("/login", controllers.LoginUserController)

	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/create-book", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)

	eJWT := e.Group("")
	eJWT.Use(mid.JWT([]byte(data.SECRET_JWT)))
	eJWT.GET("/users", controllers.GetUsersController)
	eJWT.GET("/users/:id", controllers.GetUserController)
	eJWT.DELETE("/users/:id", controllers.DeleteUserController)
	eJWT.PUT("/users/:id", controllers.UpdateUserController)
	eJWT.POST("/create-book", controllers.CreateBookController)
	eJWT.DELETE("/books/:id", controllers.DeleteBookController)
	eJWT.PUT("/books/:id", controllers.UpdateBookController)

	return e
}
