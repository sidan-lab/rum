package builder

import (
	"github.com/sidan-lab/rum/models"
	types "github.com/sidan-lab/rum/models/builder_types"
)

func (builder *TxBuilder) SpendingPlutusScript(languageVersion types.LanguageVersion) *TxBuilder {
	switch languageVersion {
	case types.LanguageVersionV1:
		builder.SpendingPlutusScriptV1()
	case types.LanguageVersionV2:
		builder.SpendingPlutusScriptV2()
	case types.LanguageVersionV3:
		builder.SpendingPlutusScriptV3()
	}
	return builder
}

func (builder *TxBuilder) SpendingPlutusScriptV1() *TxBuilder {
	v := types.LanguageVersionV1
	builder.AddingScriptInput = &v
	return builder
}

func (builder *TxBuilder) SpendingPlutusScriptV2() *TxBuilder {
	v := types.LanguageVersionV2
	builder.AddingScriptInput = &v
	return builder
}

func (builder *TxBuilder) SpendingPlutusScriptV3() *TxBuilder {
	v := types.LanguageVersionV3
	builder.AddingScriptInput = &v
	return builder
}

func (builder *TxBuilder) TxIn(
	txHash string,
	txIndex uint32,
	amount []models.Asset,
	address string,
) *TxBuilder {
	if builder.TxInItem != nil {
		// TODO: queue input
	}

	if builder.AddingScriptInput != nil {
		builder.TxInItem = types.ScriptTxIn{
			TxIn: types.TxInParameter{
				TxHash:  txHash,
				TxIndex: txIndex,
				Amount:  &amount,
				Address: &address,
			},
			ScriptTxIn: types.ScriptTxInParameter{
				ScriptSource: nil,
				DatumSource:  nil,
				Redeemer:     nil,
			},
		}
	} else {
		builder.TxInItem = types.PubKeyTxIn{
			TxIn: types.TxInParameter{
				TxHash:  txHash,
				TxIndex: txIndex,
				Amount:  &amount,
				Address: &address,
			},
		}
	}
	return builder
}

func (builder *TxBuilder) TxInScript(scriptCbor string) *TxBuilder {
	if builder.TxInItem == nil {
		panic("Undefined input")
	}
	switch txInItem := builder.TxInItem.(type) {
	case types.PubKeyTxIn:
		builder.TxInItem = types.SimpleScriptTxIn{
			TxIn: txInItem.TxIn,
			SimpleScriptTxIn: types.ProvidedSimpleScriptTxInSource{
				ScriptCbor: scriptCbor,
			},
		}
	case types.ScriptTxIn:
		if builder.AddingScriptInput == nil {
			panic("Plutus script must have version specified")
		}

		txInItem.ScriptTxIn.ScriptSource = types.ProvidedScriptSource{
			ScriptCbor:      scriptCbor,
			LanguageVersion: *builder.AddingScriptInput,
		}

		builder.AddingScriptInput = nil
		builder.TxInItem = txInItem
	case types.SimpleScriptTxIn:
		txInItem.SimpleScriptTxIn = types.ProvidedSimpleScriptTxInSource{
			ScriptCbor: scriptCbor,
		}
		builder.TxInItem = txInItem
	}
	return builder
}

func (builder *TxBuilder) TxInDatumValue(data WData) *TxBuilder {
	if builder.TxInItem == nil {
		panic("Undefined input")
	}
	switch txInItem := builder.TxInItem.(type) {
	case types.PubKeyTxIn:
		panic("Datum cannot be defined for a pubkey tx in")
	case types.SimpleScriptTxIn:
		panic("Datum cannot be defined for a simple script tx in")
	case types.ScriptTxIn:
		rawData, err := data.ToCbor()
		if err != nil {
			panic("Error converting datum to CBOR")
		}
		txInItem.ScriptTxIn.DatumSource = types.ProvidedDatumSource{
			Data: rawData,
		}
		builder.TxInItem = txInItem
	}
	return builder
}

func (builder *TxBuilder) TxInInlineDatumPresent() *TxBuilder {
	if builder.TxInItem == nil {
		panic("Undefined input")
	}
	switch txInItem := builder.TxInItem.(type) {
	case types.PubKeyTxIn:
		panic("Datum cannot be defined for a pubkey tx in")
	case types.SimpleScriptTxIn:
		panic("Datum cannot be defined for a simple script tx in")
	case types.ScriptTxIn:
		txInItem.ScriptTxIn.DatumSource = types.InlineDatumSource{
			TxHash:  txInItem.TxIn.TxHash,
			TxIndex: txInItem.TxIn.TxIndex,
		}
		builder.TxInItem = txInItem
	}
	return builder
}

func (builder *TxBuilder) TxInRedeemerValue(redeemer WRedeemer) *TxBuilder {
	if builder.TxInItem == nil {
		panic("Undefined input")
	}
	switch txInItem := builder.TxInItem.(type) {
	case types.PubKeyTxIn:
		panic("Redeemer cannot be defined for a pubkey tx in")
	case types.SimpleScriptTxIn:
		panic("Redeemer cannot be defined for a simple script tx in")
	case types.ScriptTxIn:
		rawRedeemer, err := redeemer.Data.ToCbor()
		if err != nil {
			panic("Error converting redeemer to CBOR")
		}
		txInItem.ScriptTxIn.Redeemer = &types.Redeemer{
			Data:    rawRedeemer,
			ExUnits: redeemer.ExUnits,
		}
		builder.TxInItem = txInItem
	}
	return builder
}

func (builder *TxBuilder) SpendingTxInReference(txHash string, txIndex uint32, scriptHash string, scriptSize uint) *TxBuilder {
	if builder.TxInItem == nil {
		panic("Undefined input")
	}
	switch txInItem := builder.TxInItem.(type) {
	case types.PubKeyTxIn:
		panic("Script reference cannot be defined for a pubkey tx in")
	case types.SimpleScriptTxIn:
		panic("Script reference cannot be defined for a simple script tx in")
	case types.ScriptTxIn:
		if builder.AddingScriptInput == nil {
			panic("Plutus script must have version specified")
		}
		txInItem.ScriptTxIn.ScriptSource = types.InlineScriptSource{
			RefTxIn: types.RefTxIn{
				TxHash:  txHash,
				TxIndex: txIndex,
				// Script size is already accounted for in script source
				ScriptSize: nil,
			},
			ScriptHash:      scriptHash,
			LanguageVersion: *builder.AddingScriptInput,
			ScriptSize:      scriptSize,
		}
		builder.TxInItem = txInItem
	}

	return builder
}

func (builder *TxBuilder) SpendingReferenceTxInInlineDatumPresent() *TxBuilder {
	return builder.TxInInlineDatumPresent()
}

func (builder *TxBuilder) SpendingReferenceTxInRedeemerValue(redeemer WRedeemer) *TxBuilder {
	return builder.TxInRedeemerValue(redeemer)
}

func (builder *TxBuilder) ReadOnlyTxInReference(txHash string, txIndex uint32, scriptSize *uint) *TxBuilder {
	builder.TxBuilderBody.ReferenceInputs = append(
		builder.TxBuilderBody.ReferenceInputs,
		types.RefTxIn{
			TxHash:     txHash,
			TxIndex:    txIndex,
			ScriptSize: scriptSize,
		},
	)
	return builder
}

func (builder *TxBuilder) TxInCollateral(
	txHash string,
	txIndex uint32,
	amount []models.Asset,
	address string,
) *TxBuilder {
	if builder.CollateralItem != nil {
		builder.TxBuilderBody.Collaterals = append(
			builder.TxBuilderBody.Collaterals,
			*builder.CollateralItem,
		)
	}
	builder.CollateralItem = &types.PubKeyTxIn{
		TxIn: types.TxInParameter{
			TxHash:  txHash,
			TxIndex: txIndex,
			Amount:  &amount,
			Address: &address,
		},
	}
	return builder
}
