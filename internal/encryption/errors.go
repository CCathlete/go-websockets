package encryption

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPasswordHasing   = fmt.Errorf("error hashing password")
	ErrPAsswordEncoding = fmt.Errorf("error encoding password")
	ErrPasswordMismatch = bcrypt.ErrMismatchedHashAndPassword
)
