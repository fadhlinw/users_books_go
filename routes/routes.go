package routes

import (
	"tugas/controllers"
	"tugas/middleware"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	middleware.LogMiddleware(e)

	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/create-user", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/create-book", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)

	return e
}
