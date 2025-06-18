package builder_types

import "github.com/sidan-lab/rum/models"

type UtxoInput struct {
	OutputIndex uint32 `json:"outputIndex"`
	TxHash      string `json:"txHash"`
}

type UtxoOutput struct {
	Address    string         `json:"address"`
	Amount     []models.Asset `json:"amount"`
	DataHash   *string        `json:"dataHash"`
	PlutusData *string        `json:"plutusData"`
	ScriptRef  *string        `json:"scriptRef"`
	ScriptHash *string        `json:"scriptHash"`
}

type UTxO struct {
	Input  UtxoInput  `json:"input"`
	Output UtxoOutput `json:"output"`
}
