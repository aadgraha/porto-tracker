-- Assets (BTC, ETH, AAPL)
CREATE TABLE assets (
    id           BIGSERIAL PRIMARY KEY,
    symbol       TEXT NOT NULL,
    name         TEXT,
    asset_type   TEXT NOT NULL CHECK (asset_type IN ('crypto', 'stock')),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE(symbol, asset_type)
);

-- Transactions (append-only)
CREATE TABLE transactions (
    id            BIGSERIAL PRIMARY KEY,

    asset_id      BIGINT NOT NULL
        REFERENCES assets(id),

    tx_type       TEXT NOT NULL CHECK (
        tx_type IN ('buy', 'sell', 'transfer_in', 'transfer_out')
    ),

    quantity      NUMERIC(36, 18) NOT NULL,

    price         NUMERIC(36, 18), -- nullable for transfers

    fee           NUMERIC(36, 18) NOT NULL DEFAULT 0,

    occurred_at   TIMESTAMPTZ NOT NULL,

    -- correction / audit fields
    corrected_by  BIGINT
        REFERENCES transactions(id),

    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),

    -- prevent self-correction
    CONSTRAINT transactions_no_self_correction
        CHECK (corrected_by IS NULL OR corrected_by <> id)
);


CREATE INDEX idx_transactions_asset_time
ON transactions(asset_id, occurred_at);

-- Price snapshots (historical)
CREATE TABLE price_snapshots (
    id            BIGSERIAL PRIMARY KEY,
    asset_id      BIGINT NOT NULL REFERENCES assets(id),
    price         NUMERIC(36, 18) NOT NULL,
    captured_at   TIMESTAMPTZ NOT NULL,
    UNIQUE(asset_id, captured_at)
);

CREATE INDEX idx_price_snapshots_asset_time
ON price_snapshots(asset_id, captured_at);
