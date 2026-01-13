-- name: UpsertPriceSnapshot :one
INSERT INTO price_snapshots (asset_id, price, captured_at)
VALUES ($1, $2, $3)
ON CONFLICT (asset_id, captured_at)
DO UPDATE SET price = EXCLUDED.price
RETURNING *;

-- name: GetLatestPrice :one
SELECT *
FROM price_snapshots
WHERE asset_id = $1
ORDER BY captured_at DESC
LIMIT 1;
