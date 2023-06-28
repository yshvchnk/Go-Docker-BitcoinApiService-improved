package service

import (
	// "bitcoin-app/store"
	"errors"
	"fmt"
	// "log"
)

var (
	ErrSubscriptionCheckFailed = errors.New("subscription check failed")
	ErrEmailAlreadySubscribed  = errors.New("email already subscribed")
	ErrFailedToSaveEmail       = errors.New("failed to save email")
)

const emailStoragePath = "../emails.json"

type EmailStorage interface {
	IsEmailSubscribed(email string) (bool, error)
	SaveEmailToFile(email string) error
}

type EmailService struct {
	storage EmailStorage
}

func NewEmailService(storage EmailStorage) *EmailService {
	return &EmailService{
		storage: storage,
	}
}

func (es *EmailService) SubscribeEmail(email string) error {

	subscribed, err := es.storage.IsEmailSubscribed(email)
	if err != nil {
		return fmt.Errorf("failed to check subscription: %s: %s", ErrSubscriptionCheckFailed, err.Error())
	}

	if subscribed {
		return ErrEmailAlreadySubscribed
	}

	err = es.storage.SaveEmailToFile(email)
	if err != nil {
		return fmt.Errorf("failed to save email: %s: %s", ErrFailedToSaveEmail, err.Error())
	}

	return nil
}