package builder_types

type ValidityRange struct {
	InvalidBefore    *uint64 `json:"invalidBefore"`
	InvalidHereafter *uint64 `json:"invalidHereafter"`
}
