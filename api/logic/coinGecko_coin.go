package logic

import (
	"liliya/models"
	"net/http"
)

func (l Logic) GetInfoCoinsForPage(id, page string) ([]models.Coin, int) {
	output, err := l.plugins.GetInfoCoins(id, page)
	if err != nil {
		return nil, http.StatusBadRequest
	}

	return output, http.StatusOK
}
