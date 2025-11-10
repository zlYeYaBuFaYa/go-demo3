package v1

import (
	"go-demo3/internal/global"
	"go-demo3/internal/models"
	"go-demo3/internal/response"
	"go-demo3/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var bookService = service.NewBookService()

func ListBooks(c *gin.Context) {
	books, err := bookService.ListBooks()
	if err != nil {
		global.LogS.Error("图书查询失败", zap.Error(err))
		response.FailResponse(c, "查询失败")
		return
	}
	global.LogS.Info("图书查询成功", zap.Error(err))
	response.SuccessResponse(c, books)
}

func GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailResponse(c, "参数错误")
		return
	}
	book, err := bookService.GetBook(uint(id))
	if err != nil {
		response.FailResponse(c, "未找到图书")
		return
	}
	response.SuccessResponse(c, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		response.FailResponse(c, "参数错误")
		return
	}
	if err := bookService.CreateBook(&book); err != nil {
		response.FailResponse(c, "添加失败")
		return
	}
	response.SuccessResponse(c, book)
}

func UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailResponse(c, "参数错误")
		return
	}
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		response.FailResponse(c, "参数错误")
		return
	}
	book.ID = uint(id)
	if err := bookService.UpdateBook(&book); err != nil {
		response.FailResponse(c, "更新失败")
		return
	}
	response.SuccessResponse(c, book)
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailResponse(c, "参数错误")
		return
	}
	if err := bookService.DeleteBook(uint(id)); err != nil {
		response.FailResponse(c, "删除失败")
		return
	}
	response.SuccessResponse(c, nil)
}
