package models

type Coin struct {
	Rank                 int     `json:"market_cap_rank"`
	Id                   string  `json:"id"`
	Image                string  `json:"image"`
	Symbol               string  `json:"symbol"`
	MarketCapitalization float64 `json:"market_cap"`
	TotalVolume          float64 `json:"total_volume"`
	Supply               Supply
	Change               Change
	Price                Price
	PriceForDate         PriceForDate
	Volume24h            Volume24h
}

type Supply struct {
	CirculatingSupply float64 `json:"circulating_supply"`
	TotalSupply       float64 `json:"total_supply"`
	RemainingSupply   float64 `json:"remaining_supply"`
}

type Change struct {
	Price               float64 `json:"price_change_24h"`
	MarketCap           float64 `json:"market_cap_change_24h"`
	MarketCapPercentage float64 `json:"market_cap_change_percentage_24h"`
}

type Price struct {
	CurrentPrice             float64 `json:"current_price"`
	PriceChangePercentage24h float64 `json:"price_change_percentage_24h"`
	MaxPrice                 float64 `json:"ath"`
	MinPrice                 float64 `json:"atl"`
}

type PriceForDate struct {
	MaxPriceForDate string `json:"ath_date"`
	MinPriceForDate string `json:"atl_date"`
}

type Volume24h struct {
	High float64 `json:"high_24h"`
	Low  float64 `json:"low_24h"`
}

type CoinAllInfo struct {
	Rank                     int     `json:"market_cap_rank"`
	Id                       string  `json:"id"`
	Image                    string  `json:"image"`
	Symbol                   string  `json:"symbol"`
	MarketCapitalization     float64 `json:"market_cap"`
	TotalVolume              float64 `json:"total_volume"`
	CirculatingSupply        float64 `json:"circulating_supply"`
	TotalSupply              float64 `json:"total_supply"`
	RemainingSupply          float64 `json:"remaining_supply"`
	Price                    float64 `json:"price_change_24h"`
	MarketCap                float64 `json:"market_cap_change_24h"`
	MarketCapPercentage      float64 `json:"market_cap_change_percentage_24h"`
	CurrentPrice             float64 `json:"current_price"`
	PriceChangePercentage24h float64 `json:"price_change_percentage_24h"`
	MaxPrice                 float64 `json:"ath"`
	MinPrice                 float64 `json:"atl"`
	MaxPriceForDate          string  `json:"ath_date"`
	MinPriceForDate          string  `json:"atl_date"`
	High                     float64 `json:"high_24h"`
	Low                      float64 `json:"low_24h"`
}
