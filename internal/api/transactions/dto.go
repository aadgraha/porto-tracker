package transaction

import "time"

type CreateTransactionRequest struct {
	AssetID    int64     `json:"asset_id"`
	TxType     string    `json:"tx_type"`
	Quantity   string    `json:"quantity"` // NUMERIC → string
	Price      *string   `json:"price,omitempty"`
	Fee        *string   `json:"fee,omitempty"`
	OccurredAt time.Time `json:"occurred_at"`
}

type CorrectTransactionRequest struct {
	AssetID    int64     `json:"asset_id"`
	TxType     string    `json:"tx_type"`
	Quantity   string    `json:"quantity"`
	Price      *string   `json:"price,omitempty"`
	Fee        *string   `json:"fee,omitempty"`
	OccurredAt time.Time `json:"occurred_at"`
}
