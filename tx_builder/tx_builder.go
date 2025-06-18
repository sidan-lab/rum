package txbuilder

import (
	"github.com/sidan-lab/rum/models"
	types "github.com/sidan-lab/rum/models/builder_types"
)

type TxBuilder struct {
	// serializer: WhiskyCSL,
	TxBuilderBody          *types.TxBuilderBody
	ProtocolParams         *models.Protocol
	TxInItem               *types.TxIn
	WithdrawalItem         *types.Withdrawal
	VoteItem               *types.Vote
	MintItem               *types.MintItem
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

func NewTxBuilder(param TxBuilderParam) *TxBuilder {
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
