package postgresrepo

import (
	"database/sql"
	"fmt"
)

// Wraps the error returned by the database with a custom error.
type RepoError func(error) error

// Wraps the error returned by the database with a custom error.
var (
	ErrRecordNotFound RepoError = func(err error) error {
		return fmt.Errorf("%w: %w", err, sql.ErrNoRows)
	}

	ErrTxCommit RepoError = func(err error) error {
		return fmt.Errorf("%w: couldn't commit transaction", err)
	}

	ErrTxRollback RepoError = func(err error) error {
		return fmt.Errorf("%w: rollback (UsingTx = true)", err)
	}

	ErrDiscardedTx RepoError = func(err error) error {
		return fmt.Errorf("%w: rollback (UsingTx = false)", err)
	}

	// User errors.

	ErrUpdatingPassword RepoError = func(err error) error {
		return fmt.Errorf("%w: error updating password", err)
	}

	ErrInsertingUser RepoError = func(err error) error {
		return fmt.Errorf("%w: error inserting user", err)
	}

	// Token errors.

	ErrDeleteToken RepoError = func(err error) error {
		return fmt.Errorf("%w: error deleting token", err)
	}
)
