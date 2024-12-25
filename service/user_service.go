package service

import (
	"github.com/Real-Musafir/bookshop/model"
	repo "github.com/Real-Musafir/bookshop/repository"
	"github.com/Real-Musafir/bookshop/utils"
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

func (us *UserService) FindOneUserByEmail(email string, sessionContext mongo.SessionContext) (interface{}, error) {
	res, err := us.repository.FindOneByKey("email", email, sessionContext)

	if err != nil {
		return nil, err
	}

	var user model.User

	if err := utils.MapToStruct(res.(map[string]interface{}), &user); err != nil {
		return nil, err
	}

	return user, nil
}


func GetUserService(repository repo.IMongoRepository) IUserService{
	return &UserService{
		repository: repository,
	}
}