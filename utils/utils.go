package utils

import (
	"errors"
	"regexp"
)

const DateFormat = "2006-01-02"

// ValidateClassName checks if the class name contains only letters
func ValidateClassName(name string) error {
	if len(name) < 3 {
		return errors.New("class name must be at least 3 characters")
	}
	validName := regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	if !validName(name) {
		return errors.New("class name should contain only letters and spaces")
	}
	return nil
}
