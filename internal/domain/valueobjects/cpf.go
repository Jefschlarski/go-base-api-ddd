package valueobjects

import (
	"strings"
	"taskmanager/internal/common/validate"
)

type Cpf string

func (p Cpf) Formater() Cpf {
	return Cpf(strings.TrimSpace(string(p)))
}

func (p Cpf) Validate() error {
	if err := validate.Cpf(string(p)); err != nil {
		return err
	}
	return nil
}
