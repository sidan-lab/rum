package builder_types

type Vote interface {
	isVote()
}

type BasicVote struct {
	Inner VoteType `json:"basicVote"`
}

func (BasicVote) isVote() {}

type ScriptVote struct {
	Vote         VoteType     `json:"vote"`
	Redeemer     *Redeemer    `json:"redeemer"`
	ScriptSource ScriptSource `json:"scriptSource"`
}

type SimpleScriptVote struct {
	Vote               VoteType           `json:"vote"`
	SimpleScriptSource SimpleScriptSource `json:"simpleScriptSource"`
}

type VoteType struct {
	Voter           Voter           `json:"voter"`
	GovActionId     RefTxIn         `json:"govActionId"`
	VotingProcedure VotingProcedure `json:"votingProcedure"`
}

type Voter interface {
	isVoter()
}

type ConstitutionalCommitteeHotCred struct {
	Inner Credential `json:"constitutionalCommitteeHotCred"`
}

func (ConstitutionalCommitteeHotCred) isVoter() {}

type DRepId struct {
	Inner string `json:"dRepId"`
}

func (DRepId) isVoter() {}
func (DRepId) isDRep() {}

type StakingPoolKeyHash struct {
	Inner string `json:"stakingPoolKeyHash"`
}

func (StakingPoolKeyHash) isVoter() {}

type VotingProcedure struct {
	VoteKind VoteKind `json:"voteKind"`
	Anchor   *Anchor  `json:"anchor"`
}

type VoteKind uint

const (
	VoteKindNo      = 0
	VoteKindYes     = 1
	VoteKindAbstain = 2
)
