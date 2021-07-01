package service

import (
	_struct "github.com/PutskouDzmitry/golang-training-Library/pkg/entity"
	"github.com/PutskouDzmitry/golang-training-Library/pkg/repository"
)

type BookService struct {
	repo repository.BooksRepo
}

func (b BookService) ReadAll() ([]_struct.Book, error) {
	return b.repo.ReadAll()
}

func (b BookService) Read(id string) (_struct.Book, error) {
	return b.repo.Read(id)
}

func (b BookService) Add(book _struct.Book) (string, error) {
	return b.repo.Add(book)
}

func (b BookService) Update(id string, value int) (string, error) {
	return b.repo.Update(id, value)
}

func (b BookService) Delete(id string) error {
	return b.repo.Delete(id)
}

func NewBookService(repo repository.BooksRepo) *BookService {
	return &BookService{repo: repo}
}
