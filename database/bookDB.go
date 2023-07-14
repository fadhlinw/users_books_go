package database

import (
	"tugas/config"
	"tugas/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBook(id int) (*models.Book, error) {
	book := &models.Book{}
	if err := config.DB.First(book, id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func CreateBook(book *models.Book) error {
	if err := config.DB.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBook(id int) error {
	book := &models.Book{}
	if err := config.DB.Delete(book, id).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBook(id int, book *models.Book) error {
	if err := config.DB.Model(&models.Book{}).Where("id = ?", id).Updates(book).Error; err != nil {
		return err
	}
	return nil
}
