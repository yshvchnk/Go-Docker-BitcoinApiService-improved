package store

import (
	"encoding/json"
	"os"
	"github.com/pkg/errors"
)

type EmailStorage struct {
	StoragePath string
}

func NewEmailStorage(storagePath string) (*EmailStorage, error) {
	es := &EmailStorage{
		StoragePath: storagePath,
	}

	err := es.loadEnv()
	if err != nil {
		return nil, err
	}

	return es, nil
}

func (es *EmailStorage) loadEnv() error {
	es.StoragePath = os.Getenv("STORAGE")
	return nil
}

func (es *EmailStorage) GetEmailsFromFile() ([]string, error) {
	data, err := os.ReadFile(es.StoragePath)
	if err != nil {
		return nil, err
	}

	var emails []string
	err = json.Unmarshal(data, &emails)
	if err != nil {
		return nil, err
	}

	return emails, nil
}

func (es *EmailStorage) IsEmailSubscribed(email string) (bool, error) {
	emails, err := es.GetEmailsFromFile()
	if err != nil {
		return false, errors.Wrap(err, "Failed to load email addresses")
	}

	for _, e := range emails {
		if e == email {
			return true, nil
		}
	}

	return false, nil
}

func (es *EmailStorage) SaveEmailToFile(email string) error {
	emails, err := es.GetEmailsFromFile()
	if err != nil {
		emails = []string{}
	}

	emails = append(emails, email)
	data, err := json.Marshal(emails)
	if err != nil {
		return err
	}

	err = os.WriteFile(es.StoragePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}