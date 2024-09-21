package valueobjects

import (
	"strings"
	"taskmanager/internal/common/validate"
)

type Phone string

func (p Phone) Formater() Phone {
	return Phone(strings.TrimSpace(string(p)))
}

func (p Phone) Validate() error {
	if err := validate.Phone(string(p)); err != nil {
		return err
	}
	return nil
}
