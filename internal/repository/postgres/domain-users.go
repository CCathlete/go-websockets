package postgresrepo

import (
	"context"
	"time"
)

func (repo *PGRepo) DeleteUser(uid int32) (err error) {
	ctx, cancel := context.WithTimeout(repo.Context, 3*time.Second)
	defer cancel()

	err = repo.queryEngine.DeleteUser(ctx, uid)

	return
}
