package valueobjects

import (
	"strings"
	"taskmanager/internal/common/validate"
)

type Email string

func (e Email) Formater() Email {
	return Email(strings.TrimSpace(string(e)))
}

func (e Email) Validate() error {

	if err := validate.Email(string(e)); err != nil {
		return err
	}
	return nil
}
