package builder_types

import (
	"github.com/sidan-lab/rum/models"
)

type OutputScriptSource interface {
	isOutputScriptSource()
}

func (ProvidedSimpleScriptSource) isOutputScriptSource() {}
func (ProvidedScriptSource) isOutputScriptSource()       {}

type Output struct {
	Address         string             `json:"address"`
	Amount          []models.Asset     `json:"amount"`
	Datum           Datum              `json:"datum"`
	ReferenceScript OutputScriptSource `json:"referenceScript"`
}
