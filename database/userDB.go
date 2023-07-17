package database

import (
	"tugas/config"
	"tugas/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id int) (*models.User, error) {
	user := &models.User{}
	if err := config.DB.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserbyEmail(email string) error {
	user := &models.User{}
	if err := config.DB.Where("email = ?", user.Email).First(user).Error; err != nil {
		return err
	}
	return nil
}
func CreateUser(user *models.User) error {

	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func LoginUser(user *models.User) error {
	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	if err := config.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}
func UpdateUser(id int, user *models.User) error {
	if err := config.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}
