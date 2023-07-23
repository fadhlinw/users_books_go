package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"tugas/config"
	"tugas/models"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetBooksController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/books", nil)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)

	if assert.NoError(t, GetBooksController(c)) {
		assert.Equal(t, http.StatusOK, record.Code)
	}
}

func TestGetBookController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/books/3", nil)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("3")

	if assert.NoError(t, GetBookController(c)) {
		assert.Equal(t, http.StatusOK, record.Code)
	}

	request = httptest.NewRequest(http.MethodGet, "/books/sherlok", nil)
	record = httptest.NewRecorder()
	c = e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("sherlok")
	error := GetBookController(c)

	if assert.Error(t, GetBookController(c)) {
		e, ok := error.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusBadRequest, e.Code)
		assert.Equal(t, "Invalid ID", e.Message)
	}

	request = httptest.NewRequest(http.MethodGet, "/books/100", nil)
	record = httptest.NewRecorder()
	c = e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("100")
	err := GetBookController(c)

	if assert.Error(t, GetBookController(c)) {
		e, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusNotFound, e.Code)
		assert.Equal(t, "Book Not Found", e.Message)
	}
}

func TestCreateBookController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	book := models.Book{
		Judul:     "fadhli test 9",
		Pengarang: "fadhlitest 9",
		Penerbit:  "Gramedia",
	}
	bookJSON, _ := json.Marshal(book)
	request := httptest.NewRequest(http.MethodPost, "/create-book", strings.NewReader(string(bookJSON)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)

	if assert.NoError(t, CreateBookController(c)) {
		assert.Equal(t, http.StatusCreated, record.Code)
	}
}

func TestDeleteBookController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	request := httptest.NewRequest(http.MethodDelete, "/books/14", nil)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("14")

	if assert.NoError(t, DeleteBookController(c)) {
		assert.Equal(t, http.StatusOK, record.Code)
	}

	request = httptest.NewRequest(http.MethodDelete, "/books/delete", nil)
	record = httptest.NewRecorder()
	c = e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("delete")
	error := DeleteBookController(c)

	if assert.Error(t, DeleteBookController(c)) {
		response, ok := error.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "Invalid ID", response.Message)
	}
}

func TestUpdateBookController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	book := models.Book{
		Judul:     "Test Book Updated 9",
		Pengarang: "fadhli test",
		Penerbit:  "Gramedia",
	}
	bookJSON, _ := json.Marshal(book)
	request := httptest.NewRequest(http.MethodPut, "/books/3", strings.NewReader(string(bookJSON)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("3")

	if assert.NoError(t, UpdateBookController(c)) {
		assert.Equal(t, http.StatusOK, record.Code)
	}

	request = httptest.NewRequest(http.MethodPut, "/books/update", strings.NewReader(string(bookJSON)))
	record = httptest.NewRecorder()
	c = e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("update")
	error := UpdateBookController(c)

	if assert.Error(t, UpdateBookController(c)) {
		response, ok := error.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "Invalid Request Payload", response.Message)
	}
}
