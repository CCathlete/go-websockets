package postgresrepo

import "fmt"

var (
	ErrTxCommit         = fmt.Errorf("couldn't commit transaction")
	ErrTxRollback       = fmt.Errorf("rollback (UsingTx = true)")
	ErrDiscardedTx      = fmt.Errorf("rollback (UsingTx = false)")
	ErrUpdatingPassword = fmt.Errorf("error updating password")
)
