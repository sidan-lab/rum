package builder

import types "github.com/sidan-lab/rum/models/builder_types"

func (builder *TxBuilder) MintPlutusSctipt(languageVersion types.LanguageVersion) *TxBuilder {
	switch languageVersion {
	case types.LanguageVersionV1:
		builder.MintPlutusScriptV1()
	case types.LanguageVersionV2:
		builder.MintPlutusScriptV2()
	case types.LanguageVersionV3:
		builder.MintPlutusScriptV3()
	}
	return builder
}

func (builder *TxBuilder) MintPlutusScriptV1() *TxBuilder {
	v := types.LanguageVersionV1
	builder.AddingPlutusMint = &v
	return builder
}

func (builder *TxBuilder) MintPlutusScriptV2() *TxBuilder {
	v := types.LanguageVersionV2
	builder.AddingPlutusMint = &v
	return builder
}

func (builder *TxBuilder) MintPlutusScriptV3() *TxBuilder {
	v := types.LanguageVersionV3
	builder.AddingPlutusMint = &v
	return builder
}

func (builder *TxBuilder) Mint(quantity int64, policy string, name string) *TxBuilder {
	if builder.MintItem != nil {
		builder.QueueMint()
	}

	if builder.AddingPlutusMint != nil {
		builder.MintItem = types.ScriptMint{
			Mint: types.MintParameter{
				PolicyID:  policy,
				AssetName: name,
				Amount:    quantity,
			},
			Redeemer:     nil,
			ScriptSource: nil,
		}
	} else {
		builder.MintItem = types.SimpleScriptMint{
			Mint: types.MintParameter{
				PolicyID:  policy,
				AssetName: name,
				Amount:    quantity,
			},
			ScriptSource: nil,
		}
	}
	return builder
}

func (builder *TxBuilder) MintingScript(scriptCbor string) *TxBuilder {
	if builder.MintItem == nil {
		panic("Undefined mint")
	}

	switch mintItem := builder.MintItem.(type) {
	case types.ScriptMint:
		if builder.AddingPlutusMint == nil {
			panic("Plutus mints must have version specified")
		}
		mintItem.ScriptSource = types.ProvidedScriptSource{
			ScriptCbor:      scriptCbor,
			LanguageVersion: *builder.AddingPlutusMint,
		}
		builder.AddingPlutusMint = nil
		builder.MintItem = mintItem
	case types.SimpleScriptMint:
		mintItem.ScriptSource = types.ProvidedSimpleScriptSource{
			ScriptCbor: scriptCbor,
		}
		builder.MintItem = mintItem
	}
	return builder
}

func (builder *TxBuilder) MintTxInReference(txHash string, txIndex uint32, scriptHash string, scriptSize uint) *TxBuilder {
	if builder.MintItem == nil {
		panic("Undefined mint")
	}

	switch mintItem := builder.MintItem.(type) {
	case types.ScriptMint:
		if builder.AddingPlutusMint == nil {
			panic("Plutus mints must have version specified")
		}
		mintItem.ScriptSource = types.InlineScriptSource{
			RefTxIn: types.RefTxIn{
				TxHash:  txHash,
				TxIndex: txIndex,
				// Script size is already accounted for in script source
				ScriptSize: nil,
			},
			ScriptHash:      scriptHash,
			LanguageVersion: *builder.AddingPlutusMint,
			ScriptSize:      scriptSize,
		}
		builder.MintItem = mintItem
	case types.SimpleScriptMint:
		mintItem.ScriptSource = types.InlineSimpleScriptSource{
			RefTxIn: types.RefTxIn{
				TxHash:  txHash,
				TxIndex: txIndex,
				// Script size is already accounted for in script source
				ScriptSize: nil,
			},
			SimpleScriptHash: scriptHash,
			ScriptSize:       scriptSize,
		}
		builder.MintItem = mintItem
	}
	return builder
}

func (builder *TxBuilder) MintRedeemerValue(redeemer WRedeemer) *TxBuilder {
	if builder.MintItem == nil {
		panic("Undefined mint")
	}

	rawRedeemer, err := redeemer.Data.ToCbor()
	if err != nil {
		panic("Error converting redeemer to CBOR")
	}

	switch mintItem := builder.MintItem.(type) {
	case types.ScriptMint:
		mintItem.Redeemer = &types.Redeemer{
			Data:    rawRedeemer,
			ExUnits: redeemer.ExUnits,
		}
		builder.MintItem = mintItem
	case types.SimpleScriptMint:
		panic("Redeemer values cannot be defined for native script mints")
	}
	return builder
}

func (builder *TxBuilder) MintTxInRedeemerValue(redeemer WRedeemer) *TxBuilder {
	return builder.MintRedeemerValue(redeemer)
}
