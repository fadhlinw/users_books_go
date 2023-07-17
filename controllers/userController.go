package controllers

import (
	"net/http"
	"strconv"

	"tugas/database"
	"tugas/middleware"
	"tugas/models"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	user, err := database.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User Tidak Ada")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func LoginUserController(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	err := database.LoginUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	token, err := middleware.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	UserResponse := models.UserResponse{int(user.ID), user.Name, user.Email, token}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"user":   UserResponse,
	})
}

func CreateUserController(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	err := database.GetUserbyEmail(user.Email)
	if err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Email Sudah Terdaftar")
	}

	err = database.CreateUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	err = database.DeleteUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User Tidak Ada")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Success",
		"message": "User Deleted",
	})
}

func UpdateUserController(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	err = database.UpdateUser(id, &user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Success",
		"message": "User Updated",
	})
}
