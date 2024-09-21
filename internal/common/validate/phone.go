package validate

import (
	"errors"
	"regexp"
	"strings"
)

// Phone validates the format of a Brazilian phone number.
//
// It takes a phone number as a string parameter.
// Returns error if it's invalid.
func Phone(phone string) error {
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")

	if !strings.HasPrefix(phone, "+55") {
		return errors.New("invalid phone number: missing country code +55")
	}

	phone = strings.TrimPrefix(phone, "+55")

	if len(phone) != 11 {
		return errors.New("invalid phone number: must be 11 digits")
	}

	phonePattern := `^[1-9]{2}9[0-9]{8}$`
	if !regexp.MustCompile(phonePattern).MatchString(phone) {
		return errors.New("invalid phone number: does not match Brazilian format")
	}

	return nil
}
