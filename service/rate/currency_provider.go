package service

import (
	"os"
)

type ExchangeRateProvider interface {
	GetCurrencyRate() (float64, error)
}

type(
	CoinGeckoAPI struct {
		CoinGeckoURL string
		CoinGeckoIds string
		CoinGeckoVsCurrencies string
	}

	CoinPaprikaAPI struct {
		CoinPaprikaURL string
		CoinPaprikaFrom string
		CoinPaprikaTo string
		CoinPaprikaAmount string
	}

	CryptoCompareAPI struct {
		CryptoCompareURL string
		CryptoCompareFrom string
		CryptoCompareTo string
	}
)

func getCurrencyRate(url string, response CurrencyProviderResponse) (float64, error) {
	return RetrieveAndDecodeFromAPI(url, response)
}

func (api *CoinGeckoAPI) GetCurrencyRate() (float64, error) {
	url := api.CoinGeckoURL + "?ids=" + api.CoinGeckoIds + "&vs_currencies=" + api.CoinGeckoVsCurrencies
	var coinGeckoResp CoinGeckoResponse
	return getCurrencyRate(url, &coinGeckoResp)
}

func (api *CoinPaprikaAPI) GetCurrencyRate() (float64, error) {
	url := api.CoinPaprikaURL + "?base_currency_id=" + api.CoinPaprikaFrom + "&quote_currency_id=" + api.CoinPaprikaTo + "&amount=" + api.CoinPaprikaAmount
	var coinPaprikaResp CoinPaprikaResponse
	return getCurrencyRate(url, &coinPaprikaResp)
}

func (api *CryptoCompareAPI) GetCurrencyRate() (float64, error) {
	url := api.CryptoCompareURL + "?fsym=" + api.CryptoCompareFrom + "&tsyms=" + api.CryptoCompareTo
	var cryptoCompareResp CryptoCompareResponse
	return getCurrencyRate(url, &cryptoCompareResp)
}

type APIProvider struct {
	providers []ExchangeRateProvider
}

func NewCurrencyAPIProvider() *APIProvider {
	providers := []ExchangeRateProvider{
		&CoinGeckoAPI{
			CoinGeckoURL: os.Getenv("COIN_GECKO_API"),
			CoinGeckoIds: os.Getenv("COIN_GECKO_IDS"),
			CoinGeckoVsCurrencies: os.Getenv("COIN_GECKO_VS_CURRENCIES"),
		},

		&CryptoCompareAPI{
			CryptoCompareURL: os.Getenv("CRYPTO_COMPARE_URL"),
			CryptoCompareFrom: os.Getenv("CRYPTO_COMPARE_FROM"),
			CryptoCompareTo: os.Getenv("CRYPTO_COMPARE_TO"),
		},

		&CoinPaprikaAPI{
			CoinPaprikaURL: os.Getenv("COIN_PAPRIKA_URL"),
			CoinPaprikaFrom: os.Getenv("COIN_PAPRIKA_FROM"),
			CoinPaprikaTo: os.Getenv("COIN_PAPRIKA_TO"),
			CoinPaprikaAmount: os.Getenv("COIN_PAPRIKA_AMOUNT"),
		},
	}
	

	return &APIProvider{
		providers: providers,
	}
}