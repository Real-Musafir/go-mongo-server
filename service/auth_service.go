package service

import (
	"fmt"

	"github.com/Real-Musafir/bookshop/dto"
	"github.com/Real-Musafir/bookshop/model"
	repo "github.com/Real-Musafir/bookshop/repository"
	"go.mongodb.org/mongo-driver/mongo"
)


type IAuthService interface {
	Login(loginDto dto.LoginDto, sessionContext mongo.SessionContext) (map[string]interface{}, error)
}

type AuthService struct  {
	repository repo.Repository
	userService IUserService
}

func (as *AuthService) Login(loginDto dto.LoginDto, sessionContext mongo.SessionContext) (map[string]interface{}, error) {
	res, err := as.userService.FindOneUserByEmail(loginDto.Email, sessionContext)
	if err != nil {
		return nil, err
	}

	user := res.(model.User)

	if user.Password == loginDto.Password {
		return map[string]any{
			"user":user,
			"token": "",
		}, nil
	}


	return nil, fmt.Errorf("Invalid credentials")

}

func GetAuthService(repository repo.Repository, userService IUserService) IAuthService {
	return &AuthService{
		repository: repository,
		userService: userService,
	}
}