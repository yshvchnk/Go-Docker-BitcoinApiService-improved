package integration

import (
	"bitcoin-app/handler"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

func TestHandler(t *testing.T) {

	reqRate, err := http.NewRequest(http.MethodGet, "/api/rate", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	recRate := httptest.NewRecorder()

	handler.HandleRate(recRate, reqRate)

	if recRate.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, recRate.Code)
	}

	var rateValue float64
	err = json.Unmarshal(recRate.Body.Bytes(), &rateValue)
	if err != nil {
			t.Errorf("failed to decode response body: %v", err)
	}

	if rateValue <= 800000 || rateValue > 1500000 {
    t.Errorf("unexpected rate value: %f", rateValue)
}

	reqSubscribe, err := http.NewRequest(http.MethodPost, "/api/subscribe", strings.NewReader("email=test2@example.com"))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	recSubscribe := httptest.NewRecorder()

	handler.HandleSubscribe(recSubscribe, reqSubscribe)

	if recSubscribe.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, recSubscribe.Code)
	}
}