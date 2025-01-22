package models

import "vigilante/internal/sql/sqlc"

type Preference struct {
	*sqlc.Preference
}
