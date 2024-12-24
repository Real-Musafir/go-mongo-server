package service

import (
	repo "github.com/Real-Musafir/bookshop/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserService interface {
	CreateUser(data interface{}, sessionContext mongo.SessionContext) (interface{}, error)
}

type UserService struct {
	repository repo.IMongoRepository
}

func (us *UserService) CreateUser(data interface{}, sessionContext mongo.SessionContext) (interface{}, error) {
	return us.repository.Create(data, sessionContext)
}