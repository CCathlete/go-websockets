package encryption

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Wraps a specific error in the package functions with a more general error.
type EncryptionError func(error) error

// Wraps a specific error in the package functions with a more general error.
var (
	ErrPasswordHasing EncryptionError = func(err error) error {
		return fmt.Errorf("error hashing password: %w", err)
	}

	ErrPAsswordEncoding EncryptionError = func(err error) error {
		return fmt.Errorf("error encoding password: %w", err)
	}

	ErrPasswordMismatch EncryptionError = func(err error) error {
		return fmt.Errorf("%w: %w",
			bcrypt.ErrMismatchedHashAndPassword,
			err,
		)
	}
)
