package postgresrepo

import (
	"database/sql"
	"fmt"
)

var (
	ErrRecordNotFound = sql.ErrNoRows
	ErrTxCommit       = fmt.Errorf("couldn't commit transaction")
	ErrTxRollback     = fmt.Errorf("rollback (UsingTx = true)")
	ErrDiscardedTx    = fmt.Errorf("rollback (UsingTx = false)")
	// User errors.
	ErrUpdatingPassword = fmt.Errorf("error updating password")
	ErrInsertingUser    = fmt.Errorf("error inserting user")
	// Token errors.
	ErrDeleteToken = fmt.Errorf("error deleting token")
)
