package builder_types

type Redeemer struct {
	Data    string `json:"data"`
	ExUnits Budget `json:"exUnits"`
}

type Budget struct {
	Mem   uint64 `json:"mem"`
	Steps uint64 `json:"steps"`
}
