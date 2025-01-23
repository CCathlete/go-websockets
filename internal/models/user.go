package models

import "vigilante/internal/sql/sqlc"

// Adapter for sqlc.User.
type User struct {
	*sqlc.User
}

func NewUser() *User {
	return &User{
		User: &sqlc.User{},
	}
}
