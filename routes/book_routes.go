package routes

import (
	"github.com/Real-Musafir/bookshop/controller"
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
		bookController.CreateBook,
	)

	router.POST(
		"/",
		bookController.GetAllBooks,
	)

	router.GET(
		"/:bookId",
		bookController.GetBookById,
	)

	router.POST(
		"/:bookId/update",
		bookController.UpdateBookById,
	)

	router.DELETE(
		"/:bookId/delete",
		bookController.DeleteBookById,
	)
}