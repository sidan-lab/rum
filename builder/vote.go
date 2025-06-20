package builder

import (
	types "github.com/sidan-lab/rum/models/builder_types"
)

func (builder *TxBuilder) VotingPlutusScript(version types.LanguageVersion) *TxBuilder {
	switch version {
	case types.LanguageVersionV1:
		builder.VotingPlutusScriptV1()
	case types.LanguageVersionV2:
		builder.VotingPlutusScriptV2()
	case types.LanguageVersionV3:
		builder.VotingPlutusScriptV3()
	}
	return builder
}

func (builder *TxBuilder) VotingPlutusScriptV1() *TxBuilder {
	v := types.LanguageVersionV1
	builder.AddingPlutusVote = &v
	return builder
}

func (builder *TxBuilder) VotingPlutusScriptV2() *TxBuilder {
	v := types.LanguageVersionV2
	builder.AddingPlutusVote = &v
	return builder
}

func (builder *TxBuilder) VotingPlutusScriptV3() *TxBuilder {
	v := types.LanguageVersionV3
	builder.AddingPlutusVote = &v
	return builder
}

func (builder *TxBuilder) VoteTxInReference(
	txHash string,
	txIndex uint32,
	voteScriptHash string,
	scriptSize uint,
) *TxBuilder {
	if builder.VoteItem == nil {
		panic("Undefined output")
	}
	switch voteItem := builder.VoteItem.(type) {
	case types.BasicVote:
		panic("Script reference cannot be defined for a pubkey vote")
	case types.SimpleScriptVote:
		builder.TxBuilderBody.Votes = append(
			builder.TxBuilderBody.Votes,
			types.InlineSimpleScriptSource{
				RefTxIn: types.RefTxIn{
					TxHash:  txHash,
					TxIndex: txIndex,
					// Script size is already accounted for in script source
					ScriptSize: nil,
				},
				SimpleScriptHash: voteScriptHash,
				ScriptSize:       scriptSize,
			},
		)
		builder.VoteItem = nil
	case types.ScriptVote:
		if builder.AddingPlutusVote == nil {
			panic("Plutus votes require a language version")
		}
		voteItem.ScriptSource = types.InlineScriptSource{
			RefTxIn: types.RefTxIn{
				TxHash:  txHash,
				TxIndex: txIndex,
				// Script size is already accounted for in script source
				ScriptSize: nil,
			},
			ScriptHash:      voteScriptHash,
			LanguageVersion: *builder.AddingPlutusVote,
			ScriptSize:      scriptSize,
		}
		builder.VoteItem = voteItem
	}
	return builder
}

func (builder *TxBuilder) Vote(
	voter types.Voter,
	govActionID types.RefTxIn,
	votingProcedure types.VotingProcedure,
) *TxBuilder {
	if builder.VoteItem != nil {
		builder.QueueVote()
	}

	if builder.AddingPlutusVote != nil {
		builder.VoteItem = types.ScriptVote{
			Vote: types.VoteType{
				Voter:           voter,
				GovActionID:     govActionID,
				VotingProcedure: votingProcedure,
			},
			Redeemer:     nil,
			ScriptSource: nil,
		}
	} else {
		builder.VoteItem = types.BasicVote{
			Inner: types.VoteType{
				Voter:           voter,
				GovActionID:     govActionID,
				VotingProcedure: votingProcedure,
			},
		}
	}
	return builder
}

// pub fn vote_script(&mut self, script_cbor: &str) -> &mut Self {

func (builder *TxBuilder) VoteScript(scriptCbor string) *TxBuilder {
	if builder.VoteItem == nil {
		panic("Undefined vote")
	}
	voteItem := builder.VoteItem
	builder.VoteItem = nil

	switch voteItem := voteItem.(type) {
	case types.BasicVote:
		panic("Script reference cannot be defined for a pubkey vote")
	case types.SimpleScriptVote:
		voteItem.SimpleScriptSource = types.ProvidedSimpleScriptSource{
			ScriptCbor: scriptCbor,
		}
		builder.VoteItem = voteItem
	case types.ScriptVote:
		if builder.AddingPlutusVote == nil {
			panic("Plutus votes require a language version")
		}
		voteItem.ScriptSource = types.ProvidedScriptSource{
			ScriptCbor:      scriptCbor,
			LanguageVersion: *builder.AddingPlutusVote,
		}
		builder.VoteItem = voteItem
		builder.AddingPlutusVote = nil
	}
	return builder
}

func (builder *TxBuilder) VoteRedeemerValue(redeemer WRedeemer) *TxBuilder {
	if builder.VoteItem == nil {
		panic("Undefined vote")
	}
	voteItem := builder.VoteItem
	builder.VoteItem = nil
	switch voteItem := voteItem.(type) {
	case types.BasicVote:
		panic("Redeemer cannot be defined for a basic vote")
	case types.SimpleScriptVote:
		panic("Redeemer cannot be defined for a native script vote")
	case types.ScriptVote:
		rawRedeemer, err := redeemer.Data.ToCbor()
		if err != nil {
			panic("Error converting redeemer to CBOR")
		}
		voteItem.Redeemer = &types.Redeemer{
			Data:    rawRedeemer,
			ExUnits: redeemer.ExUnits,
		}
		builder.VoteItem = voteItem
	}
	return builder
}

func (builder *TxBuilder) VoteReferenceTxInRedeemerValue(redeemer WRedeemer) *TxBuilder {
	return builder.VoteRedeemerValue(redeemer)
}
