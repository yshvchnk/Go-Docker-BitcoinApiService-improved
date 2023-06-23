package service_test

import (
	"bitcoin-app/service"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	exitCode := m.Run()

	os.Exit(exitCode)
}

//tests the successful scenario where the API responds with the expected rate
func TestGetBitcoinRate_Success(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
			"bitcoin": {
				"uah": 1000000
			}
		}`))
	}))
	defer mockAPI.Close()

	originalAPI := os.Getenv("COIN_GECKO_API")
	os.Setenv("COIN_GECKO_API", mockAPI.URL)
	defer os.Setenv("COIN_GECKO_API", originalAPI)

	rate, err := service.GetBitcoinRate()

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	expectedRate := 1000000.0
	if rate != expectedRate {
		t.Errorf("Expected rate %f, got: %f", expectedRate, rate)
	}
}

// tests the case where the API returns an error response with a non-OK status code. 
func TestGetBitcoinRate_APIError(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer mockAPI.Close()

	originalAPI := os.Getenv("COIN_GECKO_API")
	os.Setenv("COIN_GECKO_API", mockAPI.URL)
	defer os.Setenv("COIN_GECKO_API", originalAPI)

	rate, err := service.GetBitcoinRate()

	if err == nil {
		t.Error("Expected an error, got nil")
	}
	expectedError := "couldn't get a response from the API, status code: 500"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got: '%s'", expectedError, err.Error())
	}
	if rate != 0 {
		t.Errorf("Expected rate 0, got: %f", rate)
	}
}

//tests the scenario where the API response doesn't contain the expected rate
func TestGetBitcoinRate_RateNotFound(t *testing.T) {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
			"bitcoin": {
				"usd": 40000
			}
		}`))
	}))
	defer mockAPI.Close()

	originalAPI := os.Getenv("COIN_GECKO_API")
	os.Setenv("COIN_GECKO_API", mockAPI.URL)
	defer os.Setenv("COIN_GECKO_API", originalAPI)

	rate, err := service.GetBitcoinRate()

	if err == nil {
		t.Error("Expected an error, got nil")
	}

	expectedError := "failed to retrieve the rate for the bitcoin"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got: '%s'", expectedError, err.Error())
	}

	if rate != 0 {
		t.Errorf("Expected rate 0, got: %f", rate)
	}
}
