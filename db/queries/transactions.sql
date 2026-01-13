-- name: CreateTransaction :one
INSERT INTO transactions (
    asset_id, tx_type, quantity, price, fee, occurred_at
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: ListTransactionsByAsset :many
SELECT *
FROM transactions
WHERE asset_id = $1
ORDER BY occurred_at ASC;

-- name: ListAllTransactions :many
SELECT *
FROM transactions
ORDER BY occurred_at ASC;
