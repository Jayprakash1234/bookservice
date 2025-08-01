package handlers

import (
	"net/http"

	"github.com/Jayprakash1234/bookservice/database"
	"github.com/Jayprakash1234/bookservice/models"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

func AddBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//book.ID = uuid.New().String()
	database.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}
