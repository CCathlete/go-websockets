-- name: InsertRememberMeToken :exec
   insert into remember_tokens (user_id, remember_token)
   values ($1, $2)
;

-- name: DeleteToken :exec
   delete from remember_tokens
    where remember_token = $1
;

-- name: CheckForToken :one
   select exists (
             select 1
               from remember_tokens
              where user_id = $1
                and remember_token = $2
          ) as exists
;

-- name: GetTokenByID :one
   select
     from remember_tokens
    where user_id = $1
      and remember_token = $2
;
