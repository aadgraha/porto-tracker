package asset

type CreateAssetRequest struct {
	Symbol    string  `json:"symbol"`
	Name      *string `json:"name,omitempty"`
	AssetType string  `json:"asset_type"`
}

type UpdateAssetRequest struct {
	Symbol    string  `json:"symbol"`
	Name      *string `json:"name,omitempty"`
	AssetType string  `json:"asset_type"`
}
