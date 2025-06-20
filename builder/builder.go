// Package builder
package builder

import (
	"github.com/sidan-lab/rum/models"
	types "github.com/sidan-lab/rum/models/builder_types"
)

type TxBuilder struct {
	// serializer: WhiskyCSL,
	TxBuilderBody          *types.TxBuilderBody
	ProtocolParams         *models.Protocol
	TxInItem               types.TxIn
	WithdrawalItem         *types.Withdrawal
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

// pub fn change_output_datum(&mut self, data: WData) -> &mut Self {

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

// pub fn input_for_evaluation(&mut self, input: &UTxO) -> &mut Self {

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

// pub fn queue_input(&mut self) {
//
// pub fn queue_withdrawal(&mut self) {
//
// pub fn queue_vote(&mut self) {
//
// pub fn queue_mint(&mut self) {
//
// pub fn queue_all_last_item(&mut self) {
//
// pub fn add_utxos_from(
