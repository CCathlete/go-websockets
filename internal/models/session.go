package models

import "vigilante/internal/sql/sqlc"

// An adapter for sqlc.Session.
type Session struct {
	*sqlc.Session
}
