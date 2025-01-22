package postgresrepo

import (
	"context"
	"log"
	"time"
	"vigilante/internal/encryption"
	"vigilante/internal/models"
	"vigilante/internal/sql/sqlc"
)

func (repo *PGRepo) DeleteUser(uid int32) (err error) {
	ctx, cancel := context.WithTimeout(repo.Context, 3*time.Second)
	defer cancel()

	err = repo.queryEngine.DeleteUser(ctx, uid)

	return
}

// ---------------------------------------------------------------------

type UpdateUserParams = sqlc.UpdateUserParams

func (repo *PGRepo) UpdateUser(params UpdateUserParams,
) (usr models.User, err error) {
	ctx, cancel := context.WithTimeout(repo.Context, 3*time.Second)
	defer cancel()

	rawUsr, err := repo.queryEngine.UpdateUser(ctx, params)
	if err != nil {
		return
	}

	usr = models.User{
		User: &rawUsr,
	}

	return
}

// ---------------------------------------------------------------------

type InsertUserParams = sqlc.InsertUserParams

func (repo *PGRepo) InsertUser(params InsertUserParams,
) (uid int32, err error) {
	ctx, cancel := context.WithTimeout(repo.Context, 3*time.Second)
	defer cancel()

	return
}

// ---------------------------------------------------------------------

type UpdatePasswordParams struct {
	Password string
	ID       int32
}

func (repo *PGRepo) UpdatePassword(params UpdatePasswordParams,
) (err error) {
	ctx, cancel := context.WithTimeout(repo.Context, 3*time.Second)
	defer cancel()

	hashedPasswordString, err := encryption.HashPassword(params.Password)
	if err != nil {
		log.Printf("UpdatePassword: %v\n", err)
		err = ErrUpdatingPassword
	}

	// Setting the parameters as they should be sent to the db + running the query.
	queryParams := sqlc.UpdatePasswordParams{
		PasswordHash: hashedPasswordString,
		ID:           params.ID,
	}
	err = repo.queryEngine.UpdatePassword(ctx, queryParams)
	if err != nil {
		log.Printf("UpdatePassword: %v\n", err)
		err = ErrUpdatingPassword
	}

	return
}

// ---------------------------------------------------------------------
