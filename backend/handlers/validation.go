package handlers

import (
	"fmt"
	"regexp"
	"strings"
)

// validateEmail validates email format using regex
func validateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email is required")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}

	return nil
}

// validatePassword validates password strength
func validatePassword(password string) error {
	if password == "" {
		return fmt.Errorf("password is required")
	}

	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	if len(password) > 128 {
		return fmt.Errorf("password too long (max 128 characters)")
	}

	return nil
}

// validateName validates user name
func validateName(name string) error {
	if len(name) > 100 {
		return fmt.Errorf("name too long (max 100 characters)")
	}

	// Allow empty name for optional updates
	if name != "" {
		name = strings.TrimSpace(name)
		if len(name) < 2 {
			return fmt.Errorf("name must be at least 2 characters long")
		}
	}

	return nil
}
