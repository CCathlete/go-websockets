package postgresrepo

import "log"

func (repo *PGRepo) AuthenticatePassword(
	email, password string,
) (authenticated bool) {

	// Retrieving the user from the database.
	returnedRow, err := repo.queryEngine.GetUserByEmail(
		repo.Context, email,
	)
	if err != nil {
		log.Printf("AuthenticatePassword: %v\n", err)
	}

	return
}
