package models

type Input struct {
	OutputIndex int    `json:"output_index" binding:"required"`
	TxHash      string `json:"tx_hash" binding:"required"`
}

type Output struct {
	Address    string  `json:"address" binding:"required"`
	Amount     []Asset `json:"amount" binding:"required"`
	DataHash   string  `json:"data_hash,omitempty"`
	PlutusData string  `json:"plutus_data,omitempty"`
	ScriptRef  string  `json:"script_ref,omitempty"`
	ScriptHash string  `json:"script_hash,omitempty"`
}

type UTxO struct {
	Input  Input  `json:"input" binding:"required"`
	Output Output `json:"output" binding:"required"`
}

func MakeScriptUtxo(txHash string, outputIndex int, address string, amount []Asset, plutusData string, dataHash string) UTxO {
	return UTxO{
		Input: Input{
			OutputIndex: outputIndex,
			TxHash:      txHash,
		},
		Output: Output{
			Address:    address,
			Amount:     amount,
			DataHash:   dataHash,
			PlutusData: plutusData,
		},
	}
}
