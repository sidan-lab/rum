package builder_types

import "github.com/sidan-lab/rum/models"

type TxIn interface {
	isTxIn()
	ToUTxO() UTxO
}

type PubKeyTxIn struct {
	TxIn TxInParameter `json:"txIn"`
}

func (PubKeyTxIn) isTxIn() {}

func (txIn PubKeyTxIn) ToUTxO() UTxO {
	return UTxO{
		Input: UtxoInput{
			OutputIndex: txIn.TxIn.TxIndex,
			TxHash:      txIn.TxIn.TxHash,
		},
		Output: UtxoOutput{
			Address:    *txIn.TxIn.Address,
			Amount:     *txIn.TxIn.Amount,
			DataHash:   nil,
			PlutusData: nil,
			ScriptRef:  nil,
			ScriptHash: nil,
		},
	}
}

type SimpleScriptTxIn struct {
	TxIn             TxInParameter             `json:"txIn"`
	SimpleScriptTxIn SimpleScriptTxInParameter `json:"simpleScriptTxIn"`
}

func (SimpleScriptTxIn) isTxIn() {}

func (txIn SimpleScriptTxIn) ToUTxO() UTxO {
	return UTxO{
		Input: UtxoInput{
			OutputIndex: txIn.TxIn.TxIndex,
			TxHash:      txIn.TxIn.TxHash,
		},
		Output: UtxoOutput{
			Address:    *txIn.TxIn.Address,
			Amount:     *txIn.TxIn.Amount,
			DataHash:   nil,
			PlutusData: nil,
			ScriptRef:  nil,
			ScriptHash: nil,
		},
	}
}

type ScriptTxIn struct {
	TxIn       TxInParameter       `json:"txIn"`
	ScriptTxIn ScriptTxInParameter `json:"scriptTxIn"`
}

func (ScriptTxIn) isTxIn() {}

func (txIn ScriptTxIn) ToUTxO() UTxO {
	return UTxO{
		Input: UtxoInput{
			OutputIndex: txIn.TxIn.TxIndex,
			TxHash:      txIn.TxIn.TxHash,
		},
		Output: UtxoOutput{
			Address:    *txIn.TxIn.Address,
			Amount:     *txIn.TxIn.Amount,
			DataHash:   nil,
			PlutusData: nil,
			ScriptRef:  nil,
			ScriptHash: nil,
		},
	}
}

type RefTxIn struct {
	TxHash     string `json:"txHash"`
	TxIndex    uint32 `json:"txIndex"`
	ScriptSize *uint  `json:"scriptSize"`
}

type TxInParameter struct {
	TxHash  string          `json:"txHash"`
	TxIndex uint32          `json:"txIndex"`
	Amount  *[]models.Asset `json:"amount"`
	Address *string         `json:"address"`
}

type SimpleScriptTxInParameter interface {
	isSimpleScriptTxInParameter()
}

type ProvidedSimpleScriptTxInSource struct {
	ScriptCbor string `json:"scriptCbor"`
}

func (ProvidedSimpleScriptTxInSource) isSimpleScriptTxInParameter() {}

type InlineSimpleScriptTxInSource struct {
	RefTxIn          RefTxIn `json:"refTxIn"`
	SimpleScriptHash string  `json:"simpleScriptHash"`
	ScriptSize       uint    `json:"scriptSize"`
}

func (InlineSimpleScriptTxInSource) isSimpleScriptTxInParameter() {}

type ScriptTxInParameter struct {
	ScriptSource ScriptSource `json:"scriptSource"`
	DatumSource  DatumSource  `json:"datumSource"`
	Redeemer     *Redeemer    `json:"redeemer"`
}
