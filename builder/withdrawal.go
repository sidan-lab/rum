package builder

import types "github.com/sidan-lab/rum/models/builder_types"

func (builder *TxBuilder) WithdrawalPlutusScript(languageVersion types.LanguageVersion) *TxBuilder {
	switch languageVersion {
	case types.LanguageVersionV1:
		builder.WithdrawalPlutusScriptV1()
	case types.LanguageVersionV2:
		builder.WithdrawalPlutusScriptV2()
	case types.LanguageVersionV3:
		builder.WithdrawalPlutusScriptV3()
	}
	return builder
}

func (builder *TxBuilder) WithdrawalPlutusScriptV1() *TxBuilder {
	v := types.LanguageVersionV1
	builder.AddingPlutusWithdrawal = &v
	return builder
}

func (builder *TxBuilder) WithdrawalPlutusScriptV2() *TxBuilder {
	v := types.LanguageVersionV2
	builder.AddingPlutusWithdrawal = &v
	return builder
}

func (builder *TxBuilder) WithdrawalPlutusScriptV3() *TxBuilder {
	v := types.LanguageVersionV3
	builder.AddingPlutusWithdrawal = &v
	return builder
}

func (builder *TxBuilder) WithdrawalTxInReference(
	txHash string,
	txIndex uint32,
	withdrawlScriptHash string,
	scriptSize uint,
) *TxBuilder {
	if builder.WithdrawalItem == nil {
		panic("Undefined output")
	}
	withdrawalItem := builder.WithdrawalItem
	builder.WithdrawalItem = nil
	switch withdrawalItem := withdrawalItem.(type) {
	case types.PubKeyWithdrawal:
		panic("Script reference cannot be defined for a pubkey withdrawal")
	case types.SimpleScriptWithdrawal:
		withdrawalItem.ScriptSource = types.InlineSimpleScriptSource{
			RefTxIn: types.RefTxIn{
				TxHash:     txHash,
				TxIndex:    txIndex,
				ScriptSize: nil,
			},
			SimpleScriptHash: withdrawlScriptHash,
			ScriptSize:       scriptSize,
		}
		builder.WithdrawalItem = withdrawalItem
	case types.PlutusScriptWithdrawal:
		if builder.AddingPlutusWithdrawal == nil {
			panic("Plutus withdrawals require a language version")
		}
		withdrawalItem.ScriptSource = types.InlineScriptSource{
			RefTxIn: types.RefTxIn{
				TxHash:     txHash,
				TxIndex:    txIndex,
				ScriptSize: nil,
			},
			ScriptHash:      withdrawlScriptHash,
			LanguageVersion: *builder.AddingPlutusWithdrawal,
			ScriptSize:      scriptSize,
		}
		builder.WithdrawalItem = withdrawalItem
	}
	return builder
}

func (builder *TxBuilder) Withdrawal(stakeAddress string, coin uint64) *TxBuilder {
	if builder.WithdrawalItem != nil {
		builder.QueueWithdrawal()
	}
	if builder.AddingPlutusWithdrawal != nil {
		builder.WithdrawalItem = types.PlutusScriptWithdrawal{
			Address:      stakeAddress,
			Coin:         coin,
			ScriptSource: nil,
			Redeemer:     nil,
		}
	} else {
		builder.WithdrawalItem = types.PubKeyWithdrawal{
			Address: stakeAddress,
			Coin:    coin,
		}
	}
	return builder
}

func (builder *TxBuilder) WithdrawalScript(scriptCbor string) *TxBuilder {
	if builder.WithdrawalItem == nil {
		panic("Undefined withdrawal")
	}
	withdrawalItem := builder.WithdrawalItem
	builder.WithdrawalItem = nil
	switch withdrawalItem := withdrawalItem.(type) {
	case types.PubKeyWithdrawal:
		panic("Script cannot be defined for a pubkey withdrawal")
	case types.SimpleScriptWithdrawal:
		withdrawalItem.ScriptSource = types.ProvidedSimpleScriptSource{
			ScriptCbor: scriptCbor,
		}
		builder.WithdrawalItem = withdrawalItem
	case types.PlutusScriptWithdrawal:
		if builder.AddingPlutusWithdrawal == nil {
			panic("Plutus withdrawals require a language version")
		}
		withdrawalItem.ScriptSource = types.ProvidedScriptSource{
			ScriptCbor:      scriptCbor,
			LanguageVersion: *builder.AddingPlutusWithdrawal,
		}
		builder.WithdrawalItem = withdrawalItem
		builder.AddingPlutusWithdrawal = nil
	}

	return builder
}

func (builder *TxBuilder) WithdrawalRedeemerValue(redeemer WRedeemer) *TxBuilder {
	if builder.WithdrawalItem == nil {
		panic("Undefined withdrawal")
	}
	withdrawalItem := builder.WithdrawalItem
	builder.WithdrawalItem = nil
	switch withdrawalItem := withdrawalItem.(type) {
	case types.PubKeyWithdrawal:
		panic("Redeemer cannot be defined for a pubkey withdrawal")
	case types.SimpleScriptWithdrawal:
		panic("Redeemer cannot be defined for a native script withdrawal")
	case types.PlutusScriptWithdrawal:
		rawRedeemer, err := redeemer.Data.ToCbor()
		if err != nil {
			panic("Error converting redeemer to CBOR")
		}

		withdrawalItem.Redeemer = &types.Redeemer{
			Data:    rawRedeemer,
			ExUnits: redeemer.ExUnits,
		}
		builder.WithdrawalItem = withdrawalItem
	}
	return builder
}

func (builder *TxBuilder) WithdrawalReferenceTxInRedeemerValue(redeemer WRedeemer) *TxBuilder {
	return builder.WithdrawalRedeemerValue(redeemer)
}
