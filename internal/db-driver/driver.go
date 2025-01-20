package driver

import (
	"database/sql"
	"time"
)

// Holds the db connection object.
type DB struct {
	*sql.DB
}

var dbConn = &DB{}

const (
	maxOpenDbConn = 25
	maxIdleDbConn = 25
	maxDbLifetime = 5 * time.Minute
)
