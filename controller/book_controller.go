package controller

import (
	"fmt"

	"github.com/Real-Musafir/bookshop/dto"
	"github.com/Real-Musafir/bookshop/service"
	"github.com/Real-Musafir/bookshop/utils"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService service.IBookService
	responseService utils.ResponseService
}

func (bc *BookController) CreateBook(ctx *gin.Context){
	var bookDto dto.CreateBookDto
	if err := ctx.ShouldBindJSON(&bookDto); err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", "BookController_CreateBook", err))
		return
	}

	data, err := bc.bookService.CreateBook(bookDto, nil)
	if err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Data saving failed", "BookController_CreateBook", err))
		return
	}

	bc.responseService.Success(ctx, 201, data, "Successfully saved!")

}

func (bc *BookController) GetAllBooks(ctx *gin.Context){
	data, err := bc.bookService.GetAllBooks(nil)
	if err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Data getting failed", "BookController_GetAllBooks", err))
		return
	}

	bc.responseService.Success(ctx, 200, data, "fetched successfully!")
}

func (bc *BookController) GetBookById(ctx *gin.Context){
	var bookId = ctx.Param("bookId")
	if bookId == "" {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", "BookController_GetAllBooks", "Book Id is not present in request param"))
		return
	}

	data, err := bc.bookService.GetBookById(bookId, nil)
	if err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", err.Error(), "BookController_GetBookById", err))
		return
	}

	bc.responseService.Success(ctx, 200, data, "fetched successfully!!!")
}

func (bc *BookController) UpdateBookById(ctx *gin.Context){
	var bookDto dto.UpdateBookDto
	if err := ctx.ShouldBindJSON(&bookDto); err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", "BookController_UpdateBookById", err))
		return
	}

	data, err := bc.bookService.UpdateBookById(bookDto, nil)
	if err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Data saving failed", "BookController_UpdateBookById", err))
		return
	}

	bc.responseService.Success(ctx, 201, data, "Successfully updated!")
}

func (bc *BookController) DeleteBookById(ctx *gin.Context){
	var bookId = ctx.Param("bookId")
	if bookId == "" {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", "BookController_DeleteBookById", "Book Id is not present in request param"))
		return
	}

	data, err := bc.bookService.DeleteBookById(bookId, nil)
	if err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Data fetching failed", "BookController_DeleteBookById", err))
		return
	}

	bc.responseService.Success(ctx, 200, data, "Deleted successfully!")
}

func GetBookController(bookService service.IBookService, responseService utils.ResponseService) *BookController {
	return &BookController{
		bookService: bookService,
		responseService: responseService,
	}
}




