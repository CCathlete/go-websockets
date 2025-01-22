package repository

import "vigilante/internal/sql/sqlc"

// An adapter for sqlc.Querier which is an interface that defines all the methods that are available for the database.
type Repository interface {
	*sqlc.Querier
}
