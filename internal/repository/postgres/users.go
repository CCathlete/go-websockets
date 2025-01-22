package postgresrepo

import (
	"context"
	"time"
)

func (repo *PGRepo) DeleteUser(uid int) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return
}
