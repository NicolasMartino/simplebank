-- name: CreateEntry :one
INSERT INTO entries(
    account_id,
    amount
    )
VALUES($1, $2)
RETURNING *;


-- name: FindEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: FindEntriesWithPagination :many
SELECT * FROM entries
LIMIT $1
OFFSET $2;

-- name: UpdateEntry :one
Update entries 
SET amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;