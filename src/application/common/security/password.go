package security

import "golang.org/x/crypto/bcrypt"

// Hash generates a hashed password from a given plain text password.
//
// Parameters:
// - password: the plain text password to be hashed.
//
// Returns:
// - []byte: the hashed password.
// - error: an error if the hashing process fails.
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Compare compares a hashed password with a plain text password.
//
// Parameters:
// - hashedPassword: the hashed password to compare.
// - password: the plain text password to compare against the hashed password.
// Return type:
// - error: an error if the comparison fails.
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
