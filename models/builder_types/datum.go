package builder_types

type Datum interface {
	isDatum()
}

type InlineDatum struct {
	Inner string `json:"inline"`
}

func (InlineDatum) isDatum() {}


type HashDatum struct {
	Inner string `json:"hash"`
}

func (HashDatum) isDatum() {}


type EmbeddedDatum struct {
	Inner string `json:"embedded"`
}

func (EmbeddedDatum) isDatum() {}
