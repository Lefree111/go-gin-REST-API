package controller

import (
	"errors"
	"github.com/Lefree111/go-gin-rest-api/go-crud-api/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateBook(c *gin.Context) {
	var book *database.Book
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	res := database.DB.Create(book)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a book",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
	return

}

func ReadBook(c *gin.Context) {
	var book database.Book
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)
	res := database.DB.Find(&book, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "book not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
	return
}

func ReadBooks(c *gin.Context) {
	var books []database.Book
	res := database.DB.Find(&books)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("book not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
	return
}

func UpdateBook(c *gin.Context) {
	var book database.Book
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)
	err := c.ShouldBind(&book)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updateBook database.Book
	res := database.DB.Model(&book).Where("id = ?", id).Updates(updateBook)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "book not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
	return
}

func DeleteBook(c *gin.Context) {
	var book database.Book
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)
	res := database.DB.Find(&book, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "book not found",
		})
		return
	}
	database.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{
		"message": "book deleted successfully",
	})
	return
}
