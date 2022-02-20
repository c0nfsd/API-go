package main

import (
	"example/API-go/helper"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/books", helper.GetBooks)
	router.GET("/books/:id", helper.BookById)
	router.POST("/books", helper.CreateBook)
	router.PATCH("/checkout", helper.CheckoutBook)
	router.PATCH("/return", helper.ReturnBook)
	router.Run("localhost:8080")

}
