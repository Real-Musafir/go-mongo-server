package controller

import (
	"github.com/Real-Musafir/bookshop/model"
	"github.com/Real-Musafir/bookshop/service"
	"github.com/Real-Musafir/bookshop/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService     service.IUserService
	responseService utils.ResponseService
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var dto model.UserCreateDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		// Extract custom error message from the DTO
		customMessage := utils.ExtractCustomErrorMessage(err, &model.UserCreateDto{})
		uc.responseService.Failure(ctx, 400, nil, customMessage)
		return
	}

	data, err := uc.userService.CreateUser(dto, nil)
	if err != nil {
		ctx.Error(err)
		return
	}

	uc.responseService.Success(ctx, 200, data, "Successfully saved!")
}

func GetUserController(userService service.IUserService, responseService utils.ResponseService) *UserController {
	return &UserController{
		userService:     userService,
		responseService: responseService,
	}
}
