package coinGecko

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
	"liliya/models"
	"net/http"
)

type CoinGeckoCategoryPlugin struct {
	api string
}

func NewCoinGeckoCategoryPlugin(api string) *CoinGeckoCategoryPlugin {
	return &CoinGeckoCategoryPlugin{api: api}
}

func (s CoinGeckoCategoryPlugin) GetCategoriesByID(id string) ([]string, error) {

	var categories *models.CategoryNames

	url := fmt.Sprintf(`https://api.coingecko.com/api/v3/coins/%s`, id)

	resp, err := http.Get(url)
	if err != nil {
		logrus.Error(`error from http.Get():%d`, err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		logrus.Error(`API returned non-200 status: %s`, resp.Status)
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&categories)
	if err != nil {
		logrus.Error(`error decoding JSON: %s`, err)
		return nil, err
	}

	if len(categories.Names) == 0 {
		logrus.Error(`no categories found`)
		return nil, fmt.Errorf(`error unmarshalling categories:`, err)
	}

	var categoriesNames []string
	for _, category := range categories.Names {
		categoriesNames = append(categoriesNames, category)
	}

	return categoriesNames, nil
}

//https://api.coingecko.com/api/v3/coins/categories/list

func (s CoinGeckoCategoryPlugin) GetCategories() ([]models.Category, error) {
	var categories []models.Category

	url := fmt.Sprintf(`https://api.coingecko.com/api/v3/coins/categories/list`)

	resp, err := http.Get(url)
	if err != nil {
		logrus.Error(`error from http.Get():%d`, err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		logrus.Error(`API returned non-200 status: %s`, resp.Status)
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&categories)
	if err != nil {
		logrus.Error(`error decoding JSON: %s`, err)
		return nil, err
	}

	if len(categories) == 0 {
		logrus.Error(`no categories found`)
		return nil, fmt.Errorf(`error unmarshalling categories:`, err)
	}

	return categories, nil
}
