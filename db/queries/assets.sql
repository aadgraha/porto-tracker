-- name: CreateAsset :one
INSERT INTO assets (symbol, name, asset_type)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAssetByID :one
SELECT *
FROM assets
WHERE id = $1;

-- name: ListAssets :many
SELECT *
FROM assets
ORDER BY id;

-- name: UpdateAsset :one
UPDATE assets
SET
  symbol = $2,
  name = $3,
  asset_type = $4
WHERE id = $1
RETURNING *;

-- name: DeleteAsset :exec
DELETE FROM assets
WHERE id = $1;
