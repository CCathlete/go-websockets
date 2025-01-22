package encryption

import "fmt"

var (
	ErrPasswordHasing   = fmt.Errorf("error hashing password")
	ErrPAsswordEncoding = fmt.Errorf("error encoding password")
)
