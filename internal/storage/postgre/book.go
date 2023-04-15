package postgre

import (
	"github.com/alibekabdrakhman1/first-lab/internal/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func (r *BookRepository) GetByAuthor(author string) ([]model.Book, error) {
	var all []model.Book
	if err := r.db.Find(&all).Error; err != nil {
		return nil, err
	}
	ans := make([]model.Book, 0)
	for _, book := range all {
		if book.Author == author {
			ans = append(ans, book)
		}
	}
	return ans, nil
}

func (r *BookRepository) GetByName(name string) (model.Book, error) {
	var ans model.Book
	err := r.db.Where("name = ?", name).First(&ans).Error
	if err != nil {
		return model.Book{}, err
	}
	return ans, nil
}

func (r *BookRepository) Create(book model.Book) error {
	if err := r.db.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func (r *BookRepository) Update(book model.Book) error {
	return r.db.Where("name = ?", book.Name).Updates(&book).Error
}

func (r *BookRepository) Delete(name string) error {
	return r.db.Delete(&model.Book{}, "name = ?", name).Error
}

func (r *BookRepository) GetAll() ([]model.Book, error) {
	var resp []model.Book
	err := r.db.Find(&resp)
	return resp, err.Error
}
func (r *BookRepository) GetAvailableBooks() ([]model.Book, error) {
	var all []model.Book
	if err := r.db.Find(&all).Error; err != nil {
		return nil, err
	}
	ans := make([]model.Book, 0)
	for _, book := range all {
		if book.Available == true {
			ans = append(ans, book)
		}
	}
	return ans, nil
}

func NewBookRepository(DB *gorm.DB) *BookRepository {
	return &BookRepository{
		db: DB,
	}
}
