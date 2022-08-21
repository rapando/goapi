package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rapando/goapi/pkg/common/models"
)

func (h handler) DeleteBook(c *gin.Context) {
	var id = c.Param("id")
	var book models.Book

	var result = h.DB.First(&book, id)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&book)
	c.Status(http.StatusOK)
}
