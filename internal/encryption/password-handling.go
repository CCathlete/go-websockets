package encryption

import (
	"bytes"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string,
) (hashedPassString string, err error) {

	hashedPasswordBinary, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	// Converting the hashed password to base64 because it has a wider range of ascii characters we can use to represent a digit in the string form.
	var conversionBuffer *bytes.Buffer
	b64Encoder := base64.NewEncoder(
		base64.StdEncoding,
		conversionBuffer,
	)
	if err != nil {
		log.Printf("HashPassword: %v\n", err)
		err = ErrPAsswordEncoding(err)
		return
	}

	_, err = b64Encoder.Write(
		hashedPasswordBinary,
	)
	if err != nil {
		log.Printf("HasPassword: %v\n", err)
		err = ErrPasswordHasing(err)
		return
	}

	err = b64Encoder.Close()
	if err != nil {
		log.Printf("HashPassword: %v\n", err)
		err = ErrPasswordHasing(err)
		return
	}

	hashedPassString = conversionBuffer.String()

	return
}
