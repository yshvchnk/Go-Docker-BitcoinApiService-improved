package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func GetBitcoinRate() (float64, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var coinGeckoAPI = os.Getenv("COIN_GECKO_API")
	var ids = os.Getenv("IDS")
	var vsCurrencies = os.Getenv("VS_CURRENCIES")

	url := coinGeckoAPI + "?ids=" + ids + "&vs_currencies=" + vsCurrencies

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("couldn't get a response from the API, status code: %d", resp.StatusCode)
	}

	var data map[string]map[string]float64

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	rate, ok := data["bitcoin"]["uah"]
	if !ok {
		return 0, errors.New("failed to retrieve the rate for the bitcoin")
	}

	return rate, nil
}