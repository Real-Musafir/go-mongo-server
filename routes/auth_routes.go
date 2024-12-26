package routes

import (
	"github.com/Real-Musafir/bookshop/controller"
	repo "github.com/Real-Musafir/bookshop/repository"
	"github.com/Real-Musafir/bookshop/service"
	"github.com/Real-Musafir/bookshop/utils"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	repository := repo.GetRepository()
	userService := service.GetUserService(*repository)
	authService := service.GetAuthService(*repository, userService)
	responseService := utils.GetResponseService()
	authController := controller.GetAuthController(authService, *responseService)

	router.POST(
		"/login",
		authController.Login,
	)
}