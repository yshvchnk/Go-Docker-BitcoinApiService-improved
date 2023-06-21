package files

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

var storage string

func init() {
	err := LoadEnv()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return errors.Wrap(err, "Error loading .env file")
	}

	storage = os.Getenv("STORAGE")
	return nil
}

func IsEmailSubscribed(email string) (bool, error) {
	emails, err := GetEmailsFromFile()
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

func SaveEmailToFile(email string) error {
	emails, err := GetEmailsFromFile()
	if err != nil {
		emails = []string{}
	}

	emails = append(emails, email)
	data, err := json.Marshal(emails)
	if err != nil {
		return err
	}

	err = os.WriteFile(storage, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetEmailsFromFile() ([]string, error) {

	data, err := os.ReadFile(storage)
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