// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: tokens.sql

package sqlc

import (
	"context"
)

const checkForToken = `-- name: CheckForToken :one
   select exists (
             select 1
               from remember_tokens
              where user_id = $1
                and remember_token = $2
          ) as exists
`

type CheckForTokenParams struct {
	UserID        int32  `json:"user_id"`
	RememberToken string `json:"remember_token"`
}

func (q *Queries) CheckForToken(ctx context.Context, arg CheckForTokenParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkForToken, arg.UserID, arg.RememberToken)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const deleteToken = `-- name: DeleteToken :exec
   delete from remember_tokens
    where remember_token = $1
`

func (q *Queries) DeleteToken(ctx context.Context, rememberToken string) error {
	_, err := q.db.ExecContext(ctx, deleteToken, rememberToken)
	return err
}

const getTokenByID = `-- name: GetTokenByID :one
   select
     from remember_tokens
    where user_id = $1
      and remember_token = $2
`

type GetTokenByIDParams struct {
	UserID        int32  `json:"user_id"`
	RememberToken string `json:"remember_token"`
}

type GetTokenByIDRow struct {
}

func (q *Queries) GetTokenByID(ctx context.Context, arg GetTokenByIDParams) (GetTokenByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getTokenByID, arg.UserID, arg.RememberToken)
	var i GetTokenByIDRow
	err := row.Scan()
	return i, err
}

const insertRememberMeToken = `-- name: InsertRememberMeToken :exec
   insert into remember_tokens (user_id, remember_token)
   values ($1, $2)
`

type InsertRememberMeTokenParams struct {
	UserID        int32  `json:"user_id"`
	RememberToken string `json:"remember_token"`
}

func (q *Queries) InsertRememberMeToken(ctx context.Context, arg InsertRememberMeTokenParams) error {
	_, err := q.db.ExecContext(ctx, insertRememberMeToken, arg.UserID, arg.RememberToken)
	return err
}
