package e2e

import (
	"bitcoin-app/handler"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"strings"
	"testing"
)

func TestGetBitcoinRate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler.HandleRate))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Errorf("Error retrieving Bitcoin rate: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %s", err)
	}

	validPattern := regexp.MustCompile(`\b[1-9]\d*\b`)
	if !validPattern.Match(body) {
		t.Errorf("Expected response body to contain a digit greater than 0, but it does not")
	}
}

func TestSubscribeToBitcoinRate(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(handler.HandleSubscribe))
	defer server.Close()

	formData := url.Values{}
	formData.Set("email", "test@example.com")
	requestBody := strings.NewReader(formData.Encode())

	resp, err := http.Post(server.URL, "application/x-www-form-urlencoded", requestBody)
	if err != nil {
		t.Fatalf("Error subscribing to Bitcoin rate: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

}

func TestSendEmailNotification(t *testing.T) {

	emailHandler, err := handler.NewEmailHandler("emails.json")
	if err != nil {
		t.Fatalf("Failed to create EmailHandler: %s", err)
	}

	server := httptest.NewServer(http.HandlerFunc(emailHandler.HandleSendEmails))
	defer server.Close()

	resp, err := http.Post(server.URL, "application/json", nil)
	if err != nil {
		t.Fatalf("Error sending email notification: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

}
