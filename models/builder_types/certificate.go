// Package builder_types
package builder_types

import "encoding/json"

type Certificate interface {
	isCertificate()
}

type BasicCertificate struct {
	Inner CertificateType `json:"basicCertificate"`
}

func (BasicCertificate) isCertificate() {}

type ScriptCertificate struct {
	Cert         CertificateType `json:"cert"`
	Redeemer     *Redeemer       `json:"redeemer"`
	ScriptSource ScriptSource    `json:"scriptSource"`
}

func (ScriptCertificate) isCertificate() {}

type SimpleScriptCertificate struct {
	Cert               CertificateType    `json:"cert"`
	SimpleScriptSource SimpleScriptSource `json:"simpleScriptSource"`
}

func (SimpleScriptCertificate) isCertificate() {}

type CertificateType interface {
	isCertificateType()
}

type RegisterPool struct {
	PoolParams PoolParams `json:"poolParams"`
}

func (RegisterPool) isCertificateType() {}
func (r RegisterPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"registerPool": r,
	})
}

type PoolParams struct {
	VrfKeyHash    string        `json:"vrfKeyHash"`
	Operator      string        `json:"operator"`
	Pledge        string        `json:"pledge"`
	Cost          string        `json:"cost"`
	Margin        [2]uint64     `json:"margin"`
	Relays        []Relay       `json:"relays"`
	Owners        []string      `json:"owners"`
	RewardAddress string        `json:"rewardAddress"`
	Metadata      *PoolMetadata `json:"metadata"`
}

type Anchor struct {
	AnchorUrl      string `json:"anchorUrl"`
	AnchorDataHash string `json:"anchorDataHash"`
}

type Relay interface {
	isRelay()
}

type SingleHostAddr struct {
	Ipv4 *string `json:"ipv4"`
	Ipv6 *string `json:"ipv6"`
	Port *uint16 `json:"port"`
}

func (SingleHostAddr) isRelay() {}
func (s SingleHostAddr) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"singleHostAddr": s,
	})
}

type SingleHostName struct {
	DomainName string  `json:"domainName"`
	Port       *uint16 `json:"port"`
}

func (SingleHostName) isRelay() {}
func (s SingleHostName) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"singleHostName": s,
	})
}

type MultiHostName struct {
	DomainName string `json:"domainName"`
}

func (MultiHostName) isRelay() {}
func (m MultiHostName) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"multiHostName": m,
	})
}

type PoolMetadata struct {
	Url  string `json:"url"`
	Hash string `json:"hash"`
}

type RegisterStake struct {
	StakeKeyAddress string `json:"stakeKeyAddress"`
	Coin            uint64 `json:"coin"`
}

func (RegisterStake) isCertificateType() {}
func (r RegisterStake) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"registerStake": r,
	})
}

type DelegateStake struct {
	StakeKeyAddress string `json:"stakeKeyAddress"`
	PoolID          string `json:"poolId"`
}

func (DelegateStake) isCertificateType() {}
func (d DelegateStake) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"delegateStake": d,
	})
}

type DeregisterStake struct {
	StakeKeyAddress string `json:"stakeKeyAddress"`
}

func (DeregisterStake) isCertificateType() {}
func (d DeregisterStake) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"deregisterStake": d,
	})
}

type RetirePool struct {
	PoolID string `json:"poolId"`
	Epoch  uint32 `json:"epoch"`
}

func (RetirePool) isCertificateType() {}
func (r RetirePool) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"retirePool": r,
	})
}

type VoteDelegation struct {
	StakeKeyAddress string `json:"stakeKeyAddress"`
	DRep            DRep   `json:"drep"`
}

func (VoteDelegation) isCertificateType() {}
func (v VoteDelegation) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"voteDelegation": v,
	})
}

type StakeAndVoteDelegation struct {
	StakeKeyAddress string `json:"stakeKeyAddress"`
	PoolKeyHash     string `json:"poolKeyHash"`
	DRep            DRep   `json:"drep"`
}

func (StakeAndVoteDelegation) isCertificateType() {}
func (s StakeAndVoteDelegation) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"stakeAndVoteDelegation": s,
	})
}

type StakeRegistrationAndDelegation struct {
	StakeKeyAddress string `json:"stakeKeyAddress"`
	PoolKeyHash     string `json:"poolKeyHash"`
	Coin            uint64 `json:"coin"`
}

func (StakeRegistrationAndDelegation) isCertificateType() {}
func (s StakeRegistrationAndDelegation) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"stakeRegistrationAndDelegation": s,
	})
}

type VoteRegistrationAndDelegation struct {
	StakeKeyAddress string `json:"stakeKeyAddress"`
	DRep            DRep   `json:"drep"`
	Coin            uint64 `json:"coin"`
}

func (VoteRegistrationAndDelegation) isCertificateType() {}
func (v VoteRegistrationAndDelegation) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"voteRegistrationAndDelegation": v,
	})
}

type StakeVoteRegistrationAndDelegation struct {
	StakeKeyAddress string `json:"stakeKeyAddress"`
	PoolKeyHash     string `json:"poolKeyHash"`
	DRep            DRep   `json:"drep"`
	Coin            uint64 `json:"coin"`
}

func (StakeVoteRegistrationAndDelegation) isCertificateType() {}
func (s StakeVoteRegistrationAndDelegation) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"stakeVoteRegistrationAndDelegation": s,
	})
}

type DRep interface {
	isDRep()
}

type AlwaysAbstain struct{}

func (AlwaysAbstain) isDRep() {}
func (AlwaysAbstain) MarshalJSON() ([]byte, error) {
	return json.Marshal("alwaysAbstain")
}

type AlwaysNoConfidence struct{}

func (AlwaysNoConfidence) isDRep() {}
func (AlwaysNoConfidence) MarshalJSON() ([]byte, error) {
	return json.Marshal("alwaysNoConfidence")
}

type CommitteeHotAuth struct {
	CommitteeColdKeyAddress string `json:"committeeColdKeyAddress"`
	CommitteeHotKeyAddress  string `json:"committeeHotKeyAddress"`
}

func (CommitteeHotAuth) isCertificateType() {}
func (c CommitteeHotAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"committeeHotAuth": c,
	})
}

type CommitteeColdResign struct {
	CommitteeColdKeyAddress string  `json:"committeeColdKeyAddress"`
	Anchor                  *Anchor `json:"anchor"`
}

func (CommitteeColdResign) isCertificateType() {}
func (c CommitteeColdResign) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"committeeColdResign": c,
	})
}

type DRepRegistration struct {
	DrepID string  `json:"drepId"`
	Coin   uint64  `json:"coin"`
	Anchor *Anchor `json:"anchor"`
}

func (DRepRegistration) isCertificateType() {}
func (d DRepRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"drepRegistration": d,
	})
}

type DRepDeregistration struct {
	DrepID string `json:"drepId"`
	Coin   uint64 `json:"coin"`
}

func (DRepDeregistration) isCertificateType() {}
func (d DRepDeregistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"drepDeregistration": d,
	})
}

type DRepUpdate struct {
	DrepID string  `json:"drepId"`
	Anchor *Anchor `json:"anchor"`
}

func (DRepUpdate) isCertificateType() {}
func (d DRepUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"drepUpdate": d,
	})
}
