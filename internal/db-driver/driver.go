package driver

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Holds a db connection.
type DB struct {
	*sql.DB
}

const (
	maxOpenDbConns = 25
	maxIdleDbConns = 25
	maxDbLifetime  = 5 * time.Minute
)

// A factory for a db connection.
func ConnectToPostgres(connectionString string,
) (dbConn *DB, err error) {

	// --------------- Creating a new db connection --------------------

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		err = fmt.Errorf("couldn't create a db connection: %w", err)
		return
	}

	db.SetMaxOpenConns(maxOpenDbConns)
	db.SetMaxIdleConns(maxIdleDbConns)
	db.SetConnMaxLifetime(maxDbLifetime)
	dbConn = &DB{
		DB: db,
	}

	// -------------------- Pinging the db -----------------------------

	err = dbConn.Ping()
	if err != nil {
		err = fmt.Errorf("created a db connection but couldn't ping: %w", err)
	} else {
		log.Println("*** Pinged DB successfully ***")
	}

	return
}
