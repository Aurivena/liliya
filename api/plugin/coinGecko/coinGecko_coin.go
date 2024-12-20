package coinGecko

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"liliya/models"
	"net/http"
)

type CoinsToSearch struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CoinGeckoCoinPlugin struct {
	api string
}

func NewCoinGeckoCoinPlugin(api string) *CoinGeckoCoinPlugin {
	return &CoinGeckoCoinPlugin{api: api}
}

func (s CoinGeckoCoinPlugin) GetInfoCoins(currency, page string) ([]models.Coin, error) {
	var output []models.CoinAllInfo
	url := fmt.Sprintf(`https://api.coingecko.com/api/v3/coins/markets?vs_currency=%s&page=%s`, currency, page)

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

	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		logrus.Error(`error decoding JSON: %s`, err)
		return nil, err
	}

	if len(output) == 0 {
		logrus.Error(`no coins found`)
		return nil, fmt.Errorf(`error unmarshalling coins:`, err)
	}
	var coins []models.Coin

	for _, coin := range output {

		remainingSupply := coin.TotalSupply - coin.CirculatingSupply

		coin := models.Coin{
			Rank:                 coin.Rank,
			Id:                   coin.Id,
			Symbol:               coin.Symbol,
			Image:                coin.Image,
			MarketCapitalization: coin.MarketCapitalization,
			TotalVolume:          coin.TotalVolume,
			Supply: models.Supply{
				CirculatingSupply: coin.CirculatingSupply,
				TotalSupply:       coin.TotalSupply,
				RemainingSupply:   remainingSupply,
			},
			Change: models.Change{
				Price:               coin.Price,
				MarketCap:           coin.MarketCap,
				MarketCapPercentage: coin.MarketCapPercentage,
			},
			Price: models.Price{
				CurrentPrice:             coin.CurrentPrice,
				PriceChangePercentage24h: coin.PriceChangePercentage24h,
				MaxPrice:                 coin.MaxPrice,
				MinPrice:                 coin.MinPrice,
			},
			PriceForDate: models.PriceForDate{
				MaxPriceForDate: coin.MaxPriceForDate,
				MinPriceForDate: coin.MinPriceForDate,
			},
			Volume24h: models.Volume24h{
				High: coin.High,
				Low:  coin.Low,
			},
		}
		coins = append(coins, coin)
	}

	return coins, nil
}
