package controller

import (
	"fmt"

	"github.com/Real-Musafir/bookshop/dto"
	"github.com/Real-Musafir/bookshop/service"
	"github.com/Real-Musafir/bookshop/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.IAuthService
	responseService utils.ResponseService
}

func (ac *AuthController) Login(ctx *gin.Context) {
	var loginDto dto.LoginDto
	if err := ctx.ShouldBindJSON(&loginDto); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Bad Request",
			"statusCode": 400,
			"error": err.Error(),
		})

		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", err.Error(), err))
	}

	data, err := ac.authService.Login(loginDto, nil)
	if err != nil {
		ctx.Error(err)
		return
	}

	ac.responseService.Success(ctx, 200, data, "Successfully login!")
}

func GetAuthController(authService service.IAuthService, responseService utils.ResponseService) *AuthController {
	return &AuthController {
		authService: authService,
		responseService: responseService,
	}
}