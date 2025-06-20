// Package builder
package builder

import (
	"fmt"

	"github.com/sidan-lab/rum/models"
	types "github.com/sidan-lab/rum/models/builder_types"
)

type TxBuilder struct {
	// serializer: WhiskyCSL,
	TxBuilderBody          *types.TxBuilderBody
	ProtocolParams         *models.Protocol
	TxInItem               types.TxIn
	WithdrawalItem         types.Withdrawal
	VoteItem               types.Vote
	MintItem               types.MintItem
	CollateralItem         *types.PubKeyTxIn
	TxOutput               *types.Output
	AddingScriptInput      *types.LanguageVersion
	AddingPlutusMint       *types.LanguageVersion
	AddingPlutusWithdrawal *types.LanguageVersion
	AddingPlutusVote       *types.LanguageVersion
	// Fetcher: Option<Box<dyn Fetcher>>,
	// Evaluator: Option<Box<dyn Evaluator>>,
	// Submitter: Option<Box<dyn Submitter>>,
	ExtraInputs         []types.UTxO
	SelectionThreshold  uint64
	ChainedTxs          []string
	InputsForEvaluation map[string]types.UTxO
}

type TxBuilderParam struct {
	// pub evaluator: Option<Box<dyn Evaluator>>,
	// pub fetcher: Option<Box<dyn Fetcher>>,
	// pub submitter: Option<Box<dyn Submitter>>,
	Params *models.Protocol
}

func New(param TxBuilderParam) *TxBuilder {
	return &TxBuilder{
		// serializer: WhiskyCSL::new(param.params.clone()).unwrap(),
		TxBuilderBody:          types.NewTxBuilderBody(),
		ProtocolParams:         param.Params,
		TxInItem:               nil,
		WithdrawalItem:         nil,
		VoteItem:               nil,
		MintItem:               nil,
		CollateralItem:         nil,
		TxOutput:               nil,
		AddingScriptInput:      nil,
		AddingPlutusMint:       nil,
		AddingPlutusWithdrawal: nil,
		AddingPlutusVote:       nil,
		// fetcher: param.fetcher,
		// evaluator: match param.evaluator {
		// 	Some(evaluator) => Some(evaluator),
		// 	None => Some(Box::new(OfflineTxEvaluator::new())),
		// },
		// submitter: param.submitter,
		ExtraInputs:         []types.UTxO{},
		SelectionThreshold:  5_000_000,
		ChainedTxs:          []string{},
		InputsForEvaluation: map[string]types.UTxO{},
	}
}

func NewCore() *TxBuilder {
	return New(TxBuilderParam{
		Params: nil,
	})
}

func (builder *TxBuilder) RequiredSignerHash(pubKeyHash string) *TxBuilder {
	builder.TxBuilderBody.RequiredSignatures = append(builder.TxBuilderBody.RequiredSignatures, pubKeyHash)
	return builder
}

func (builder *TxBuilder) ChangeAddress(address string) *TxBuilder {
	builder.TxBuilderBody.ChangeAddress = address
	return builder
}

func (builder *TxBuilder) ChangeOutputDatum(data WData) *TxBuilder {
	rawData, err := data.ToCbor()
	if err != nil {
		panic("Error converting datum to CBOR")
	}
	builder.TxBuilderBody.ChangeDatum = types.InlineDatum{
		Inner: rawData,
	}
	return builder
}

func (builder *TxBuilder) InvalidBefore(slot uint64) *TxBuilder {
	builder.TxBuilderBody.ValidityRange.InvalidBefore = &slot
	return builder
}

func (builder *TxBuilder) InvalidHereafter(slot uint64) *TxBuilder {
	builder.TxBuilderBody.ValidityRange.InvalidHereafter = &slot
	return builder
}

func (builder *TxBuilder) MetadataValue(tag string, metadata string) *TxBuilder {
	builder.TxBuilderBody.Metadata = append(builder.TxBuilderBody.Metadata, types.Metadata{
		Tag:      tag,
		Metadata: metadata,
	})
	return builder
}

func (builder *TxBuilder) SigningKey(skeyHex string) *TxBuilder {
	builder.TxBuilderBody.SigningKey = append(builder.TxBuilderBody.SigningKey, skeyHex)
	return builder
}

func (builder *TxBuilder) ChainTx(txHex string) *TxBuilder {
	builder.ChainedTxs = append(builder.ChainedTxs, txHex)
	return builder
}

func (builder *TxBuilder) InputForEvaluation(input types.UTxO) *TxBuilder {
	utxoID := fmt.Sprintf("%s%d", input.Input.TxHash, input.Input.OutputIndex)
	currentUtxo, ok := builder.InputsForEvaluation[utxoID]
	if ok {
		dataHash := input.Output.DataHash
		if dataHash == nil {
			dataHash = currentUtxo.Output.DataHash
		}

		plutusData := input.Output.PlutusData
		if plutusData == nil {
			plutusData = currentUtxo.Output.PlutusData
		}

		scriptRef := input.Output.ScriptRef
		if scriptRef == nil {
			scriptRef = currentUtxo.Output.ScriptRef
		}

		scriptHash := input.Output.ScriptHash
		if scriptHash == nil {
			scriptHash = currentUtxo.Output.ScriptHash
		}

		updatedUtxo := types.UTxO{
			Output: types.UtxoOutput{
				Address:    input.Output.Address,
				Amount:     input.Output.Amount,
				DataHash:   dataHash,
				PlutusData: plutusData,
				ScriptRef:  scriptRef,
				ScriptHash: scriptHash,
			},
			Input: input.Input,
		}
		builder.InputsForEvaluation[utxoID] = updatedUtxo
	} else {
		builder.InputsForEvaluation[utxoID] = input
	}

	return builder
}

func (builder *TxBuilder) SelectUtxosFrom(extraInputs []types.UTxO, threshold uint64) *TxBuilder {
	builder.SelectionThreshold = threshold
	builder.ExtraInputs = append(builder.ExtraInputs, extraInputs...)
	return builder
}

func (builder *TxBuilder) SetFee(fee string) *TxBuilder {
	builder.TxBuilderBody.Fee = &fee
	return builder
}

func (builder *TxBuilder) Network(network types.Network) *TxBuilder {
	builder.TxBuilderBody.Network = &network
	return builder
}

func (builder *TxBuilder) QueueInput() *TxBuilder {
	switch txInItem := builder.TxInItem.(type) {
	case types.ScriptTxIn:
		if txInItem.ScriptTxIn.DatumSource == nil {
			panic("Datum in a script input cannot be None")
		}
		if txInItem.ScriptTxIn.Redeemer == nil {
			panic("Redeemer in script input cannot be None")
		}
		if txInItem.ScriptTxIn.ScriptSource == nil {
			panic("Script source in script input cannot be None")
		}
	case types.SimpleScriptTxIn:
	case types.PubKeyTxIn:
	}

	input := builder.TxInItem
	builder.InputForEvaluation(input.ToUTxO())
	builder.TxBuilderBody.Inputs = append(builder.TxBuilderBody.Inputs, input)
	builder.TxInItem = nil
	return builder
}

func (builder *TxBuilder) QueueWithdrawal() *TxBuilder {
	switch withdrawalItem := builder.WithdrawalItem.(type) {
	case types.PlutusScriptWithdrawal:
		if withdrawalItem.Redeemer == nil {
			panic("Redeemer in script withdrawal cannot be None")
		}
		if withdrawalItem.ScriptSource == nil {
			panic("Script source in script withdrawal cannot be None")
		}
	case types.SimpleScriptWithdrawal:
		if withdrawalItem.ScriptSource == nil {
			panic("Script source missing from native script withdrawal")
		}
	case types.PubKeyWithdrawal:
	}
	builder.TxBuilderBody.Withdrawals = append(builder.TxBuilderBody.Withdrawals, builder.WithdrawalItem)
	builder.WithdrawalItem = nil
	return builder
}

func (builder *TxBuilder) QueueVote() *TxBuilder {
	switch voteItem := builder.VoteItem.(type) {
	case types.ScriptVote:
		if voteItem.Redeemer == nil {
			panic("Redeemer in script vote cannot be None")
		}
		if voteItem.ScriptSource == nil {
			panic("Script source in script vote cannot be None")
		}
	case types.SimpleScriptVote:
		if voteItem.SimpleScriptSource == nil {
			panic("Script source is missing from native script vote")
		}
	case types.BasicVote:
	}
	builder.TxBuilderBody.Votes = append(builder.TxBuilderBody.Votes, builder.VoteItem)
	builder.VoteItem = nil
	return builder
}

func (builder *TxBuilder) QueueMint() *TxBuilder {
	switch mintItem := builder.MintItem.(type) {
	case types.ScriptMint:
		if mintItem.ScriptSource == nil {
			panic("Missing mint script information")
		}
		builder.TxBuilderBody.Mints = append(builder.TxBuilderBody.Mints, mintItem)
	case types.SimpleScriptMint:
		if mintItem.ScriptSource == nil {
			panic("Missing mint script information")
		}
		builder.TxBuilderBody.Mints = append(builder.TxBuilderBody.Mints, mintItem)
	}
	builder.MintItem = nil
	return builder
}

func (builder *TxBuilder) QueueAllLastItem() *TxBuilder {
	if builder.TxOutput != nil {
		builder.TxBuilderBody.Outputs = append(builder.TxBuilderBody.Outputs, *builder.TxOutput)
		builder.TxOutput = nil
	}
	if builder.TxInItem != nil {
		builder.QueueInput()
	}
	if builder.CollateralItem != nil {
		builder.TxBuilderBody.Collaterals = append(
			builder.TxBuilderBody.Collaterals,
			*builder.CollateralItem,
		)
		builder.CollateralItem = nil
	}
	if builder.WithdrawalItem != nil {
		builder.QueueWithdrawal()
	}
	if builder.VoteItem != nil {
		builder.QueueVote()
	}
	if builder.MintItem != nil {
		builder.QueueMint()
	}
	return builder
}

// func (builder *TxBuilder) AddUTxOsFrom(extraInputs []types.UTxO, threshold uint64) *TxBuilder {
//
// 	return builder
// }
