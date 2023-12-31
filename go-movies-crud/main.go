package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"errors"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In search of Lost time ", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby ", Author: "F.Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace ", Author: "Leo Tolstoy ", Quantity: 6},
	{ID: "4", Title: "Lolita", Author: "Vladimir Nabokov", Quantity: 9},
	{ID: "5", Title: "Middlemarch ", Author: "George Eliot", Quantity: 3},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)

}
func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}
