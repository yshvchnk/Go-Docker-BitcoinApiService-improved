package service

import (
	"bitcoin-app/file"
	"fmt"
)

func SubscribeEmail(email string) error {

	subscribed, err := file.IsEmailSubscribed(email)
	if err != nil {
		return fmt.Errorf("failed to check subscription: %v", err)
	}

	if subscribed {
		return fmt.Errorf("email already subscribed")
	}

	err = file.SaveEmailToFile(email)
	if err != nil {
		return fmt.Errorf("failed to save email: %v", err)
	}

	return nil
}