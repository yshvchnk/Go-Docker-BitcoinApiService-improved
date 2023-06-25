package integration

import (
	"bitcoin-app/handler"
	"bitcoin-app/service"
	"bitcoin-app/file"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"log"
)

func TestHandler(t *testing.T) {

	testEmail := "test2@example.com"

	reqSubscribe, err := http.NewRequest(http.MethodPost, "/api/subscribe", strings.NewReader("email="+testEmail))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	recSubscribe := httptest.NewRecorder()

	handler.HandleSubscribe(recSubscribe, reqSubscribe)

	if recSubscribe.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, recSubscribe.Code)
	}

	if err := service.SubscribeEmail(testEmail); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	es, err := file.NewEmailStorage("emails.json")
	if err != nil {
			log.Fatal(err)
	}

	subscribed, err := es.IsEmailSubscribed(testEmail);
	if err != nil {
		t.Errorf("failed to check subscription: %v", err)
	}

	if !subscribed {
		t.Error("email was not added")
	}

}