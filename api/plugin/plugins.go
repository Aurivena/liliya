package plugin

import (
	"liliya/api/plugin/coinGecko"
	"liliya/models"
)

type CoinGeckoCoins interface {
	GetInfoCoins(currency, page string) ([]models.Coin, error)
}

type CoinGeckoMarkets interface {
	GetAllTMarketsByID(id string) ([]string, error)
}

type CoinGeckoCategory interface {
	GetCategoriesByID(id string) ([]string, error)
	GetCategories() ([]models.Category, error)
}

type Plugin struct {
	CoinGeckoCoins
	CoinGeckoMarkets
	CoinGeckoCategory
}

func NewPlugin(cfg *models.Config, env *models.Env) *Plugin {
	return &Plugin{
		CoinGeckoCoins:    coinGecko.NewCoinGeckoCoinPlugin(cfg.Server.CoinGecko),
		CoinGeckoCategory: coinGecko.NewCoinGeckoCategoryPlugin(cfg.Server.CoinGecko),
		CoinGeckoMarkets:  coinGecko.NewCoinGeckoMarketPlugin(cfg.Server.CoinGecko),
	}
}
