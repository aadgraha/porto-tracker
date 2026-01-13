-- name: GetNetQuantityByAsset :many
SELECT
    asset_id,
    SUM(
        CASE
            WHEN tx_type IN ('buy', 'transfer_in') THEN quantity
            WHEN tx_type IN ('sell', 'transfer_out') THEN -quantity
        END
    ) AS net_quantity
FROM transactions
GROUP BY asset_id;
