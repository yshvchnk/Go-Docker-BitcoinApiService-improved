package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type CurrencyProviderResponse interface {
	GetCurrencyRate() float64
}

func RetrieveAndDecodeFromAPI(url string, response CurrencyProviderResponse) (float64, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("couldn't get a response from the API, status code: %d", resp.StatusCode)
	}

	bodyBuffer := bytes.NewBuffer(body)
	err = json.NewDecoder(bodyBuffer).Decode(response)
	if err != nil {
		return 0, err
	}

	log.Printf("%T - Response: %s", response, body)
	return response.GetCurrencyRate(), nil
}