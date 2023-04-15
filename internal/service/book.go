package service

import (
	"github.com/alibekabdrakhman1/first-lab/internal/model"
	"github.com/alibekabdrakhman1/first-lab/internal/storage"
	"github.com/google/uuid"
)

type IBookService interface {
	GetByAuthor(author string) ([]model.Book, error)
	GetByName(name string) (model.Book, error)
	Create(book model.Book) error
	Update(book model.Book) error
	Delete(name string) error
	GetAll() ([]model.Book, error)
	GetAvailableBooks() ([]model.Book, error)
}

type BookService struct {
	repo *storage.Storage
}

func (s *BookService) GetAvailableBooks() ([]model.Book, error) {
	return s.repo.Book.GetAvailableBooks()
}

func (s *BookService) GetByAuthor(author string) ([]model.Book, error) {
	return s.repo.Book.GetByAuthor(author)
}

func (s *BookService) GetByName(name string) (model.Book, error) {
	return s.repo.Book.GetByName(name)
}

func (s *BookService) Create(book model.Book) error {
	book.ID = uuid.New().String()
	return s.repo.Book.Create(book)
}

func (s *BookService) Update(book model.Book) error {
	return s.repo.Book.Update(book)
}

func (s *BookService) Delete(name string) error {
	return s.repo.Book.Delete(name)
}

func (s *BookService) GetAll() ([]model.Book, error) {
	return s.repo.Book.GetAll()
}

func NewBookService(r *storage.Storage) *BookService {
	return &BookService{
		repo: r,
	}
}
