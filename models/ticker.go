package models

type Tickers struct {
	Tickers []Ticker
}
type Ticker struct {
	Base         string `json:"base"`
	Market       Market `json:"market"`
	Target       string `json:"target"`
	TrustScore   string `json:"trust_score"`
	TokenInfoUrl string `json:"token_info_url"`
	CoinID       string `json:"coin_id"`
	TargetCoinID string `json:"target_coin_id"`
}

type Market struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
}
