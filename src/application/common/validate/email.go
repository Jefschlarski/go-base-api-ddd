package validate

import (
	"errors"
	"net"
	"regexp"
	"strings"
)

// Email validates the format of an email address.
//
// It takes an email address as a string parameter.
// Returns error if it's invalid.
func Email(email string) error {

	if !strings.Contains(email, "@") {
		return errors.New("invalid email address, missing '@'")
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return errors.New("invalid email address, too many '@'")
	}

	local := parts[0]
	if err := validateLocalPart(local); err != nil {
		return err
	}

	domain := parts[1]
	if err := validateDomainPart(domain); err != nil {
		return err
	}

	return nil
}

// validateLocalPart validates the format of a local part.
func validateLocalPart(local string) error {

	allowedChars := "[a-zA-Z0-9!#$%&'*+/=?^_{|}`~-]+"

	if !regexp.MustCompile(allowedChars).MatchString(local) {
		return errors.New("local part of email is invalid: characters not allowed")
	}
	return nil
}

// validateDomainPart validates the format of a domain part.
func validateDomainPart(domain string) error {

	if ip := net.ParseIP(domain); ip != nil {
		return nil
	}

	if !strings.Contains(domain, ".") {
		return errors.New("invalid domain: missing dot (.)")
	}

	labels := strings.Split(domain, ".")
	for _, label := range labels {
		if err := validateLabel(label); err != nil {
			return err
		}
	}

	return nil
}

// validateLabel validates the format of a domain label.
func validateLabel(label string) error {

	allowedChars := "[a-zA-Z0-9-]+"

	if !regexp.MustCompile(allowedChars).MatchString(label) {
		return errors.New("domain label is invalid: characters not allowed")
	}

	return nil
}
