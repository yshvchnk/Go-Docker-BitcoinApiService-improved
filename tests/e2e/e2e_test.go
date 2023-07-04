package test

import (
	"bitcoin-app/handler"
	"bitcoin-app/store"
	rate "bitcoin-app/service/rate"
	send "bitcoin-app/service/send"
	subscribe "bitcoin-app/service/subscribe"
	test "bitcoin-app/tests"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"strings"
	"testing"
	"log"
)

const storagePath = "emails.json"

func TestGetBitcoinRateWillReturnSuccess(t *testing.T) {
	bitcoinAPI := rate.NewCurrencyAPIProvider()

	bitcoinRateHandler := handler.NewCurrencyRateHandler(bitcoinAPI)

	server := httptest.NewServer(http.HandlerFunc(bitcoinRateHandler.HandleCurrencyRate))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Errorf("Error retrieving Bitcoin rate: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %s", err)
	}

	validPattern := regexp.MustCompile(`\b\d+(\.\d+)?\b`)
	if !validPattern.Match(body) {
		t.Errorf("Expected response body to contain a digit greater than 0, but it does not")
	}
}

func TestSubscribeToBitcoinRateWillReturnSuccess(t *testing.T) {

	err := test.ClearFileContents(storagePath)
	if err != nil {
		t.Fatalf("Error cleaning up emails.json: %s", err)
	}

	emailServiceSubscribe := subscribe.NewEmailServiceSubscribe()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.HandleSubscribeEmails(w, r, emailServiceSubscribe)
	}))
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

func TestSendEmailNotificationWillReturnSuccess(t *testing.T) {
	emailSender := &send.EmailSenderPath{
		StoragePath: storagePath,
	}

	emailStorage, err := store.NewEmailStorage(storagePath)
	if err != nil {
		log.Fatal(err)
	}

	emailHandler, err := handler.NewEmailSendHandler(*emailStorage, emailSender)
	if err != nil {
		log.Fatal(err)
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