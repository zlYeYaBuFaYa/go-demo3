package router

import (
	"go-demo3/internal/app/v1"
	"github.com/gin-gonic/gin"
)

func registerBookRoutes(r *gin.Engine) {
	books := r.Group("/api/books")
	{
		books.GET("", v1.ListBooks)
		books.GET(":id", v1.GetBook)
		books.POST("", v1.CreateBook)
		books.PUT(":id", v1.UpdateBook)
		books.DELETE(":id", v1.DeleteBook)
	}
}
