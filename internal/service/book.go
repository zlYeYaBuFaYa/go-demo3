package service

import (
	"go-demo3/internal/models"
	"go-demo3/internal/store"
)

type BookService struct {
	store *store.BookStore
}

func NewBookService() *BookService {
	return &BookService{store: store.NewBookStore()}
}

func (s *BookService) CreateBook(book *models.Book) error {
	return s.store.Create(book)
}

func (s *BookService) GetBook(id uint) (*models.Book, error) {
	return s.store.Get(id)
}

func (s *BookService) ListBooks() ([]models.Book, error) {
	return s.store.List()
}

func (s *BookService) UpdateBook(book *models.Book) error {
	return s.store.Update(book)
}

func (s *BookService) DeleteBook(id uint) error {
	return s.store.Delete(id)
}
