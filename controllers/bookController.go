package controllers

import (
	"net/http"
	"strconv"
	"tugas/database"
	"tugas/models"

	"github.com/labstack/echo"
)

func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success",
		"books":  books,
	})
}

func GetBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	book, err := database.GetBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Book Not Found")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success",
		"Books":  book,
	})
}

func CreateBookController(c echo.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Request Payload")
	}
	err := database.CreateBook(&book)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "Success",
		"message": book,
	})
}

func DeleteBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = database.DeleteBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Book Not Found")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Success",
		"Message": "Book Deleted",
	})
}

func UpdateBookController(c echo.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Request Payload")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = database.UpdateBook(id, &book)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Success",
		"message": "Book Updated",
	})
}
