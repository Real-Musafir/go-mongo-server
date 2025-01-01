package service

import (
	"fmt"

	"github.com/Real-Musafir/bookshop/dto"
	"github.com/Real-Musafir/bookshop/model"
	repo "github.com/Real-Musafir/bookshop/repository"
	"github.com/Real-Musafir/bookshop/utils"
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
		return nil, fmt.Errorf("401::%s::%s::%v", "User Not Found", "AuthService_Login", "User does not exist with this email")
	}

	user := res.(model.User)
	token, err := utils.CreateToken(map[string]any{
		"user_id": user.Id.Hex(),
		"user_name": user.Name,
	})

	if err != nil {
		return nil, fmt.Errorf("500::%s::%s::%v", "Internal Server Error", "AuthService_Login", err)
	}



	if user.Password == loginDto.Password {
		return map[string]any{
			"user":user,
			"token": token,
		}, nil
	}


	return nil, fmt.Errorf("401::%s::%s::%v", "Invalid credential", "AuthService_Login", "Password Mismatch")

}

func GetAuthService(repository repo.Repository, userService IUserService) IAuthService {
	return &AuthService{
		repository: repository,
		userService: userService,
	}
}