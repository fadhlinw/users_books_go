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

func TestGetUsersController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/users", nil)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)

	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, record.Code)
	}
}

func TestGetUserController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/users/2", nil)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("2")

	if assert.NoError(t, GetUserController(c)) {
		assert.Equal(t, http.StatusOK, record.Code)
	}

	request = httptest.NewRequest(http.MethodGet, "/users/aku", nil)
	record = httptest.NewRecorder()
	c = e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("aku")
	error := GetUserController(c)

	if assert.Error(t, GetBookController(c)) {
		response, ok := error.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "Invalid ID", response.Message)
	}

	request = httptest.NewRequest(http.MethodGet, "/users/100", nil)
	record = httptest.NewRecorder()
	c = e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("100")
	err := GetUserController(c)

	if assert.Error(t, GetBookController(c)) {
		response, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, "User Tidak Ada", response.Message)
	}
}

func TestLoginUserController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	user := models.User{
		Email:    "fadhlitest12@gmail.com",
		Password: "12345",
	}
	userJSON, _ := json.Marshal(user)
	request := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(userJSON)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)

	if assert.NoError(t, LoginUserController(c)) {
		assert.Equal(t, http.StatusCreated, record.Code)
	} else {
		assert.Equal(t, http.StatusBadRequest, record.Code)
	}
}

func TestCreateUserController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	user := models.User{
		Name:     "fadhli test 9",
		Email:    "fadhlitest123456789101112@gmail.com",
		Password: "12345",
	}
	userJSON, _ := json.Marshal(user)
	request := httptest.NewRequest(http.MethodPost, "/create-user", strings.NewReader(string(userJSON)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)

	if assert.NoError(t, CreateUserController(c)) {
		assert.Equal(t, http.StatusCreated, record.Code)
	}
}

func TestDeleteUserController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	request := httptest.NewRequest(http.MethodDelete, "/users/29", nil)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("29")

	if assert.NoError(t, DeleteUserController(c)) {
		assert.Equal(t, http.StatusOK, record.Code)
	}

	request = httptest.NewRequest(http.MethodDelete, "/users/delete", nil)
	record = httptest.NewRecorder()
	c = e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("delete")
	error := DeleteUserController(c)

	if assert.Error(t, DeleteBookController(c)) {
		response, ok := error.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "Invalid ID", response.Message)
	}
}

func TestUpdateUserController(t *testing.T) {

	config.InitDB()
	e := echo.New()
	user := models.User{
		Name:     "Test User Updated 10",
		Email:    "fadhlitest12@gmail.com",
		Password: "12345",
	}
	userJSON, _ := json.Marshal(user)
	request := httptest.NewRequest(http.MethodPut, "/users/2", strings.NewReader(string(userJSON)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	record := httptest.NewRecorder()
	c := e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("2")

	if assert.NoError(t, UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, record.Code)
	}

	request = httptest.NewRequest(http.MethodPut, "/users/update", strings.NewReader(string(userJSON)))
	record = httptest.NewRecorder()
	c = e.NewContext(request, record)
	c.SetParamNames("id")
	c.SetParamValues("update")
	error := UpdateUserController(c)

	if assert.Error(t, UpdateBookController(c)) {
		response, ok := error.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "Invalid request payload", response.Message)
	}
}
