package routes

import (
	"github.com/Real-Musafir/bookshop/controller"
	"github.com/Real-Musafir/bookshop/middleware"
	repo "github.com/Real-Musafir/bookshop/repository"
	"github.com/Real-Musafir/bookshop/service"
	"github.com/Real-Musafir/bookshop/utils"
	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(router *gin.RouterGroup) {
	repository := repo.GetRepository()
	bookService := service.GetBookService(*repository)

	responseService := utils.GetResponseService()
	bookController := controller.GetBookController(bookService, *responseService)

	router.POST(
		"/create",
		middleware.AuthenticateRequest(false),
		bookController.CreateBook,
	)

	router.POST(
		"/",
		middleware.AuthenticateRequest(false),
		bookController.GetAllBooks,
	)

	router.GET(
		"/:bookId",
		middleware.AuthenticateRequest(false),
		bookController.GetBookById,
	)

	router.POST(
		"/:bookId/update",
		middleware.AuthenticateRequest(false),
		bookController.UpdateBookById,
	)

	router.DELETE(
		"/:bookId/delete",
		middleware.AuthenticateRequest(false),
		bookController.DeleteBookById,
	)
}