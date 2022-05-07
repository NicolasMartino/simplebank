-- name: CreateUser :one
INSERT INTO users(
    email,
    hashed_password,
    first_name,
    last_name,
    password_change_at
)
VALUES($1, $2, $3, $4, $5)
RETURNING *;

-- name: FindUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: FindUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUserHash :one
update users
set hashed_password = $2, password_change_at = (select now()::timestamptz)
where id =$1
RETURNING *;
