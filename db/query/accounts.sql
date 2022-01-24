-- name: CreateAccount :one
INSERT INTO accounts(
    owner, 
    balance, 
    currency
)
VALUES($1, $2, $3)
RETURNING *;

-- name: FindAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: FindAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1 
FOR NO KEY UPDATE; -- avoids deadlocks

-- name: FindAllAccountWithPagination :many
SELECT * FROM accounts
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
Update accounts 
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: UpdateAccountAddToBalance :one
Update accounts 
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;