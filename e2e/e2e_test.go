package e2e

import (
	"fmt"
	"testing"
)

// tests retrieving the Bitcoin rate.
func TestGetBitcoinRate(t *testing.T) {
	rate, err := GetBitcoinRate()

	if err != nil {
		t.Errorf("Error retrieving Bitcoin rate: %s", err)
	}

	expectedRate := 1000000.0
	if rate != expectedRate {
		t.Errorf("Expected Bitcoin rate %f, but got %f", expectedRate, rate)
	}
}

// tests subscribing to the Bitcoin rate and verifying the subscription.
func TestSubscribeToBitcoinRate(t *testing.T) {

	email := "test@example.com"

	err := SubscribeToBitcoinRate(email)

	if err != nil {
		t.Errorf("Error subscribing to Bitcoin rate: %s", err)
	}

	subscriptions := GetSubscribedEmails()
	subscribed := false
	for _, subscribedEmail := range subscriptions {
		if subscribedEmail == email {
			subscribed = true
			break
		}
	}

	if !subscribed {
		t.Errorf("Email %s is not subscribed to the Bitcoin rate", email)
	}
}

// tests sending an email notification when the Bitcoin rate crosses a threshold.
func TestSendEmailNotification(t *testing.T) {

	threshold = 1200000.0

	SetBitcoinRate(threshold + 1000.0)

	err := CheckBitcoinRateAndSendNotifications()

	if err != nil {
		t.Errorf("Error sending email notification: %s", err)
	}

	notifications := GetSentEmailNotifications()

	expectedNotifications := 1
	if len(notifications) != expectedNotifications {
		t.Errorf("Expected %d email notifications, but got %d", expectedNotifications, len(notifications))
	}

}


// Helper functions for test data and dependencies

var bitcoinRate float64
var subscribedEmails []string
var sentEmailNotifications []EmailNotification
var threshold float64

// Simulates retrieving the Bitcoin rate from an external API
func GetBitcoinRate() (float64, error) {
	bitcoinRate = 1000000.0
	return bitcoinRate, nil
}

// simulates subscribing an email to the Bitcoin rate
func SubscribeToBitcoinRate(email string) error {
	subscribedEmails = append(subscribedEmails, email)
	return nil
}

// retrieves the emails that are subscribed to the Bitcoin rate
func GetSubscribedEmails() []string {
	return subscribedEmails
}

// sets the current Bitcoin rate for testing purposes
func SetBitcoinRate(rate float64) {
	bitcoinRate = rate
}

// simulates checking the Bitcoin rate and sending email notifications
func CheckBitcoinRateAndSendNotifications() error {
	for _, email := range subscribedEmails {
		if bitcoinRate > threshold {
			notification := EmailNotification{
				Recipient: email,
				Content:   fmt.Sprintf("Bitcoin rate crossed the threshold %.2f", threshold),
			}
			sentEmailNotifications = append(sentEmailNotifications, notification)
		}
	}
	return nil
}

// retrieves the email notifications that have been sent
func GetSentEmailNotifications() []EmailNotification {
	return sentEmailNotifications
}

// EmailNotification represents an email notification
type EmailNotification struct {
	Recipient string
	Content   string
}
