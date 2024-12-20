package logic

import (
	"liliya/models"
	"net/http"
)

func (l Logic) GetCategoriesByID(id string) ([]string, int) {
	output, err := l.plugins.GetCategoriesByID(id)
	if err != nil {
		return nil, http.StatusBadRequest
	}

	return output, http.StatusOK
}

func (l Logic) GetCategories() ([]models.Category, int) {
	output, err := l.plugins.GetCategories()
	if err != nil {
		return nil, http.StatusBadRequest
	}

	return output, http.StatusOK
}
