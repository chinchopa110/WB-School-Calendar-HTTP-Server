package Validation

import (
	"errors"
	"regexp"
	"time"
)

func IsValidEvent(userIdStr string, key string, date string, description string) error {
	if len(userIdStr) == 0 {
		return errors.New("user_id is required")
	}
	if len(key) == 0 {
		return errors.New("key is required")
	}
	if len(date) == 0 {
		return errors.New("date is required")
	}
	if len(description) == 0 {
		return errors.New("description is required")
	}

	if err := IsValidDate(date); err != nil {
		return err
	}

	return nil
}

func IsValidUser(key string) error {
	if len(key) == 0 {
		return errors.New("key is required")
	}
	return nil
}

func IsValidDate(date string) error {
	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !dateRegex.MatchString(date) {
		return errors.New("invalid date format (yyyy-mm-dd)")
	}

	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return err
	}
	return nil
}
