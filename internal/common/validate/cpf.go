package validate

import (
	"errors"
	"regexp"
	"strconv"
)

// Cpf validates the format of a Brazilian CPF number.
//
// It takes a CPF number as a string parameter.
// Returns error if it's invalid.
func Cpf(cpf string) error {
	// Remove caracteres não numéricos
	re := regexp.MustCompile(`[^\d]`)
	cpf = re.ReplaceAllString(cpf, "")

	// Verifica se o CPF tem 11 dígitos
	if len(cpf) != 11 {
		return errors.New("invalid CPF: must be 11 digits")
	}

	// Verifica se todos os dígitos são iguais (ex: 111.111.111-11)
	if allDigitsEqual(cpf) {
		return errors.New("invalid CPF: all digits are equal")
	}

	// Valida os dígitos verificadores
	if !validateDigits(cpf) {
		return errors.New("invalid CPF: check digits do not match")
	}

	return nil
}

// allDigitsEqual checks if all digits in the CPF are the same.
func allDigitsEqual(cpf string) bool {
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}
	return true
}

// validateDigits validates the check digits of the CPF.
func validateDigits(cpf string) bool {
	// Calcula o primeiro dígito verificador
	sum := 0
	for i := 0; i < 9; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		sum += num * (10 - i)
	}
	firstCheckDigit := (sum * 10 % 11) % 10

	// Calcula o segundo dígito verificador
	sum = 0
	for i := 0; i < 10; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		sum += num * (11 - i)
	}
	secondCheckDigit := (sum * 10 % 11) % 10

	return firstCheckDigit == int(cpf[9]-'0') && secondCheckDigit == int(cpf[10]-'0')
}
