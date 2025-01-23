package postgresrepo

import (
	"database/sql"
	"fmt"
)

type (
	RepoError func(error) error
)

var ErrRecordNotFound RepoError = func(err error) error {
	return fmt.Errorf("%w: %w", err, sql.ErrNoRows)
}

var ErrTxCommit RepoError = func(err error) error {
	return fmt.Errorf("%w: couldn't commit transaction", err)
}

var ErrTxRollback RepoError = func(err error) error {
	return fmt.Errorf("%w: rollback (UsingTx = true)", err)
}

var ErrDiscardedTx RepoError = func(err error) error {
	return fmt.Errorf("%w: rollback (UsingTx = false)", err)
}

// User errors.
var ErrUpdatingPassword RepoError = func(err error) error {
	return fmt.Errorf("%w: error updating password", err)
}

var ErrInsertingUser RepoError = func(err error) error {
	return fmt.Errorf("%w: error inserting user", err)
}

// Token errors.
var ErrDeleteToken RepoError = func(err error) error {
	return fmt.Errorf("%w: error deleting token", err)
}
