package postgres

import (
	"context"
	"log"
	"time"
	"vigilante/internal/encryption"
	"vigilante/internal/models"
	"vigilante/internal/sql/sqlc"
)

func (repo *Repo) GetUserByID(uid int32) (usr models.User, err error) {
	ctx, cancel := context.WithTimeout(repo.Context, 3*time.Second)
	defer cancel()

	rawUsr, err := repo.queryEngine.GetUserByID(ctx, uid)
	if err != nil {
		err = ErrGettingUser(err)
		log.Println(err)
		return
	}

	usr = *models.NewUser()

	usr.ID = rawUsr.ID
	usr.FirstName = rawUsr.FirstName
	usr.LastName = rawUsr.LastName
	usr.Email = rawUsr.Email
	usr.AccessLevel = rawUsr.AccessLevel
	usr.UserActive = rawUsr.UserActive
	usr.CreatedAt = rawUsr.CreatedAt
	usr.UpdatedAt = rawUsr.UpdatedAt
	usr.DeletedAt = rawUsr.DeletedAt

	return
}

func (repo *Repo) DeleteUser(uid int32) (err error) {
	ctx, cancel := context.WithTimeout(repo.Context, 3*time.Second)
	defer cancel()

	err = repo.queryEngine.DeleteUser(ctx, uid)

	return
}

// ---------------------------------------------------------------------

type UpdateUserParams = sqlc.UpdateUserParams

func (repo *Repo) UpdateUser(params UpdateUserParams,
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

type InsertUserParams struct {
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int32
	UserActive  bool
}

func (repo *Repo) InsertUser(params InsertUserParams,
) (uid int32, err error) {
	ctx, cancel := context.WithTimeout(repo.Context, 3*time.Second)
	defer cancel()

	// Setting up the parameters for the insert query.
	rawParams := sqlc.InsertUserParams{
		FirstName:   params.FirstName,
		LastName:    params.LastName,
		Email:       params.Email,
		AccessLevel: params.AccessLevel,
		UserActive:  params.UserActive,
	}

	hashedPasswordString, err := encryption.HashPassword(params.Password)
	if err != nil {
		log.Printf("InsertUser: Error hashing password, %v\n", err)
		err = ErrInsertingUser(err)
	}

	rawParams.PasswordHash = hashedPasswordString

	uid, err = repo.queryEngine.InsertUser(ctx, rawParams)
	if err != nil {
		err = ErrInsertingUser(err)
	}

	return
}

// ---------------------------------------------------------------------

type UpdatePasswordParams struct {
	Password string
	ID       int32
}

func (repo *Repo) UpdatePassword(params UpdatePasswordParams,
) (err error) {
	ctx, cancel := context.WithTimeout(repo.Context, 3*time.Second)
	defer cancel()

	hashedPasswordString, err := encryption.HashPassword(params.Password)
	if err != nil {
		log.Printf("UpdatePassword: %v\n", err)
		err = ErrUpdatingPassword(err)
	}

	// Setting the parameters as they should be sent to the db + running the query.
	queryParams := sqlc.UpdatePasswordParams{
		PasswordHash: hashedPasswordString,
		ID:           params.ID,
	}
	err = repo.queryEngine.UpdatePassword(ctx, queryParams)
	if err != nil {
		log.Printf("UpdatePassword: %v\n", err)
		err = ErrUpdatingPassword(err)
	}

	err = repo.queryEngine.DeleteTokenByUserID(ctx, params.ID)
	if err != nil {
		log.Printf("UpdatePassword: %v\n", err)
		err = ErrDeleteToken(err)
	}

	return
}

// ---------------------------------------------------------------------
