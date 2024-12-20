package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (con Controller) getAllMarketById(c *gin.Context) {
	id := c.Param("id")

	output, processStatus := con.logic.GetAllTMarketsByID(id)
	if processStatus != http.StatusOK {
		c.JSON(processStatus, nil)
		return
	}

	c.JSON(processStatus, output)
}
