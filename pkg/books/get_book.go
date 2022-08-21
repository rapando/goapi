package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rapando/goapi/pkg/common/models"
)

func (h handler) GetBook(c *gin.Context) {
	var id = c.Param("id")
	var book models.Book

	var result = h.DB.First(&book, id)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &book)
}
