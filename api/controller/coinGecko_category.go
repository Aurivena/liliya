package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (con Controller) getAllCategoriesById(c *gin.Context) {
	id := c.Param("id")

	output, processStatus := con.logic.GetCategoriesByID(id)
	if processStatus != http.StatusOK {
		c.JSON(processStatus, nil)
		return
	}

	c.JSON(processStatus, output)
}

func (con Controller) getAllCategories(c *gin.Context) {
	output, processStatus := con.logic.GetCategories()
	if processStatus != http.StatusOK {
		c.JSON(processStatus, nil)
		return
	}

	c.JSON(processStatus, output)
}
