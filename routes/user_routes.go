package routes

import (
	"github.com/Real-Musafir/bookshop/controller"
	repo "github.com/Real-Musafir/bookshop/repository"
	"github.com/Real-Musafir/bookshop/service"
	"github.com/Real-Musafir/bookshop/utils"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	repository := repo.GetRepository().UserRepository
	userService := service.GetUserService(repository)
	responseService := utils.GetResponseService()
	userController := controller.GetUserController(userService, *responseService)

	router.POST(
		"/add",
		userController.CreateUser,
	)
}