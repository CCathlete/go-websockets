package models

import "vigilante/internal/sql/sqlc"

// Adapter for sqlc.User.
type User struct {
	*sqlc.User
}
