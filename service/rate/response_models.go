package service

type CoinGeckoResponse struct {
	Currency struct {
		UAH float64 `json:"uah"`
	} `json:"bitcoin"`
}

func (c *CoinGeckoResponse) GetCurrencyRate() float64 {
	return c.Currency.UAH
}

type CryptoCompareResponse struct {
	UAH float64 `json:"UAH"`
}

func (c *CryptoCompareResponse) GetCurrencyRate() float64 {
	return c.UAH
}

type CoinPaprikaResponse struct {
	UAH float64 `json:"price"`
}

func (c *CoinPaprikaResponse) GetCurrencyRate() float64 {
	return c.UAH
}