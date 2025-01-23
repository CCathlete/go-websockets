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

	ErrPAsswordDecoding EncryptionError = func(err error) error {
		return fmt.Errorf("error decoding password: %w", err)
	}

	ErrPasswordMismatch EncryptionError = func(err error) error {
		return fmt.Errorf("%w: %w",
			bcrypt.ErrMismatchedHashAndPassword,
			err,
		)
	}

	ErrPasswordAuthentication EncryptionError = func(err error) error {
		return fmt.Errorf("AuthenticatePassword: %w", err)
	}

	ErrUserIsDeleted EncryptionError = func(err error) error {
		return fmt.Errorf("user is deleted: %w", err)
	}
)
