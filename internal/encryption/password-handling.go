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

	// -----------------------------------------------------------------
	// NOTE: Converting ascii characters to bytes is ok but the other way around could result in 8-bit numbers that doesn't have an ascii character. That's why base64 (6-bit numbers) is used, it maps perfectly to ascii characters.
	// -----------------------------------------------------------------

	// Converting the hashed password to base64.
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
