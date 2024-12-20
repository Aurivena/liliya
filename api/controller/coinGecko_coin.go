package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (con Controller) getAllTokens(c *gin.Context) {
	currency := c.Param("currency")
	page := c.Param("page")

	output, processStatus := con.logic.GetInfoCoinsForPage(currency, page)
	if processStatus != http.StatusOK {
		c.JSON(processStatus, nil)
		return
	}

	c.JSON(processStatus, output)
}
