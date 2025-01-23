package postgres

import (
	"encoding/base64"
	"log"
	"vigilante/internal/encryption"

	"golang.org/x/crypto/bcrypt"
)

func (repo *Repo) AuthenticatePassword(
	email, password string,
) (authenticated bool) {

	// Retrieving the user's info + an indicator is_deleted from the database.
	returnedRow, err := repo.queryEngine.GetUserByEmail(
		repo.Context, email,
	)
	if err != nil {
		err = encryption.ErrPasswordAuthentication(err)
		log.Println(err)
		return
	}

	// If the user is deleted, we can't authenticate credentials.
	is_deleted := returnedRow.IsDeleted.(bool)
	if is_deleted {
		err = encryption.ErrPasswordAuthentication(
			encryption.ErrUserIsDeleted(nil),
		)
		log.Println(err)
		return
	}

	usrPassHashString := returnedRow.PasswordHash

	// The hash was stored as a string using a base64 encoding so we need to decode it.
	usrPassHashBinary, err := base64.StdEncoding.DecodeString(usrPassHashString)

	// Checking if the password is correct.
	err = bcrypt.CompareHashAndPassword()

	return
}
