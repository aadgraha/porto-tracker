-- name: CreateAsset :one
INSERT INTO assets (symbol, name, asset_type)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAssetBySymbol :one
SELECT *
FROM assets
WHERE symbol = $1 AND asset_type = $2;

-- name: ListAssets :many
SELECT *
FROM assets
ORDER BY symbol;
