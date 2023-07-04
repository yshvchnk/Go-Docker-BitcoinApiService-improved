package handler

import (
	"encoding/json"
	"net/http"
)

type (
	CurrencyRateProvider interface {
		GetCurrencyRate() (float64, error)
	}

	CurrencyRateHandler struct {
		CurrencyAPI CurrencyRateProvider
	}
)

func NewCurrencyRateHandler(currencyAPI CurrencyRateProvider) *CurrencyRateHandler {
	return &CurrencyRateHandler{
		CurrencyAPI: currencyAPI,
	}
}

func (h *CurrencyRateHandler) HandleCurrencyRate(w http.ResponseWriter, r *http.Request) {
	rate, err := h.CurrencyAPI.GetCurrencyRate()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(rate)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}


