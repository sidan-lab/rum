package builder

import (
	"encoding/json"

	"github.com/sidan-lab/rum/common/data"
	types "github.com/sidan-lab/rum/models/builder_types"
)

type WData interface {
	ToCbor() (string, error)
}

type WCbor string

func (c WCbor) ToCbor() (string, error) {
	return string(c), nil
}

type WPlutusData struct {
	Data data.PlutusData
}

// ToCbor TODO: serialize plutus data from rust
func (w WPlutusData) ToCbor() (string, error) {
	jsonData, err := json.Marshal(w.Data)
	if err != nil {
		return "", err
	}
	jsonString := string(jsonData)
	return jsonString, nil
}

type WRedeemer struct {
	Data    WData
	ExUnits types.Budget
}

type WDatum struct {
	Type string
	Data WData
}
