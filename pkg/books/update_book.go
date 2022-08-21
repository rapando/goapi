package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rapando/goapi/pkg/common/models"
)

type UpdateBookRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) UpdateBook(c *gin.Context) {
	var err error
	var id = c.Param("id")
	var body UpdateBookRequest
	var book models.Book

	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var result = h.DB.First(&book, id)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	h.DB.Save(&book)

	c.JSON(http.StatusOK, &book)
}
