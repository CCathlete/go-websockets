-- name: GetUserByEmail :one
   select *,
          deleted_at is null as is_deleted
     from users
    where email = $1
;

-- name: GetUserByID :one
   select *,
          deleted_at is null as is_deleted
     from users
    where id = $1
;

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

-- name: UpdateUser :one
   update users
      set first_name = $1,
          last_name = $2,
          user_active = $3,
          email = $4,
          access_level = $5,
          updated_at = now()
    where id = $6
returning *
;

-- name: DeleteUser :exec
   update users
      set user_active = false,
          deleted_at = now()
    where id = $1
;

-- name: UpdatePassword :exec
   update users
      set password_hash = $1
    where id = $2
;

-- //TODO: Make sure to delete the remember_token for this user after resetting the password.