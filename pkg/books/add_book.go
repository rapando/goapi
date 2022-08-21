package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rapando/goapi/pkg/common/models"
)

type AddBookRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) AddBook(c *gin.Context) {
	var err error
	var body AddBookRequest
	var book models.Book

	err = c.BindJSON(&body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book = models.Book{
		Title:       body.Title,
		Author:      body.Author,
		Description: body.Description,
	}
	
	var result = h.DB.Create(&book)
	if result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}
	c.JSON(http.StatusCreated, &book)

}
