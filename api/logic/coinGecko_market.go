package logic

import (
	"net/http"
)

func (l Logic) GetAllTMarketsByID(id string) ([]string, int) {
	output, err := l.plugins.GetAllTMarketsByID(id)
	if err != nil {
		return nil, http.StatusBadRequest
	}

	return output, http.StatusOK
}
