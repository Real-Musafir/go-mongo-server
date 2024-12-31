package service

import (
	repo "github.com/Real-Musafir/bookshop/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type IBookService interface {
	CreateBook(data interface{}, sessionContext mongo.SessionContext) (interface{}, error)
	GetAllBooks(sessionContext mongo.SessionContext) (interface{}, error)
	GetBookById(id string, sessionContext mongo.SessionContext) (interface{}, error)
	UpdateBookById(data interface{}, sessionContext mongo.SessionContext) (interface{}, error)
	DeleteBookById(id string, sessionContext mongo.SessionContext) (interface{}, error)
}

type BookService struct {
	repository repo.Repository
}

func (bs *BookService) CreateBook(data interface{}, sessionContext mongo.SessionContext) (interface{}, error) {
	return bs.repository.BookRepository.Create(data, sessionContext)
}

func (bs *BookService) GetAllBooks(sessionContext mongo.SessionContext) (interface{}, error) {
	return bs.repository.BookRepository.FindAll(nil, sessionContext)
}

func (bs *BookService) GetBookById(id string, sessionContext mongo.SessionContext) (interface{}, error) {
	return bs.repository.BookRepository.FindOne(id, sessionContext)
}

func (bs *BookService) UpdateBookById(data interface{}, sessionContext mongo.SessionContext) (interface{}, error) {
	bookData := data.(map[string]interface{})
	id := bookData["id"].(string)
	delete(bookData, "_id") // form bookData we need to remove id after getting the id
	return bs.repository.BookRepository.Update(id, bookData, sessionContext)
}

func (bs *BookService) DeleteBookById(id string, sessionContext mongo.SessionContext) (interface{}, error) {
	return bs.repository.BookRepository.Delete(id, sessionContext)
}

func GetBookService(repository repo.Repository) IBookService {
	return &BookService{
		repository: repository,
	}
}