package postgresrepo

import (
	"context"
	"database/sql"
	"log"
	"vigilante/internal/sql/sqlc"
)

// An adapter for sqlc.Queries that implements the actual postgres DB methods.
type PGRepo struct {
	queryEngine *sqlc.Queries // Meant only for runing queries. DB changes should be done through repo methods operating on the db field.
	db          *sql.DB
	Context     context.Context
}

// A factory for creating a new PGRepo object.
func New(db *sql.DB) (repo *PGRepo) {

	repo = &PGRepo{
		queryEngine: sqlc.New(db),
		db:          db,
		Context:     context.Background(),
	}

	return
}

func (repo *PGRepo) Close() (err error) {
	err = repo.db.Close()
	return
}

// Registers a transaction object as the inner db connection.
func (repo *PGRepo) WithTx() (err error) {

	// Creates a new transaction object.
	tx, err := repo.db.BeginTx(repo.Context, nil)
	if err != nil {
		return
	}

	// NOTE: The Queries object uses the transaction but the repo maintains the original db object.

	// Turns the underlying db connection into a transaction connection.
	repo.queryEngine = repo.queryEngine.WithTx(tx)

	return
}

func (repo *PGRepo) CommitOrRollback(
	tx *sql.Tx,
	TxExecErr error,
) (err error) {

	// Registering the original db connection for queries. We already ran the queries relevant to the current tx object, just need to commit or rollback.
	repo.queryEngine = sqlc.New(repo.db)

	if tx == nil {
		err = ErrDiscardedTx(nil)
		return
	}

	switch TxExecErr {
	case nil:
		err = tx.Commit()
		if err != nil {
			log.Printf("Error committing tx: %v\n", err)
			err = ErrTxCommit(err)
		}
		return

	default:
		log.Printf("Tx execution error: %v\nRolling back...", TxExecErr)
		err = tx.Rollback()
		if err != nil {
			log.Printf("Error rolling back tx: %v\n", err)
			err = ErrTxRollback(err)
		}
	}

	return
}
