-- name: CreateTransaction :one
INSERT INTO transactions (
    asset_id,
    tx_type,
    quantity,
    price,
    fee,
    occurred_at
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;



-- name: ListActiveTransactions :many
SELECT *
FROM transactions
WHERE corrected_by IS NULL
ORDER BY occurred_at DESC;


-- name: ListTransactionsByAsset :many
SELECT *
FROM transactions
WHERE asset_id = $1
  AND corrected_by IS NULL
ORDER BY occurred_at DESC;


-- name: ListAllTransactions :many
SELECT *
FROM transactions
ORDER BY occurred_at DESC;


-- name: GetTransactionByID :one
SELECT *
FROM transactions
WHERE id = $1;


-- name: CorrectTransaction :one
WITH new_tx AS (
    INSERT INTO transactions (
        asset_id,
        tx_type,
        quantity,
        price,
        fee,
        occurred_at
    )
    VALUES (
        $2,
        $3,
        $4,
        $5,
        $6,
        $7
    )
    RETURNING id
)
UPDATE transactions AS t
SET corrected_by = (SELECT new_tx.id FROM new_tx)
WHERE t.id = $1
RETURNING t.*;
