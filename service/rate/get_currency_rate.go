package service

import (
	"errors"
)

func NewAPIProvider(providers []ExchangeRateProvider) *APIProvider {
	return &APIProvider{
		providers: providers,
	}
}

func (api *APIProvider) GetCurrencyRate() (float64, error) {

	for _, provider := range api.providers {
		rate, err := provider.GetCurrencyRate()
		if err == nil {
			return rate, nil
		}
	}

	return 0, errors.New("unable to retrieve currency rate from any provider")
}