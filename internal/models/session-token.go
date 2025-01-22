package models

import "vigilante/internal/sql/sqlc"

// An adapter for sqlc.RememberToken.
type SessionToken struct {
	*sqlc.RememberToken
}
