-- name: InsertUser :one
   insert into users (
          first_name,
          last_name,
          email,
          password_hash,
          access_level,
          user_active
          )
   values ($1, $2, $3, $4, $5, $6)
returning id
;