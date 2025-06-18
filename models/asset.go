package models

type Asset struct {
	Unit     string `json:"unit"`
	Quantity string `json:"quantity"`
}

func NewAsset(unit string, quantity string) Asset {
	return Asset{
		Unit:     unit,
		Quantity: quantity,
	}
}
