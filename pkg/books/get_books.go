package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rapando/goapi/pkg/common/models"
	"github.com/rapando/goapi/pkg/common/utils"
)

func (h handler) GetBooks(c *gin.Context) {
	var books []models.Book
	var result = h.DB.Find(&books)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	utils.Logger.Info("found all books")
	c.JSON(http.StatusOK, &books)
}
