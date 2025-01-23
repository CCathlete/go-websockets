package encryption

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Wraps the error returned by the database with a custom error.
type RepoError func(error) error

// Wraps the error returned by the database with a custom error.
var (
	ErrPasswordHasing RepoError = func(err error) error {
		return fmt.Errorf("error hashing password: %w", err)
	}

	ErrPAsswordEncoding RepoError = func(err error) error {
		return fmt.Errorf("error encoding password: %w", err)
	}

	ErrPasswordMismatch RepoError = func(err error) error {
		return fmt.Errorf("%w: %w",
			bcrypt.ErrMismatchedHashAndPassword,
			err,
		)
	}
)
