package store

import (
	"go-demo3/internal/global"
	"go-demo3/internal/models"
)

type BookStore struct{}

func NewBookStore() *BookStore { return &BookStore{} }

func (s *BookStore) Create(book *models.Book) error {
	return global.DB.Create(book).Error
}

func (s *BookStore) Get(id uint) (*models.Book, error) {
	var book models.Book
	err := global.DB.Where("id = ? AND removed = ?", id, false).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *BookStore) List() ([]models.Book, error) {
	var books []models.Book
	err := global.DB.Where("removed = ?", false).Order("id desc").Find(&books).Error
	return books, err
}

func (s *BookStore) Update(book *models.Book) error {
	return global.DB.Model(book).Updates(book).Error
}

func (s *BookStore) Delete(id uint) error {
	return global.DB.Model(&models.Book{}).Where("id = ?", id).Update("removed", true).Error
}
