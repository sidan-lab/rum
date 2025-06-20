package builder

import (
	types "github.com/sidan-lab/rum/models/builder_types"
)

func (builder *TxBuilder) RegisterPoolCertificate(poolParams *types.PoolParams) *TxBuilder {
	builder.TxBuilderBody.Certificates = append(builder.TxBuilderBody.Certificates, types.BasicCertificate{
		Inner: types.RegisterPool{
			PoolParams: *poolParams,
		},
	})
	return builder
}

func (builder *TxBuilder) RegisterStakeCertificate(stakeKeyAddress string) *TxBuilder {
	builder.TxBuilderBody.Certificates = append(builder.TxBuilderBody.Certificates, types.BasicCertificate{
		Inner: types.RegisterStake{
			StakeKeyAddress: stakeKeyAddress,
			Coin:            builder.ProtocolParams.KeyDeposit,
		},
	})
	return builder
}

func (builder *TxBuilder) DelegateStakeCertificate(
	stakeKeyAddress string,
	poolID string,
) *TxBuilder {
	builder.TxBuilderBody.Certificates = append(builder.TxBuilderBody.Certificates, types.BasicCertificate{
		Inner: types.DelegateStake{
			StakeKeyAddress: stakeKeyAddress,
			PoolID:          poolID,
		},
	})
	return builder
}

func (builder *TxBuilder) DeregisterStakeCertificate(stakeKeyAddress string) *TxBuilder {
	builder.TxBuilderBody.Certificates = append(builder.TxBuilderBody.Certificates, types.BasicCertificate{
		Inner: types.DeregisterStake{
			StakeKeyAddress: stakeKeyAddress,
		},
	})
	return builder
}

func (builder *TxBuilder) RetirePoolCertificate(poolID string, epoch uint32) *TxBuilder {
	builder.TxBuilderBody.Certificates = append(builder.TxBuilderBody.Certificates, types.BasicCertificate{
		Inner: types.RetirePool{
			PoolID: poolID,
			Epoch:  epoch,
		},
	})
	return builder
}

func (builder *TxBuilder) VoteDelegationCertificate(
	stakeKeyAddress string,
	drep types.DRep,
) *TxBuilder {
	builder.TxBuilderBody.Certificates = append(builder.TxBuilderBody.Certificates, types.BasicCertificate{
		Inner: types.VoteDelegation{
			StakeKeyAddress: stakeKeyAddress,
			DRep:            drep,
		},
	})
	return builder
}

func (builder *TxBuilder) StakeAndVoteDelegationCertificate(
	stakeKeyAddress string,
	poolKeyHash string,
	drep types.DRep,
) *TxBuilder {
	builder.TxBuilderBody.Certificates = append(builder.TxBuilderBody.Certificates, types.BasicCertificate{
		Inner: types.StakeAndVoteDelegation{
			StakeKeyAddress: stakeKeyAddress,
			PoolKeyHash:     poolKeyHash,
			DRep:            drep,
		},
	})
	return builder
}

func (builder *TxBuilder) StakeRegistrationAndDelegation(
	stakeKeyAddress string,
	poolKeyHash string,
	coin uint64,
) *TxBuilder {
	builder.TxBuilderBody.Certificates = append(builder.TxBuilderBody.Certificates, types.BasicCertificate{
		Inner: types.StakeRegistrationAndDelegation{
			StakeKeyAddress: stakeKeyAddress,
			PoolKeyHash:     poolKeyHash,
			Coin:            coin,
		},
	})
	return builder
}

func (builder *TxBuilder) VoteRegistrationAndDelegation(
	stakeKeyAddress string,
	drep types.DRep,
	coin uint64,
) *TxBuilder {
	builder.TxBuilderBody.Certificates = append(builder.TxBuilderBody.Certificates, types.BasicCertificate{
		Inner: types.VoteRegistrationAndDelegation{
			StakeKeyAddress: stakeKeyAddress,
			DRep:            drep,
			Coin:            coin,
		},
	})
	return builder
}

func (builder *TxBuilder) StakeVoteRegistrationAndDelegation(
	stakeKeyAddress string,
	poolKeyHash string,
	drep types.DRep,
	coin uint64,
) *TxBuilder {
	builder.TxBuilderBody.Certificates = append(builder.TxBuilderBody.Certificates, types.BasicCertificate{
		Inner: types.StakeVoteRegistrationAndDelegation{
			StakeKeyAddress: stakeKeyAddress,
			PoolKeyHash:     poolKeyHash,
			DRep:            drep,
			Coin:            coin,
		},
	})
	return builder
}

func (builder *TxBuilder) CommitteeHotAuth(
	committeeColdKeyAddress string,
	committeeHotKeyAddress string,
) *TxBuilder {
	builder.TxBuilderBody.Certificates = append(builder.TxBuilderBody.Certificates, types.BasicCertificate{
		Inner: types.CommitteeHotAuth{
			CommitteeColdKeyAddress: committeeColdKeyAddress,
			CommitteeHotKeyAddress:  committeeHotKeyAddress,
		},
	})
	return builder
}

func (builder *TxBuilder) CommitteeColdResign(
	committeeColdKeyAddress string,
	anchor *types.Anchor,
) *TxBuilder {

	builder.TxBuilderBody.Certificates = append(builder.TxBuilderBody.Certificates, types.BasicCertificate{
		Inner: types.CommitteeColdResign{
			CommitteeColdKeyAddress: committeeColdKeyAddress,
			Anchor:                  anchor,
		},
	})
	return builder
}

func (builder *TxBuilder) DRepRegistration(
	drepID string,
	coin uint64,
	anchor *types.Anchor,
) *TxBuilder {
	builder.TxBuilderBody.Certificates = append(builder.TxBuilderBody.Certificates, types.BasicCertificate{
		Inner: types.DRepRegistration{
			DrepID: drepID,
			Coin:   coin,
			Anchor: anchor,
		},
	})
	return builder
}

func (builder *TxBuilder) DRepDeregistration(
	drepID string,
	coin uint64,
) *TxBuilder {
	builder.TxBuilderBody.Certificates = append(builder.TxBuilderBody.Certificates, types.BasicCertificate{
		Inner: types.DRepDeregistration{
			DrepID: drepID,
			Coin:   coin,
		},
	})
	return builder
}

func (builder *TxBuilder) DRepUpdate(
	drepID string,
	anchor *types.Anchor,
) *TxBuilder {
	builder.TxBuilderBody.Certificates = append(builder.TxBuilderBody.Certificates, types.BasicCertificate{
		Inner: types.DRepUpdate{
			DrepID: drepID,
			Anchor: anchor,
		},
	})
	return builder
}

// pub fn certificate_script
// pub fn certificate_tx_in_reference
// pub fn certificate_redeemer_value
