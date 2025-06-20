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

func (builder *TxBuilder) CertificateScript(scriptCbor string, version *types.LanguageVersion) *TxBuilder {
	// Pop
	lastCert := builder.TxBuilderBody.Certificates[len(builder.TxBuilderBody.Certificates)-1]
	builder.TxBuilderBody.Certificates = builder.TxBuilderBody.Certificates[:len(builder.TxBuilderBody.Certificates)-1]
	if lastCert != nil {
		panic("Undefined certificate")
	}
	switch cert := lastCert.(type) {
	case types.BasicCertificate:
		if version != nil {
			builder.TxBuilderBody.Certificates = append(
				builder.TxBuilderBody.Certificates,
				types.ScriptCertificate{
					Cert:     cert.Inner,
					Redeemer: nil,
					ScriptSource: types.ProvidedScriptSource{
						ScriptCbor:      scriptCbor,
						LanguageVersion: *version,
					},
				},
			)
		} else {
			builder.TxBuilderBody.Certificates = append(
				builder.TxBuilderBody.Certificates,
				types.SimpleScriptCertificate{
					Cert: cert.Inner,
					SimpleScriptSource: types.ProvidedSimpleScriptSource{
						ScriptCbor: scriptCbor,
					},
				})
		}
	case types.ScriptCertificate:
		if version == nil {
			panic("Language version has to be defined for plutus certificates")
		}
		builder.TxBuilderBody.Certificates = append(
			builder.TxBuilderBody.Certificates,
			types.ScriptCertificate{
				Cert:     cert.Cert,
				Redeemer: cert.Redeemer,
				ScriptSource: types.ProvidedScriptSource{
					ScriptCbor:      scriptCbor,
					LanguageVersion: *version,
				},
			})
	case types.SimpleScriptCertificate:
		panic("Native script cert had its script defined twice")
	}
	return builder
}

func (builder *TxBuilder) CertificateTxInReference(
	txHash string,
	txIndex uint32,
	scriptHash string,
	version *types.LanguageVersion,
	scriptSize uint,
) *TxBuilder {
	// Pop
	lastCert := builder.TxBuilderBody.Certificates[len(builder.TxBuilderBody.Certificates)-1]
	builder.TxBuilderBody.Certificates = builder.TxBuilderBody.Certificates[:len(builder.TxBuilderBody.Certificates)-1]
	if lastCert != nil {
		panic("Undefined certificate")
	}
	switch cert := lastCert.(type) {
	case types.BasicCertificate:
		if version != nil {
			builder.TxBuilderBody.Certificates = append(
				builder.TxBuilderBody.Certificates,
				types.ScriptCertificate{
					Cert:     cert.Inner,
					Redeemer: nil,
					ScriptSource: types.InlineScriptSource{
						RefTxIn: types.RefTxIn{
							TxHash:  txHash,
							TxIndex: txIndex,
							// Script size is already accounted for in script source
							ScriptSize: nil,
						},
						ScriptHash:      scriptHash,
						LanguageVersion: *version,
						ScriptSize:      scriptSize,
					},
				},
			)
		} else {
			builder.TxBuilderBody.Certificates = append(
				builder.TxBuilderBody.Certificates,
				types.SimpleScriptCertificate{
					Cert: cert.Inner,
					SimpleScriptSource: types.InlineSimpleScriptSource{
						RefTxIn: types.RefTxIn{
							TxHash:  txHash,
							TxIndex: txIndex,
							// Script size is already accounted for in script source
							ScriptSize: nil,
						},
						SimpleScriptHash: scriptHash,
						ScriptSize:       scriptSize,
					},
				})
		}
	case types.ScriptCertificate:
		if version == nil {
			panic("Language version has to be defined for plutus certificates")
		}
		builder.TxBuilderBody.Certificates = append(
			builder.TxBuilderBody.Certificates,
			types.ScriptCertificate{
				Cert:     cert.Cert,
				Redeemer: cert.Redeemer,
				ScriptSource: types.InlineScriptSource{
					RefTxIn: types.RefTxIn{
						TxHash:  txHash,
						TxIndex: txIndex,
						// Script size is already accounted for in script source
						ScriptSize: nil,
					},
					ScriptHash:      scriptHash,
					LanguageVersion: *version,
					ScriptSize:      scriptSize,
				},
			})
	case types.SimpleScriptCertificate:
		panic("Native script cert had its script defined twice")
	}
	return builder
}

func (builder *TxBuilder) CertificateRedeemerValue(redeemer WRedeemer) *TxBuilder {
	// Pop
	lastCert := builder.TxBuilderBody.Certificates[len(builder.TxBuilderBody.Certificates)-1]
	builder.TxBuilderBody.Certificates = builder.TxBuilderBody.Certificates[:len(builder.TxBuilderBody.Certificates)-1]
	if lastCert != nil {
		panic("Undefined certificate")
	}
	rawRedeemer, err := redeemer.Data.ToCbor()
	if err != nil {
		panic("Error converting certificate redeemer to CBOR")
	}
	currentRedeemer := types.Redeemer{
		Data:    rawRedeemer,
		ExUnits: redeemer.ExUnits,
	}
	switch cert := lastCert.(type) {
	case types.BasicCertificate:
		builder.TxBuilderBody.Certificates = append(
			builder.TxBuilderBody.Certificates,
			types.ScriptCertificate{
				Cert:         cert.Inner,
				Redeemer:     &currentRedeemer,
				ScriptSource: nil,
			},
		)
	case types.ScriptCertificate:
		builder.TxBuilderBody.Certificates = append(
			builder.TxBuilderBody.Certificates,
			types.ScriptCertificate{
				Cert:         cert.Cert,
				Redeemer:     &currentRedeemer,
				ScriptSource: cert.ScriptSource,
			},
		)
	case types.SimpleScriptCertificate:
		panic("Native script cert cannot use redeemers")
	}
	return builder
}
