package coinGecko

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
	"liliya/models"
	"net/http"
)

type CoinGeckoMarketPlugin struct {
	api string
}

func NewCoinGeckoMarketPlugin(api string) *CoinGeckoMarketPlugin {
	return &CoinGeckoMarketPlugin{api: api}
}

func (s CoinGeckoMarketPlugin) GetAllTMarketsByID(id string) ([]string, error) {
	var tickers *models.Tickers
	url := fmt.Sprintf(`https://api.coingecko.com/api/v3/coins/%s/tickers`, id)

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

	err = json.NewDecoder(resp.Body).Decode(&tickers)
	if err != nil {
		logrus.Error(`error decoding JSON: %s`, err)
		return nil, err
	}

	if len(tickers.Tickers) == 0 {
		logrus.Error(`no tickers found`)
		return nil, fmt.Errorf(`no tickers found`)
	}

	var marketsName []string
	for _, ticker := range tickers.Tickers {
		marketsName = append(marketsName, ticker.Market.Name)
	}

	return marketsName, nil
}
